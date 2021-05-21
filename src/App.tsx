import React, {useEffect, useState} from 'react';
import './App.css';
import Message from "./components/Message";
import {BrowserRouter, Route} from "react-router-dom";


function App() {


    return (
        <div className="App">
            <BrowserRouter>
            <main className="form-signin">
                
                    <Route path="/" component={Message}/>
            </main>
            </BrowserRouter>
        </div>
    );
}

export default App;
