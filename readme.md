# How to running this app on local
1. copy `.env-example` menjadi `.env`
2. Ubah isi file `.env`   
3. Export env menggunakan `export $(cat .env | xargs -L 1)`  
4. Jika ingin menjalankan di local, isi config docker compose lalu jalankan `docker compose up -d`
5. Lakukan migration dahulu dengan `make migrate-up`
6. Jalankan docker compose aplikasi dengan `make dc.up`
7. Untuk menghentikannya gunakan `make dc.stop`

# TODO
[X] CRUD Activity   
[X] CRUD Todo   
[ ] Fix docker compose for golang
[ ] github action / circleci to push to docker hub
[ ] Integration Testing
