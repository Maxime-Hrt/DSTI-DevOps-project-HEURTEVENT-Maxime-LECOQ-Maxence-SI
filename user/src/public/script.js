document.getElementById('loadContactsBtn').addEventListener('click', loadContacts);

document.getElementById('showAddContactFormBtn').addEventListener('click', function() {
    document.getElementById('contactFormDiv').classList.toggle('hidden');
    resetForm();
});

document.getElementById('contactForm').addEventListener('submit', function(e) {
    e.preventDefault();
    const contactData = {
        name: document.getElementById('name').value,
        email: document.getElementById('email').value,
        phone: document.getElementById('phone').value
    };

    const contactId = document.getElementById('contactId').value;
    if (contactId) {
        updateContact(contactId, contactData);
    } else {
        createContact(contactData);
    }
});

function loadContacts() {
    fetch('http://localhost:8080/contacts')
        .then(response => response.json())
        .then(data => {
            const contactsDiv = document.getElementById('contacts');
            contactsDiv.innerHTML = '';
            data.forEach(contact => {
                const contactDiv = document.createElement('div');
                contactDiv.innerHTML = `
                            <span>${contact.name} - ${contact.email} - ${contact.phone}</span>
                            <section>
                                <button onclick="editContact('${contact.id}')">Modify</button>
                                <button onclick="deleteContact('${contact.id}')">Delete</button>
                            </section>
                        `;
                contactsDiv.appendChild(contactDiv);
            });
        });
}

function createContact(contactData) {
    fetch('http://localhost:8080/contacts', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(contactData)
    })
        .then(response => response.json())
        .then(data => {
            loadContacts();
            resetForm();
        });
}

function updateContact(id, contactData) {
    fetch('http://localhost:8080/contacts/' + id, {
        method: 'PUT',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(contactData)
    })
        .then(response => response.json())
        .then(data => {
            loadContacts();
            resetForm();
        });
}

function deleteContact(id) {
    fetch('http://localhost:8080/contacts/id/' + id, { method: 'DELETE' })
        .then(response => {
            loadContacts();
        });
}

function editContact(id) {
    fetch('http://localhost:8080/contacts/' + id)
        .then(response => response.json())
        .then(contact => {
            document.getElementById('contactId').value = contact.id;
            document.getElementById('name').value = contact.name;
            document.getElementById('email').value = contact.email;
            document.getElementById('phone').value = contact.phone;
            document.getElementById('contactFormDiv').classList.remove('hidden');
        });
}

function resetForm() {
    document.getElementById('contactForm').reset();
    document.getElementById('contactId').value = '';
}