import API from "./api";
import {GoogleLogin} from 'react-google-login';
import {Component} from "react";
import {GOOGLE_CLIENT_ID} from "./config";


class Login extends Component {

    onSignIn = (googleUser) => {
        const profile = googleUser.getBasicProfile();
        console.log('ID: ' + profile.getId()); // Do not send to your backend! Use an ID token instead.
        console.log('Name: ' + profile.getName());
        console.log('Image URL: ' + profile.getImageUrl());
        console.log('Email: ' + profile.getEmail()); // This is null if the 'email' scope is not present.
        console.log(googleUser, profile)

        const id_token = googleUser.getAuthResponse().id_token;

        API.defaults.headers["authorization"] = "bearer " + id_token;

        this.props.onLogIn();
    }

    render() {
        return (
            <GoogleLogin
                clientId={GOOGLE_CLIENT_ID}
                buttonText="Login"
                onSuccess={this.onSignIn}
                isSignedIn
            />
        )

    }
}

export default Login;