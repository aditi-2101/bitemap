import React, { useState } from 'react';
import './ReviewForm.css';
import axios from 'axios';

const baseUrl = process.env.REACT_APP_BACKEND_API_URL;

const ReviewForm = ({ onSubmit }) => {
  const [username, setAuthor] = useState('');
  const [rating, setRating] = useState('');
  const [review, setBody] = useState('');

  const handleSubmit = async (e) => {
    e.preventDefault();
    const accessToken = sessionStorage.getItem("access_token");
    
    const postData = {

      res_id: parseInt(sessionStorage.getItem('selectedRestaurantId')),
      //res_id: 63469,
      rating: parseFloat(rating),
      review: review
    };

    try {
      // Send a POST request using Axios
      const response = await axios.post( baseUrl+"/restaurants/review", postData, {
        headers: {
          'Authorization': `Bearer ${accessToken}`
        }
      });
  
      // Handle the response if needed
      console.log('POST request successful:', response.data);
      
      // Reset the form fields if needed
      setAuthor('');
      setRating('');
      setBody('');

      window.location.reload();
      
      // Call the onSubmit prop with the form data
      //onSubmit({ username, rating, review });
    } catch (error) {
      // Handle errors if the POST request fails
      console.error('Error:', error);
    }
  };

  return (
    <form onSubmit={handleSubmit} className="review-form">
      <input type="number" value={rating} onChange={(e) => setRating(e.target.value)} placeholder="Rating (1-5)" required />
      <textarea value={review} onChange={(e) => setBody(e.target.value)} placeholder="Your review" required />
      <button type="submit">Submit</button>
    </form>
  );
};

export default ReviewForm;
