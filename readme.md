### build
```bash
make [build_all|build_api|build_listing|build_order|build_user]
```

### build image
```bash
make build TAG=v0.0.1 APP=[all|api|listing|order|user]
```

### deploy
1. mysql
import tables to db
  - `sql/db_listing.sql` to `db_listing`
  - `sql/db_order.sql` to `db_order`
  - `sql/db_user.sql` to `db_user`

1. config
```bash
kubectl create -f k8s/config_map.yml
```

1. consul
```bash
kubectl create -f k8s/consul.yml
```

1. services
```bash
kubectl create -f k8s/user.yml
kubectl create -f k8s/listing.yml
kubectl create -f k8s/order.yml
```
