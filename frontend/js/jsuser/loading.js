setTimeout(function () {
    document.getElementById('loader-wrapper').style.display = 'none';
    document.getElementById('main-content').style.display = 'block';
  }, 1500); 

  		// Menghilangkan overlay saat halaman selesai dimuat
          document.onreadystatechange = function () {
            if (document.readyState === 'complete') {
              setTimeout(function () {
                document.getElementById('loader-wrapperz').style.display = 'none';
              }, 2000); // Sesuaikan timeout dengan durasi animasi CSS
            }
          };