import React, { useState, useEffect } from "react";
import { loadStripe } from "@stripe/stripe-js";
import { Elements } from "@stripe/react-stripe-js";
import { BrowserRouter, Routes, Route } from "react-router-dom";

import CheckoutForm from "./CheckoutForm";
import "./App.css";
import UserPaymentDetails from "./UserPaymentDetails";

// Make sure to call loadStripe outside of a component’s render to avoid
// recreating the Stripe object on every render.
// This is a public sample test API key.
// Don’t submit any personally identifiable information in requests made with this key.
// Sign in to see your own test API key embedded in code samples.
const stripePromise = loadStripe("pk_test_51OdxYiSAcOZ5DlUs5VTd0n0sLM9zqA7AiYlwrX3V7MSqJlJ8q3A6HboTgWXBU0tR3E4Z0ROpy0py55HyaVebZlo0003AdPq7UF");

// ... (previous code)

export default function App() {
  const [clientSecret, setClientSecret] = useState("");

  useEffect(() => {
    // Create PaymentIntent as soon as the page loads
    fetch("/create-payment-intent", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ items: [{ id: "xl-tshirt" }] }),
    })
      .then((res) => res.json())
      .then((data) => setClientSecret(data.clientSecret));
  }, []);

  const appearance = {
    theme: 'stripe',
  };
  const options = {
    clientSecret,
    appearance,
  };

  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<CheckoutForm />} />
      </Routes>
    </BrowserRouter>
  );
}
