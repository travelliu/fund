docker run -d  \
  -v `pwd`/db:/app/db \
  -p 8081:8081 \
  travelliu/fund
