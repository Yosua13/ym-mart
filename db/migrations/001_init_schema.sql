-- Hapus tabel jika sudah ada untuk memastikan awal yang bersih
DROP TABLE IF EXISTS order_items, orders, cart_items, carts, products, stores CASCADE;

-- Tabel Toko
CREATE TABLE stores (
    store_id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    city VARCHAR(100),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

-- Tabel Produk
CREATE TABLE products (
    product_id SERIAL PRIMARY KEY,
    store_id INT NOT NULL REFERENCES stores(store_id),
    name VARCHAR(255) NOT NULL,
    description TEXT,
    price DECIMAL(12, 2) NOT NULL,
    stock INT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

-- Tabel Keranjang Belanja
CREATE TABLE carts (
    cart_id SERIAL PRIMARY KEY,
    user_id INT NOT NULL UNIQUE,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

-- Tabel Item di dalam Keranjang
CREATE TABLE cart_items (
    cart_item_id SERIAL PRIMARY KEY,
    cart_id INT NOT NULL REFERENCES carts(cart_id),
    product_id INT NOT NULL REFERENCES products(product_id),
    quantity INT NOT NULL CHECK (quantity > 0),
    UNIQUE(cart_id, product_id) 
);

-- Tabel Transaksi/Order
CREATE TABLE orders (
    order_id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    invoice_number VARCHAR(255) NOT NULL UNIQUE,
    total_amount DECIMAL(12, 2) NOT NULL,
    status VARCHAR(50) DEFAULT 'Menunggu Pembayaran',
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

-- Tabel Item di dalam Order
CREATE TABLE order_items (
    order_item_id SERIAL PRIMARY KEY,
    order_id INT NOT NULL REFERENCES orders(order_id),
    product_id INT NOT NULL,
    product_name VARCHAR(255) NOT NULL, 
    price_at_purchase DECIMAL(12, 2) NOT NULL, 
    quantity INT NOT NULL
);