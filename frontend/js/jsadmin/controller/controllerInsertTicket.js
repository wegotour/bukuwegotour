const getTokenFromCookies = (cookieName) => {
    const cookies = document.cookie.split(';')
    for (const cookie of cookies) {
      const [name, value] = cookie.trim().split('=')
      if (name === cookieName) {
        return value
      }
    }
    return null
  }
  
  const showAlert = (message, type = 'success') => {
    Swal.fire({
      icon: type,
      text: message,
      showConfirmButton: false,
      timer: 1500,
    })
  }
  
  const insertTicket = async (event) => {
    event.preventDefault()
  
    const token = getTokenFromCookies('Login')
  
    if (!token) {
      showAlert('Header Login Not Found', 'error')
      return
    }
  
    const URLInsertTicket = 'https://asia-southeast2-wegotour-403712.cloudfunctions.net/insertdataticket'
  
    const myHeaders = new Headers()
    myHeaders.append('Login', token)
    myHeaders.append('Content-Type', 'application/json')
  
    const requestOptions = {
      method: 'POST',
      headers: myHeaders,
      body: JSON.stringify({  
        nomorid: parseInt(document.getElementById('newNomorID').value),
        title: document.getElementById('newTitle').value,
        description: document.getElementById('newDeskripsi').value,
        image: document.getElementById('newImage').value,
        status: document.getElementById('newStatus').value === 'active' ? true : false,
      }),
      redirect: 'follow',
    }
  
    try {
      const response = await fetch(URLInsertTicket, requestOptions)
      const data = await response.json()
  
      if (data.status === false) {
        showAlert(data.message, 'error')
      } else {
        showAlert('Ticket data successfully added!', 'success')
        window.location.href = 'ticket.html'
      }
    } catch (error) {
      console.error('Error:', error)
    }
  }
  
  document.getElementById('newTicketForm').addEventListener('submit', insertTicket)