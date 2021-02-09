import React from "react";
import Login from "./Login";
import Page from "./Page";

export class App extends React.Component {
    state = {isLoggedIn: false}

    render() {
        if (this.state.isLoggedIn) {
            return <Page onLogOut={this.handleLogOut}/>
        }
        else {
            return <Login onLogIn={this.handleLogIn}/>
        }
    }

    handleLogIn = () => {
        this.setState({isLoggedIn: true})
    }
    handleLogOut = () => {
        this.setState({isLoggedIn: false})
    }
}