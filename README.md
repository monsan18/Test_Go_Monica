# RUNNING STEP GOLANG SERVER ECHO FRAMEWORK

1. Open Terminal in project folder
2. Run "go run server.go"
3. Allow access when a promt show.
4. Open Postman
5. Open API documentation
6. Run HTTP Request from Postman according to documentation.
7. Start Apache Xampp and MySql if has not started yet.
8. Create new database 'ecommerce'.
9. Import .sql to database 'ecommerce' from the .sql file.

# ###############
Sample Request : See postman collection Test_Golang_Monica.postman_collection.json.

# Essay Answer #####

1. Project Review adalah sebuah proses untuk meng-evaluasi keberhasilan proyek dan memutuskan akan terus menggunakan sumber daya.
    Project review dapat dilakukan dalam beberapa bentuk, yaitu :
    - rapat / meeting dengan tim projek untuk menilai status projek saat ini,
    - tinjauan di akhir projek.

Project Planning adalah sebuah proses perencanaan proyek yang mencakup langkah-langkah untuk menyelesaikan proyek dalam jangka waktu tertentu. Biasanya dilakukan sebelum proyek dilaksanakan. Project planning juga biasanya disajikan dalam bentuk gantt chart untuk memudahkan memvisualisasikan alur waktu proyek.

2. Load balance adalah sebuah teknik yang mendistribusikan lalu lintas jaringan atau aplikasi di beberapa server. Hal tersebut bertujuan untuk meningkatkan kinerja, keandalan, dan ketersediaan. 
    Di Amazon EC2, Elastic Load Balancing (ELB) adalah layanan load balancer yang disediakan oleh AWS.

    Security Group pada AWS berperan sebagai firewall virtual untuk instance EC2, guna mengendalikan lalu lintas yang masuk dan keluar.
    Cara kerja security group pada AWS yaitu membantu dalam mengamankan lingkungan cloud dengan mengendalikan bagaimana lalu lintas akan diizinkan masuk ke mesin EC2. 
    Dengan security group, kita dapat memastikan bahwa semua lalu lintas yang mengalir di tingkat instans hanya melalui port dan protokol yang telah ditetapkan.
    Seperti whitelist (daftar putih), aturan security group tidak selalu permisif. Tidak mungkin kita membuat aturan yang menolak akses.
    Security group bersifat stateful, jadi jika permintaan masuk lolos, maka permintaan keluar juga akan lolos.

3. Cara mengatasi memory leak pada Golang :
    - menghindari referensi melingkar antara objek untuk memungkinkan pengumpulan sampah yang tepat.
    - kelola goroutine dan saluran dengan tepat untuk mencegah kebocoran dari goroutine yang diblokir atau tidak dihentikan.
    - lepaskan sumber daya yang tidak diperlukan seperti file, koneksi jaringan, dan saluran saat tidak lagi diperlukan.
    - terapkan kebijakan kadaluarsa dengan time-to-live (TTL) untuk cache dan variable global.


