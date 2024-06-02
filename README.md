<h2>Что я хочу сделать</h2>
<p>Создать приложение, в котором присутствует: регистрация и авторизация, поддержка авторизации, возможность добавления контактов, возможность отправки сообщений этим контактам с использованием технологии WebSocket.</p>
<h2>Технологии</h2>
<p>Backend-код написан на Go с нуля без использования шаблона, но с использованием сторонних Go-модулей. Для создания CLI-команд выбрал Go-библиотеку <a href="https://github.com/spf13/cobra">cobra-cli</a>. Для Frontend выбрал React.js + Vite. База данных у проекта — PostgreSQL, веб-сервер — Nginx.</p>
<h2>Что сделано</h2>
<h3>Frontend:</h3>
<ol>
  <li>Отображение страниц регистрации, логина и списка контактов (страница со списком контактов пока что без функционала).</li>
  <li>Динамическое отображение ссылок в навигации.</li>
  <li>Динамическое отображение кнопки "Logout".</li>
  <li>Сброс кэша фронтенда после нажатия на "Logout".</li>
  <li>Отображение имени авторизованного пользователя.</li>
  <li>Отображение ошибок при регистрации и авторизации.</li>
  <li>Отображение сообщений при успешной регистрации и авторизации.</li>
</ol>
<h3>Backend:</h3>
<ol>
  <li>Работа с БД посредством ORM <a href="https://github.com/go-gorm/gorm">gorm</a>.</li>
  <li>Авторизация.</li>
  <li>Регистрация.</li>
  <li>Сохранение новых пользователей в БД.</li>
  <li>Валидация.</li>
</ol>
<h3><a href="https://github.com/alekssmv/gochat/tree/main/Cli">CLI команды:</a></h3>
<ol>
  <li>Запуск миграций в БД.</li>
  <li>Удаление всех данных из БД.</li>
</ol>
<h3>Тестовая документация:</h3>
<ol>
  <li><a href="https://docs.google.com/spreadsheets/d/1j8t9UMbRxWT9KtvN-fLcNyhz7qGfPgp8y7wH0q3xWBM/edit?usp=sharing">Чек-лист</a></li>
</ol>
<h2>Запуск проекта на Ubuntu 22.04</h2>
<p>Для запуска проекта нужен Docker Compose версии 1.29.2.</p>

```bash
sudo git clone https://github.com/alekssmv/gochat && cd gochat && cp .env.example .env && sudo docker-compose up --build
