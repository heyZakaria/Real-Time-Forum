//  function fetchOnlineUsers() {
//     fetch('/api/online-users')
//         .then(response => response.json())
//         .then(users => {
//             let userList = document.getElementById('user-list');
//             userList.innerHTML = ''; // Clear previous users
//             users.forEach(user => {
//                 let li = document.createElement('li');
//                 li.textContent = user.username;
//                 userList.appendChild(li);
//             });
//         })
//         .catch(error => console.error('Error fetching online users:', error));
// }

// // Refresh online users every 10 seconds
// setInterval(fetchOnlineUsers, 10000);
// fetchOnlineUsers(); // Initial call
 