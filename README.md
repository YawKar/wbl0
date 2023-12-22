# WBL
- Сервис слушает поток заказов в nats-streaming топике
- Сервис кэширует их in-memory с конфигурируемыми cleanup interval и cache expiration duration
- Сервис имеет два замаунченных ресурса на `order/` и `view/`
  - `/order`
    - `/{uuid}`: возвращает json с информацией о заказе без инфы про платёж, доставку и товары
    - `/{uuid}/payment`: возвращает json с информацией про платёж заказа
    - `/{uuid}/delivery`: возвращает json с информацией про доставку заказа
    - `/{uuid}/items`: возвращает json с информацией по всем товарам в заказе
  - `/view`
    - `/{uuid}`: отображает bootstrap server side render со всей информацией по заказу
- Nats-streaming & Postgres поднимаются в `docker compose`
- Миграции базы данных накатываются через `dbmate`
- Protobuff модели компилируются через хелпер скрипт `compile_protos.sh`
# View
![example_gif](https://raw.githubusercontent.com/yawkar/wbl0/media/media/ui-optimized.gif)

# CLI конфигурация
## cmd/server
```
Usage of server:
  -addr string
        set server's address (server listens to and serves addr:port) (default "localhost")
  -cache-cleanup duration
        set cache cleanup interval (default 6m0s)
  -cache-expire duration
        set cache expiration time (default 3m0s)
  -client-id string
        set nats-streaming client's id (default "default_client")
  -cluster-id string
        set nats-streaming cluster's id (default "default_cluster")
  -db-url string
        set database url
  -log-level int
        debug: 0; info: -4; warn: 4; error: 8
  -nats-url string
        set nats-streaming node's url (default "nats://127.0.0.1:4222")
  -port uint
        set server's port (server listens to and serves addr:port) (default 8080)
```
## cmd/publisher
```
Usage of publisher:
  -client-id string
        set nats-streaming client's id (default "default_client")
  -cluster-id string
        set nats-streaming cluster's id (default "default_cluster")
  -log-level int
        debug: 0; info: -4; warn: 4; error: 8
  -nats-url string
        set nats-streaming node's url (default "nats://127.0.0.1:4222")
  -seed int
        set seed for faker (default 42)
  -spam-duration duration
        set spam duration (default 1m0s)
  -spam-rate duration
        set spam rate (2s means 1 message every 2 seconds) (default 1s)
```
