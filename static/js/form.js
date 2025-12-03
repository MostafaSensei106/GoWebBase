document.addEventListener('DOMContentLoaded', () => {
    const form = document.getElementById('user-form');
    const errorMessage = document.getElementById('error-message');

    if (form) {
        form.addEventListener('submit', (event) => {
            const name = document.getElementById('name').value.trim();
            const address = document.getElementById('address').value.trim();

            if (name === '' || address === '') {
                event.preventDefault();
                errorMessage.textContent = 'Both name and address are required.';
                errorMessage.style.display = 'block';
            } else {
                errorMessage.style.display = 'none';
            }
        });
    }
});
