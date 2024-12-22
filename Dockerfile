# Используем официальный образ Go
FROM golang:1.23.4 AS builder

# Устанавливаем рабочую директорию
WORKDIR /app

# Устанавливаем необходимые системные зависимости
RUN apt update && apt install -y \
    libxkbfile-dev \
    libxrandr-dev \
    pkg-config \
    xorg-dev && \
    apt-get clean && rm -rf /var/lib/apt/lists/*

# Копируем файлы проекта
COPY . .

# Скачиваем зависимости
RUN go mod tidy

# Сборка приложения
RUN go build -o app .

# Финальный образ
FROM debian:bookworm-slim

# Устанавливаем рабочую директорию
WORKDIR /app

# Устанавливаем необходимые зависимости для Chromium, Xvfb и VNC
RUN apt update && apt install -y \
    chromium \
    libx11-dev \
    xvfb \
    x11vnc \
    fluxbox \
    libx11-6 \
    libxkbfile1 \
    libxrandr2 \
    libxinerama1 \
    libxcursor1 \
    libgtk-3-0 \
    fonts-liberation && \
    apt-get clean && rm -rf /var/lib/apt/lists/*

# Копируем скомпилированное приложение
COPY --from=builder /app/app .

# Устанавливаем переменные окружения
ENV DISPLAY=:99.0

# Открываем порт для VNC
EXPOSE 5900

# Запуск Xvfb, Fluxbox и VNC
CMD Xvfb :99 -screen 0 1920x1080x16 > /dev/null 2>&1 & \
    fluxbox > /dev/null 2>&1 & \
    x11vnc -display :99 -forever -nopw -rfbport 5900 > /dev/null 2>&1 & \
    ./app