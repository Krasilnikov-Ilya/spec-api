```Простое GO приложение, возвращающее на GET запросы содержимое .yaml файла.```
 
Поддерживает обработку OPTIONS запросов с ответом 204 accepted и добавляет к ответу CORS хидеры, разрешающие встраивание результата.
 
Для корректной работы файл должен иметь название openapi.yaml и находиться в том же каталоге, что и главный класс программы.