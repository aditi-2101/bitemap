// ReviewCard.js
import React from 'react';
import StarRating from './StarRating';

const ReviewCard = ({ review }) => {
  return (
    <div className="review-card">
      <div className="review-card-header">
        <div className="review-card-author" style={{"font-weight": "bold", "font-size": "30px"}}>{review.username}</div>
        <div className="review-card-rating">
          <StarRating rating={review.rating} />
        </div>
      </div>
      <div className="review-card-body">{review.review}</div>
    </div>
  );
};

export default ReviewCard;
