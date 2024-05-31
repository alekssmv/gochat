<h2>Что я хочу сделать</h2>
<p>Создать приложение, в котором присутствует: регистрация и авторизация, поддержка авторизации, возможность добавления контактов, возможность отправки сообщений этим контактам с использованием технологии WebSocket.</p>
<h2>Технологии</h2>
<p>Backend-код написан на Go с нуля без использования шаблона, но с использованием сторонних Go-модулей. Для создания CLI-команд выбрал Go-библиотеку <a href="https://github.com/spf13/cobra">cobra-cli</a>. Для Frontend выбрал React.js + Vite. База данных у проекта — PostgreSQL, веб-сервер — Nginx.</p>
<h2>Что сделано</h2>
<h3>Frontend:</a></h3>
<ol>
  <li>Отображение страниц регистрации, логина и списка контактов (Страница со списоком контактов пока что без функционала).</li>
  <li>Динамическое отображение ссылок в навигации.</li>
  <li>Динамическое отображение кнопки "logout".</li>
  <li>Сброс кэша "фронта" после нажатия на "logout".</li>
  <li>Отображение имени авторизованного пользователя.</li>
</ol>
<h3>Backend:</a></h3>
<ol>
  <li>Работа с БД посредством ORM <a href="https://github.com/go-gorm/gorm">gorm</a>.</li>
  <li>Авторизация.</li>
  <li>Регистрация.</li>
  <li>Сохранение новых пользователей в БД.</li>
</ol>
<h3><a href="https://github.com/alekssmv/gochat/tree/main/Cli">Cli команды:</a></h3>
<ol>
  <li>Запуск миграций в бд.</li>
  <li>Удаление всех данных из бд.</li>
</ol>
<h3>Тестовая документация:</a></h3>
<ol>
  <li><a href="https://docs.google.com/spreadsheets/d/1j8t9UMbRxWT9KtvN-fLcNyhz7qGfPgp8y7wH0q3xWBM/edit?usp=sharing">Чек лист</a></li>
</ol>
<h2>Запуск проекта на Ubuntu</h2>
<p>Для запуска проекта нужен docker-compose.</p>

```bash
sudo git clone https://github.com/alekssmv/gochat && cd gochat && cp .env.example .env && sudo docker-compose up --build
```
<p>После запуска проекта нужно выполнить <a href="https://github.com/alekssmv/gochat/tree/main/Cli">миграции</a>.</p>
<p>Также посмотреть как работает проект можно <a href=http://194.35.13.18:81/>здесь</a>.</p>
