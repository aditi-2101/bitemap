import React from 'react';
import { Button } from '@mui/material';
import { useNavigate } from 'react-router-dom';

const DashboardUrl = process.env.REACT_APP_DASHBOARD_URL;

const Home = () => {
  const navigate = useNavigate();

  const goToAnalyticsDashboard = () => {
    window.open(DashboardUrl);
  };

  const goToRestaurantMap = () => {
    navigate('/map');
  };

  return (
    <div style={{ textAlign: 'center' }}>
      <h1>Choose an Option</h1>
      <Button variant="contained" onClick={goToAnalyticsDashboard} style={{ marginBottom: '20px' }}>
        Analytics Dashboard
      </Button>
      <br />
      <Button variant="contained" onClick={goToRestaurantMap}>
        Restaurant Map
      </Button>
    </div>
  );
};

export default Home;
