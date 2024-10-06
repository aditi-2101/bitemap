// StarRating.js
import React from 'react';
import './StarRating.css';

const StarRating = ({ rating }) => {
  const fullStar = '★';
  const emptyStar = '☆';

  // Calculate the number of full and empty stars
  const fullStars = Math.floor(rating);
  const hasHalfStar = rating % 1 !== 0;
  const emptyStars = 5 - fullStars - (hasHalfStar ? 1 : 0);

  return (
    <div className="star-rating">
      {[...Array(fullStars)].map((_, index) => (
        <span key={index} className="star-full">{fullStar}</span>
      ))}
      {hasHalfStar && <span className="star-half">½</span>}
      {[...Array(emptyStars)].map((_, index) => (
        <span key={index} className="star-empty">{emptyStar}</span>
      ))}
    </div>
  );
};

export default StarRating;
