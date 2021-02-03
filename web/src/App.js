import React from "react";
import Oidc from "./Oidc";
import Page from "./Page";

export class App extends React.Component {
    state = {isLoggedIn: false}

    render() {
        if (this.state.isLoggedIn) {
            return <Page/>
        }
        else {
            return <Oidc onLogIn={this.handleLogIn}/>
        }
    }

    handleLogIn = () => {
        this.setState({isLoggedIn: true})
        console.log(this.state.isLoggedIn)
    }
}