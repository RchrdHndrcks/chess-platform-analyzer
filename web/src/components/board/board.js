import React, { useState } from 'react';
import './board.css';
import B from '../../img/pieces/B.png';
import N from '../../img/pieces/N.png';
import Q from '../../img/pieces/Q.png';
import R from '../../img/pieces/R.png';
import K from '../../img/pieces/K.png';
import P from '../../img/pieces/P.png';
import b from '../../img/pieces/bb.png';
import n from '../../img/pieces/bn.png';
import q from '../../img/pieces/bq.png';
import r from '../../img/pieces/br.png';
import k from '../../img/pieces/bk.png';
import p from '../../img/pieces/bp.png';

const pieceImages = {
  "P": P,
  "R": R,
  "B": B,
  "N": N,
  "Q": Q,
  "K": K,
  "p": p,
  "r": r,
  "b": b,
  "n": n,
  "q": q,
  "k": k
};

function Square({ value, coor, onDragStart, onDragOver, onDrop }) {
  const Piece = pieceImages[value];
  return (
    <div
      className="square"
      onDragStart={onDragStart}
      onDragOver={onDragOver}
      onDrop={onDrop}
      draggable={Piece ? true : false}
    >
      {Piece && <img src={Piece} alt={value} className="piece" />}
      <div className="square-coordinate">{coor}</div>
    </div>
  );
}

function generateInitialPosition() {
  const squares = Array(64).fill(null);
  const initialPiecesOrder = ["R", "N", "B", "Q", "K", "B", "N", "R"];

  initialPiecesOrder.forEach((piece, index) => {
    squares[index] = piece.toLowerCase();
  })
  for(let i = 8; i <= 15; i++) {
    squares[i] = "p"
  }

  for (let i = 48; i <= 55; i++) {
    squares[i] = "P";
  }
  initialPiecesOrder.forEach((piece, index) => {
    squares[56 + index] = piece;
  });

  return squares;
}

function Board({ squares, onPieceDragStart, onPieceDrop }) {
  const columns = ['a', 'b', 'c', 'd', 'e', 'f', 'g', 'h'];

  const renderBoard = () => {
    const board = [];
    for (let i = 0; i < 8; i++) {
      const row = [];
      for (let j = 0; j < 8; j++) {
        const c = columns[j] + (8 - i);
        const index = i * 8 + j;
        row.push(
          <Square
            key={index}
            value={squares[index]}
            coor={c}
            onDragStart={(e) => onPieceDragStart(e, squares[index], index)}
            onDragOver={(e) => e.preventDefault()}
            onDrop={(e) => onPieceDrop(e, index)}
          />
        );
      }
      board.push(<div key={i} className="board-row">{row}</div>);
    }
    return board;
  };

  return <div className="board">{renderBoard()}</div>;
}

export default function Game() {
  const [squares, setSquares] = useState(generateInitialPosition());
  const [draggedPiece, setDraggedPiece] = useState(null);

  const handlePieceDragStart = (e, piece, index) => {
    setDraggedPiece({ piece, index });
  };

  const handlePieceDrop = (e, dropIndex) => {
    e.preventDefault();
    const newSquares = [...squares];
    const { piece, index } = draggedPiece;

    // Check if it's a valid move and update the board accordingly
    // Implement your logic for move validation here

    // For example, you can swap pieces for simplicity
    newSquares[index] = null;
    newSquares[dropIndex] = piece;

    setSquares(newSquares);
    setDraggedPiece(null);
  };

  return (
    <div className="game">
      <div className="game-board">
        <Board
          squares={squares}
          onPieceDragStart={handlePieceDragStart}
          onPieceDrop={handlePieceDrop}
        />
      </div>
    </div>
  );
}
