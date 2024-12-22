
# Test Mayflower

## Описание

Программа выполняет автоматизированное тестирование веб-приложения, включая:
- Авторизацию.
- Навигацию по страницам.
- Выполнение действий и проверку результатов.

Запуск приложения осуществляется через Docker, с возможностью просмотра графического интерфейса через VNC.

---

## Требования

1. **Docker** и **Docker Compose**:
   - Убедитесь, что они установлены:
     ```bash
     docker --version
     docker compose version
     ```
   - Если Docker не установлен, выполните следующие шаги:

     ### Установка Docker (на Linux):
     1. Обновите пакеты:
        ```bash
        sudo apt update
        sudo apt upgrade
        ```
     2. Установите Docker:
        ```bash
        sudo apt install -y docker.io
        ```
     3. Запустите и настройте автозапуск:
        ```bash
        sudo systemctl start docker
        sudo systemctl enable docker
        ```

     ### Установка Docker (на macOS и Windows):
     1. Скачайте установщик Docker Desktop с [официального сайта](https://www.docker.com/products/docker-desktop).
     2. Установите приложение и следуйте инструкциям установщика.

2. **VNC-клиент**:
   Установите VNC-клиент для подключения к графическому интерфейсу:
   - [RealVNC](https://www.realvnc.com/).
   - [TigerVNC](https://tigervnc.org/).

---

## Как запустить

### 1. Сборка Docker-образа

1. Клонируйте репозиторий:
   ```bash
   git clone https://github.com/your-username/test-mayflower.git
   cd test-mayflower
   ```

2. Соберите Docker-образ:
   ```bash
   docker build -t test-mayflower .
   ```

---

### 2. Запуск Docker-контейнера

1. Запустите контейнер с пробросом порта для VNC:
   ```bash
   docker run --rm -p 5900:5900 -e LOGIN_USERNAME=<ваш логин> -e LOGIN_PASSWORD=<ваш пароль> test-mayflower
   ```

   #### Параметры:
   - `-p 5900:5900`: Пробрасывает порт для VNC-сервера.
   - `-e LOGIN_USERNAME`: Логин для авторизации в веб-приложении.
   - `-e LOGIN_PASSWORD`: Пароль для авторизации.

   #### Пример:
   ```bash
   docker run --rm -p 5900:5900 -e LOGIN_USERNAME=test@test.com -e LOGIN_PASSWORD=password test-mayflower
   ```

---

### 3. Подключение к VNC

1. Установите и запустите VNC-клиент (например, RealVNC или TigerVNC).
2. Подключитесь к серверу:
   ```
   localhost:5900
   ```
3. Вы увидите рабочий стол с открытым браузером, где будет выполняться тестирование.
