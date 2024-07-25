# spc (stock portfolio charts)

用于统计各股票持仓比例，用图表展示


## 部署
```bash
docker run -itd -v ./data:/data registry.cn-hongkong.aliyuncs.com/tonyzzp/spc
```

docker-compose.yaml
```yaml
services:
  server:
    container_name: spc
    image: registry.cn-hongkong.aliyuncs.com/tonyzzp/spc
    restart: always
    volumes:
      - ./data:/data

volumes:
  spc_data:
    external: true
```
