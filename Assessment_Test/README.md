# Assessment Test E-Commerce
### Deskripsi Program
Aplikasi E-Commerce ini memliki fitur:
1.	Registrasi pengguna
2.	Login pengguna
3.	List produk
4.	Cart
5.	Checkout produk

### Problem & Motivation
E-Commerce merupakan teknologi dalam bentuk aplikasi yang menghubungkan bisnis, consumer dan masyarakat dalam pertukaran barang, jasa dan informasi secara elektronik. Pada aplikasi yang dikembangkan ini, jenis e-commerce yang diterapkan yaitu B2C atau business to consumer dimana model penjualan dari bisnis langsung ke penjualan.

Dalam model bisnis B2C, pengguna aplikasi dapat dibagi menjadi tiga kategori. Berikut adalah penjelasan mengenai setiap kategori pengguna pada aplikasi ini.

1.	Public

Pengguna akan masuk kategori public apabila ketika mengakses aplikasi, pengguna tersebut tidak melalui proses login terlebih dahulu. Pengguna kategori public tetap bisa mengakses aplikasi namun fitur yang tersedia dibatasi. Pengguna kategori public hanya bisa melihat produk apa yang disediakan seperti list produk, detail dari produk, dan melakukan filter dalam list produk.

2.	User

Pengguna akan masuk kategori user apabila ketika mengakses aplikasi, pengguna tersebut melakukan proses login terlebih dahulu dan teridentifikasi sebagai user. Pengguna kategori user dapat menikmati fitur seperti pegguna kategori public. Pengguna kategori user pun dapat menikmati fitur seperti memasukan produk yang diinginkan ke dalam cart atau keranjang belanja, melihat cart dan melakukan checkout atau pembelian produk.

3.	Admin

Pengguna akan masuk kategori admin apabila ketika mengakses aplikasi, pengguna tersebut melakukan proses login terlebih dahulu dan teridentifikasi sebagai admin. Pengguna kategori admin dapat menikmati fitur seperti user namun pengguna kategori admin tidak dapat memasukan produk ke cart dan tidak dapat melakukan checkout. Pengguna dengan kategori admin dapat memasukan produk baru ke list produk dan mengubah data produk.

### How to Run
Untuk dapat menggunakan aplikasi ini, lakukan clone pada repository ini.

`git clone https://github.com/febbryandika/aqua-developer.git`

Setelah melakukan clone, ubah lah informasi pada docker-compose.yml. Informasi yang diperlukan yaitu database host, database user, database password, database name, dan database port.

Kemudian jalankan docker compose.

`docker-compose up`

### Cara Penggunaan
Cara penggunaan aplikasi dapat dilihat di dokumentasi API berikut ([ini](https://documenter.getpostman.com/view/20138613/2s8YeoQDp6)).
