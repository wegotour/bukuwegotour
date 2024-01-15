  // Retrieve selected ticket data from localStorage
  var selectedTicket = JSON.parse(localStorage.getItem('selectedTicket'));

  // Display ticket information on the checkout page
  document.getElementById('ticketName').value = selectedTicket.name;
  document.getElementById('ticketPrice').value = selectedTicket.price;
  document.getElementById('ticketQuantity').value = selectedTicket.quantity;
  document.getElementById('ticketTotal').value = selectedTicket.total;