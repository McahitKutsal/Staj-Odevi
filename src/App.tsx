import './App.css';
import Message from "./components/Message";
import {BrowserRouter, Route} from "react-router-dom";


function App() {

    //localhost 3000 portuna tarayıcıdan direkt istek atıldığında Message Component'i gösterilecek
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
