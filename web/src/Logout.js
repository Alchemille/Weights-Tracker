import React from "react";
import {GoogleLogout} from "react-google-login";
import {GOOGLE_CLIENT_ID} from "./config";

export default function Logout(props) {

    return (
        <GoogleLogout
            clientId={GOOGLE_CLIENT_ID}
            buttonText="Logout"
            onLogoutSuccess={props.onLogOut}
        />)
}