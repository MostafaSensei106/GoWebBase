document.addEventListener('DOMContentLoaded', () => {
    const form = document.getElementById('user-form');
    const errorMessage = document.getElementById('error-message');
    const successMessage = document.getElementById('success-message');

    if (form) {
        form.addEventListener('submit', (event) => {
            const name = document.getElementById('name').value.trim();
            const address = document.getElementById('address').value.trim();

            errorMessage.style.display = 'none';
            successMessage.style.display = 'none';

            if (name === '' || address === '') {
                event.preventDefault();
                errorMessage.textContent = 'Both name and address are required.';
                errorMessage.style.display = 'block';

                errorMessage.scrollIntoView({ behavior: 'smooth', block: 'nearest' });
                return;
            }

            if (name.length < 2) {
                event.preventDefault();
                errorMessage.textContent = 'Name must be at least 2 characters long.';
                errorMessage.style.display = 'block';
                return;
            }

            if (address.length < 5) {
                event.preventDefault();
                errorMessage.textContent = ' Address must be at least 5 characters long.';
                errorMessage.style.display = 'block';
                return;
            }

            successMessage.textContent = '✓ Form is valid and ready to submit!';
            successMessage.style.display = 'block';
        });

        const nameInput = document.getElementById('name');
        const addressInput = document.getElementById('address');

        nameInput.addEventListener('input', () => {
            if (errorMessage.style.display === 'block') {
                errorMessage.style.display = 'none';
            }
        });

        addressInput.addEventListener('input', () => {
            if (errorMessage.style.display === 'block') {
                errorMessage.style.display = 'none';
            }
        });
    }
});
