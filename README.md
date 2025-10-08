# Soal test maggang Backend engineer dengan Springboot

Berikut adalah soal/pertanyaan yang perlu dijawab oleh peserta maggang

## knowledge base

1. Apa yang anda ketahui tentang Rest API?
2. Apa yang anda ketahui tentang Server side and Client side processing?
3. Apa yang anda ketahui tentang Monolith dan Microservices, berikan contohnya?
4. Apa yang anda ketahui tentang Design pattern inversion of Control serta Dependency Injection?
5. Apa yang anda ketahui tentang Golang programming dan serta fiber framework?
6. Apa yang anda ketahui tentang Pointer serta fungsinya?

## Design modules

Dalam suatu schenario ada requirement membuat aplikasi e-commerse seperti Tokopedia seperti berikut:

1. Catalog, pelanggan mencari product di toko
    ![catalog](imgs/catalog.png)
2. Item, bisa melihat detail informasi produk
    ![items](imgs/item.png)
3. Cart, pelanggan bisa menambahkan produk yang ingin di beli ke keranjang
    ![cart](imgs/cart.png)
4. Setelah di checkout, masuk ke list transaction
    ![list-transaction](imgs/list-transaction.png)
5. Kita juga bisa liat detail transactionya
    ![detail-transaction](imgs/detail-transaction.png)

Kemudian temen-temen buat design database, module (monolith/microservices) berdasarkan gambar atau schenario tersebut. Serta jelakan mengapa menggunakan design tersebut.

## Praktek

Berdasarkan analisa tersebut, buat project monorepo (pada repository ini) dengan menggunakan framework [gofiber/fiber](https://gofiber.io/) seperti berikut specifikasinya:

- Database: `PostgreSQL 15`
- Golang version: `golang 1.20.x or later`
- gofiber version: `v2`

terkait design system Toko, Barang, Pembelian pada ecommerse tersebut.