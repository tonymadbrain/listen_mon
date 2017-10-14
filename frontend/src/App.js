import React, { Component } from 'react';
import './App.css';
import AppsContainer from './components/AppsContainer'

class App extends Component {
  render() {
    return (
      <div className="App">
        <header className="App-header">
          <h1 className="App-title">Listen monitoing</h1>
        </header>
	<AppsContainer />
      </div>
    );
  }
}

export default App;
