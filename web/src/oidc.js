import API from "./api";

window.onSignIn = function (googleUser) {
    const profile = googleUser.getBasicProfile();
    console.log('ID: ' + profile.getId()); // Do not send to your backend! Use an ID token instead.
    console.log('Name: ' + profile.getName());
    console.log('Image URL: ' + profile.getImageUrl());
    console.log('Email: ' + profile.getEmail()); // This is null if the 'email' scope is not present.
    console.log(googleUser, profile)

    const id_token = googleUser.getAuthResponse().id_token;

    API.defaults.headers["authorization"] = "bearer " + id_token;
}

function Oidc() {

    return (
        <div>
            <div className="g-signin2" data-onsuccess="onSignIn"/>
        </div>
    );
}

export default Oidc;