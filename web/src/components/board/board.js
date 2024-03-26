import React, { useState } from 'react';
import './board.css';

function Square({ value, coor }) {
    return (
        <div className="square">
            {value}
            <div className="square-coordinate">{coor}</div>
        </div>
    );
}

function Board({ squares }) {
  const renderBoard = () => {
    const columns = ['a', 'b', 'c', 'd', 'e', 'f', 'g', 'h']
    const board = [];
    for (let i = 0; i < 8; i++) {
      const row = [];
      for (let j = 0; j < 8; j++) {
        const c = columns[j] + (8 - i)
        const index = i * 8 + j;
        row.push(<Square key={index} value={squares[index]} coor={c} />);
      }
      board.push(<div key={i} className="board-row">{row}</div>);
    }
    return board;
  };

  return <div className="board">{renderBoard()}</div>;
}

export default function Game() {
  const [history, setHistory] = useState([Array(64).fill(null)]);
  const currentSquares = history[history.length - 1];

  return (
    <div className="game">
      <div className="game-board">
        <Board squares={currentSquares} />
      </div>
    </div>
  );
}
