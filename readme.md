
# Test Mayflower

## Описание

Программа выполняет автоматизированное тестирование веб-приложения, включая:
- Авторизацию.
- Навигацию по страницам.
- Выполнение действий и проверку результатов.

Запуск приложения осуществляется через Docker, с возможностью просмотра графического интерфейса через VNC, или локально на вашей машине.

Стоит упомянуть несколько моментов, которые были осознанно пропущены при написании теста из-за ограничений по времени:
- Программа не имеет возможности масштабирования. 
- Не написаны юниты тесты на вызовы пейдж-обджектов.

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

3. **Go**:
   Для локального запуска установите Go (версия 1.23 или выше):
   - [Инструкция по установке Go](https://go.dev/doc/install).

---

## Как запустить

### 1. Запуск локально

1. Клонируйте репозиторий:
   ```bash
   git clone https://github.com/your-username/test-mayflower.git
   cd test-mayflower
   ```

2. Установите переменные окружения и запустите программу:
   #### Linux/macOS:
   ```bash
   LOGIN_USERNAME=<ваш логин> LOGIN_PASSWORD=<ваш пароль> go run main.go
   ```

   #### Windows (в PowerShell):
   ```powershell
   $env:LOGIN_USERNAME="<ваш логин>"
   $env:LOGIN_PASSWORD="<ваш пароль>"
   go run main.go
   ```

   #### Пример:
   ```bash
   LOGIN_USERNAME=test@test.com LOGIN_PASSWORD=password go run main.go
   ```

---

### 2. Сборка Docker-образа

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

### 3. Запуск Docker-контейнера

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

### 4. Подключение к VNC

1. Установите и запустите VNC-клиент (например, RealVNC или TigerVNC).
2. Подключитесь к серверу:
   ```
   localhost:5900
   ```
3. Вы увидите рабочий стол с открытым браузером, где будет выполняться тестирование.