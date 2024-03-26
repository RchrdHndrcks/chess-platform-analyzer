import React from 'react';
import ReactDOM from 'react-dom/client';
import Game from './components/board/board.js';

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render
(
  <React.StrictMode>
    <div className="container">
      <div className="row">
        <div className="col">
          <Game />
        </div>
      </div>
    </div>
  </React.StrictMode>
);
