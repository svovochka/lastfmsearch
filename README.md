Для сборки и запуска нужен только docker и docker-compose

Запустите сборку `docker-compose -f docker-compose-build.yml up`

Запустите контейнер с сервисом `docker-compose up`

Плейграунд доступен по адресу http://localhost:8080/playground

Пример GraphQL запроса:

```
{
  findTracksByName(name:"Hallelujah"){
    name
    url
    listeners
    artist{
      id
      name
      url
      summary
    }
  }
}
```