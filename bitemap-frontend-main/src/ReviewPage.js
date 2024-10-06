// ReviewPage.js
import React, { useState, useEffect } from 'react';
import ReviewCard from './ReviewCard';
import ReviewForm from './ReviewForm';
import './ReviewPage.css';
import axios from 'axios';

const baseUrl = process.env.REACT_APP_BACKEND_API_URL;




  const ReviewPage = () => {

    const [reviews, setReviews] = useState([]);
  
    const addReview = (review) => {
      setReviews([...reviews, review]);
    };
  
    useEffect(() => {
      const getReviews = async () => {
        try {
          //let restaurant_id = 63470;

          let accessToken = sessionStorage.getItem("access_token");
          let restaurant_id = sessionStorage.getItem("selectedRestaurantId");
          const response = await axios.get(baseUrl + `/restaurants/review/${restaurant_id}`, {
            headers: { Authorization: `Bearer ${accessToken}` }
          });
          setReviews(response.data.data);
          console.log(response.data);
          console.log("success"); // Assuming your API returns some data
        } catch (error) {
          console.error('Error:', error);
        }
      };
  
      getReviews(); // Call the function inside useEffect
  
    }, []);


    return (
      <div className="review-page">
        <div className="reviews-column">
          <h1>BiteMap Reviews</h1>
          <div className="reviews-container">
            {reviews.map((review, index) => (
              <ReviewCard key={index} review={review} />
            ))}
          </div>
        </div>
        <div className="form-column">
          <h2>Add a Review</h2>
          <ReviewForm onSubmit={addReview} />
        </div>
      </div>
    );
  };
  
  export default ReviewPage;