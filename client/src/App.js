// import logo from './logo.svg';
// import './App.css';

import React from "react"

import BookEdit from "./pages/BookEdit"
import CreateBook from "./pages/CreateBook"
import Dashboard from "./pages/Dashboard"
import LoginPage from "./pages/Login"
import HomePage from "./pages/Home"
import RegisterPage from "./pages/Register"

import {
  BrowserRouter as Router,
  Switch,
  Route,
} from "react-router-dom"; 

function App() {
  return (
    <Router>
      <Switch>
        {/* root route */}
        <Route path="/" exact>
          <HomePage/>
        </Route>
        <Route path="/register">
          <RegisterPage/>
        </Route>
        <Route path="/login">
          <LoginPage/>
        </Route>
        <Route path="/dashboard">
          <Dashboard/>
        </Route>
        <Route path="/books">
          <CreateBook/>
        </Route>
        <Route path="/books/:id">
          <BookEdit/>
        </Route>
      </Switch>
    </Router>
  );
}

export default App;
