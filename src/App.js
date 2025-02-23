import React from 'react';
import Query from './services/hello';
import User from './services/query';
import Count from './services/count';
import './App.css'

function App() {
    return (
        <div>
            <h1>Microservices Interface</h1>
                <Query />
                <User  />
                <Count />
        </div>
    );
};

export default App;