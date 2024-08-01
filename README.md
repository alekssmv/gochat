<h2>Что я хочу сделать</h2>
<p>Создать приложение, в котором присутствует: регистрация и авторизация, поддержка авторизации, возможность добавления контактов, возможность отправки сообщений этим контактам с использованием технологии WebSocket.</p>
<h2>Технологии</h2>
<li>Микро-сервисы: <a href="https://github.com/gin-gonic/gin">gin-gonic</a>.</li>
<li>Фронтэнд: Vue + TypeScript.</li>
<h2>Что сделано</h2>
<h3>Frontend:</h3>
<ol>
  <li>Реализована страница авторизации. (без Vue)</li>
  <li>Реализована страница регистрации. (без Vue)</li>
</ol>
<h3>Backend:</h3>
<ol>
  <li>Реализован файловый сервис</li>
</ol>
<h3>Endpoints:</h3>
<li> GET - /login - страница авторизации</li>
<li> GET - /register - страница регистрации</li>
<h2>TODO</h2>
<li>Реализовать фронтенд на vue + ts</li>
<li>Реализовать микро-сервис аутентификации/авторизации</li>
<h2>Запуск проекта на Ubuntu 22.04</h2>
<p>Для запуска проекта нужен Docker(24.0.7), Docker-Compose(1.29.2) </p>

```bash
sudo docker-compose up
```
