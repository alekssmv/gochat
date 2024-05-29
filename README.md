<h2>Описание</h2>
Backend-код написан на go с нуля без использования шаблона, но использованием сторонних go-модулей. Для создания CLI-команд выбрал go библиотеку <a href="https://github.com/spf13/cobra">cobra-cli</a>. Для Frontend выбрал React + Vite. Бд у проекта - postgres, веб-сервер - nginx.
<h2>Что сделано</h2>
<h3>Frontend:</h3>
<p>*Отображение страниц регистрации и логина</p>
<h3>Backend:</h3>
<p>*Внедрил gorm в проект</p>
<p>*Добавление зарегистрированного пользователя в бд</p>
<h3>Cli команды:</h3>
<a href="https://github.com/Alekssmv/GoChat/tree/main/src/Cli">Список команд</a>
<br></br>
<p>*Запуск миграций в бд</p>
<p>*Удаление данных из бд</p>
<h1>Запуск проекта</h1>
<div id="code-container" style="position: relative;">
  <pre><code id="code-snippet">sudo docker-compose up</code></pre>
</div>
