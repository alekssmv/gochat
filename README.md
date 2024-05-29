<h2>Описание</h2>
<p><em>Проект пока что очень сырой</em></p>
Backend-код написан на Go с нуля без использования шаблона, но с использованием сторонних Go-модулей. Для создания CLI-команд выбрал Go-библиотеку <a href="https://github.com/spf13/cobra">cobra-cli</a>. Для Frontend выбрал React + Vite. База данных у проекта — PostgreSQL, веб-сервер — Nginx.
<h2>Что сделано</h2>
<h3><a href="https://github.com/Alekssmv/GoChat/tree/main/src/Frontend">Frontend:</a></h3>
<ol>
  <li>Отображение страниц регистрации, логина и списка контактов</li>
  <li>Динамическое отображение навигации</li>
</ol>
<h3><a href="https://github.com/Alekssmv/GoChat/tree/main/src/Backend">Backend:</a></h3>
<ol>
  <li>Работа с бд посредством ORM <a href="https://github.com/go-gorm/gorm">gorm</a></li>
  <li>Работа с сессиями</li>
</ol>
<h3><a href="https://github.com/Alekssmv/GoChat/tree/main/src/Cli">Cli команды:</a></h3>
<ol>
  <li>Запуск миграций в бд</li>
  <li>Удаление данных из бд</li>
</ol>
<h1>Запуск проекта</h1>
<div id="code-container" style="position: relative;">
  <pre><code id="code-snippet">sudo docker-compose up</code></pre>
</div>
