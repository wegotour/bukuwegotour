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
  
  const getAllTicket = async () => {
    const token = getTokenFromCookies('Login')
  
    if (!token) {
      Swal.fire({
        icon: 'warning',
        title: 'Authentication Error',
        text: 'You are not logged in.',
      }).then(() => {
        window.location.href = 'https://wegotour.my.id/loginadmin.html'
      })
      return
    }
  
    const URLGetAllTicket = 'https://asia-southeast2-wegotour-403712.cloudfunctions.net/wegotour'
  
    const myHeaders = new Headers()
    myHeaders.append('Login', token)
  
    const requestOptions = {
      method: 'GET',
      headers: myHeaders,
      redirect: 'follow',
    }
  
    try {
      const response = await fetch(URLGetAllTicket, requestOptions)
      const data = await response.json()
  
      if (data.status) {
        displayTicketData(data.data, 'TicketDataBody')
      } else {
        Swal.fire({
          icon: 'error',
          title: 'Error',
          text: data.message,
        })
      }
    } catch (error) {
      console.error('Error:', error)
    }
  }
  
  const deleteTicket = async (nomorId) => {
    const token = getTokenFromCookies('Login')
  
    if (!token) {
      showAlert('Header Login Not Found', 'error')
      return
    }
  
    const URLDeleteTicket = 'https://asia-southeast2-wegotour-403712.cloudfunctions.net/deletedataticket'
  
    const myHeaders = new Headers()
    myHeaders.append('Login', token)
    myHeaders.append('Content-Type', 'application/json')
  
    const requestOptions = {
      method: 'DELETE',
      headers: myHeaders,
      body: JSON.stringify({ nomorid: nomorId }),
      redirect: 'follow',
    }
  
    try {
      const response = await fetch(URLDeleteTicket, requestOptions)
      const data = await response.json()
  
      if (data.status) {
        Swal.fire({
          icon: 'success',
          title: 'Success',
          text: 'ticket successfully deleted!',
        }).then(() => {
          getAllEmployees()
        })
      } else {
        Swal.fire({
          icon: 'error',
          title: 'Error',
          text: data.message,
        })
      }
    } catch (error) {
      console.error('Error:', error)
    }
  }
  
  // Function to handle the delete confirmation
  const deleteTicketHandler = (nomorId) => {
    Swal.fire({
      title: 'Are you sure?',
      text: "You won't be able to revert this!",
      icon: 'warning',
      showCancelButton: true,
      confirmButtonColor: '#3085d6',
      cancelButtonColor: '#d33',
      confirmButtonText: 'Yes, delete it!',
    }).then((result) => {
      if (result.isConfirmed) {
        deleteTicket(nomorId)
      }
    })
  }
  
  const editTicket = (nomorId) => {
    window.location.href = `formedit_ticket.html?nomorid=${nomorId}`
  }
  // Event listener to handle clicks on the table
  document.getElementById('TicketDataBody').addEventListener('click', (event) => {
    const target = event.target
    if (target.classList.contains('edit-link')) {
      const nomorId = parseInt(target.getAttribute('data-nomorid'))
      editTicket(nomorId)
    } else if (target.classList.contains('delete-link')) {
      const nomorId = parseInt(target.getAttribute('data-nomorid'))
      deleteTicketHandler(nomorId)
    }
  })
  
  const displayTicketData = (ticketData, tableBodyId) => {
    const ticketDataBody = document.getElementById(tableBodyId)
  
    ticketDataBody.innerHTML = ''
  
    if (ticketData && ticketData.length > 0) {
      ticketData.forEach((item) => {
        const newRow = document.createElement('tr')
        newRow.innerHTML = `
          <td class="px-4 py-3">${item.nomorid}</td>
          <td class="px-4 py-3">${item.title}</td>
          <td class="px-4 py-3">${item.description}</td>
          <td class="px-4 py-3">
            <img src="${item.image}" alt="Ticket Image" style="max-width: 100px; max-height: 100px;">
          </td>
          <td class="px-4 py-3">${item.status ? 'Active' : 'Inactive'}</td>
          <td class="px-4 py-3">
            <a href="#" class="edit-link" data-nomorid="${item.nomorid}">Edit</a>
            <a href="#" class="delete-link" data-nomorid="${item.nomorid}">Delete</a>
          </td>
        `
  
        ticketDataBody.appendChild(newRow)
      })
    } else {
      ticketDataBody.innerHTML = `<tr><td colspan="6">No ticket data found.</td></tr>`
    }
  }
  
  // Initial fetch of all tickets
  getAllTicket()