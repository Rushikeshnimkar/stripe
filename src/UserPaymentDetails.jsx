// FormPage.js
import React, { useState } from 'react';
import { BrowserRouter as  useNavigate } from 'react-router-dom';

export default function UserPaymentDetails() {
  const [formData, setFormData] = useState({
    Name: '',
    Email: '',
    Address: {
      Line1: '',
      Line2: '',
      City: '',
      State: '',
      PostalCode: '',
      Country: ''
    },
    Amount: '' 
  });

  const history =  useNavigate();

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData((prevData) => ({
      ...prevData,
      [name]: value
    }));
  };

  const handleAddressChange = (e) => {
    const { name, value } = e.target;
    setFormData((prevData) => ({
      ...prevData,
      Address: {
        ...prevData.Address,
        [name]: value
      }
    }));
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    // You can perform any necessary actions with the form data here
    console.log(formData);
    // Redirect to another page
    history.push('/confirmation');
  };

  return (
    <form onSubmit={handleSubmit}>
      <label>
        Name:
        <input type="text" name="Name" value={formData.Name} onChange={handleChange} />
      </label>

      <label>
        Email:
        <input type="email" name="Email" value={formData.Email} onChange={handleChange} />
      </label>

      <label>
        Address Line 1:
        <input type="text" name="Line1" value={formData.Address.Line1} onChange={handleAddressChange} />
      </label>

      <label>
        Address Line 2:
        <input type="text" name="Line2" value={formData.Address.Line2} onChange={handleAddressChange} />
      </label>

      <label>
        City:
        <input type="text" name="City" value={formData.Address.City} onChange={handleAddressChange} />
      </label>

      <label>
        State:
        <input type="text" name="State" value={formData.Address.State} onChange={handleAddressChange} />
      </label>

      <label>
        Postal Code:
        <input type="text" name="PostalCode" value={formData.Address.PostalCode} onChange={handleAddressChange} />
      </label>

      <label>
        Country:
        <input type="text" name="Country" value={formData.Address.Country} onChange={handleAddressChange} />
      </label>

      <label>
        Amount:
        <input type="text" name="Amount" value={formData.Amount} onChange={handleChange} />
      </label>

      <button type="submit">Proceed</button>
    </form>
  );
};
