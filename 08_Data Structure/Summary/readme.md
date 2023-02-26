# Data Structure


## Pengertian
Data structure adalah sebuah cara dalam menyimpan maupun mengatur data dalam komputer secara terstruktur. Data data ini kemudian dapat memudahkan dalam pengolahan data kedepannya. Data structure dibagi menjadi 3:
    - Array
    - Slice
    - Map

## Jenis-jenis Data Structure

### Array
Array adalah kumpulan nilai atau elemen yang disimpan secara bersamaan dalam satu variabel. Biasanya setiap nilai atau elemen diberi nomor indeks yang berbeda, dimana nomor indeks ini nanti akan digunakan untuk mengakses nilai tersebut. Array didalam *Go*, harus dideklarasikan beserta panjang dari array tersebut. Array ini sendiri bersifat statis, dimana jika panjang array sudah penuh, maka kita tidak dapat lagi menambahkan elemen lagi kedalamnya.


### Slice
Slice mungkin mirip seperti array, namun yang membedakan adalah dalam slice, kita tidak perlu mendeklarasikan panjang dari slice itu sendiri. Bisa dibilang slice adalah array yang tidak di deklarasikan panjangnya. Slice ini bersifat dinamis dimana kita dapat terus memasukkan berapa banyak elemen kedalamnya.


### Map
Map digunakan untuk menyimpan data dalam bentu pasangan key-value. Dimana setiap key akan memiliki value itu sendiri. Dalam map, satu value hanya dapat menampung satu value. Kegunaan dari map itu sendiri adalah supaya dapat cepat diakses berdasarkan key nya.