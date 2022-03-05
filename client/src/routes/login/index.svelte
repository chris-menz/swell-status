<script lang="ts">
    import Nav from "$lib/components/nav.svelte"
    import axios from "axios";
    import { user, attemptLogin } from "$lib/stores/userStore"
    import { savedSpots } from "$lib/stores/savedSpotStore";
    import { onMount } from "svelte"
    import { goto } from "$app/navigation";
    import { fly } from "svelte/transition";
    import { axiosConfig } from "$lib/utils/axiosConfig"
    import { endpoint } from "$lib/utils/endpoint"

    let username: string;
    let password: string;

    let newUsername: string;
    let newPassword: string;
    let first_name: string;
    let last_name: string;
    let email: string;

    let displayInvalidCreds: boolean;
    let displaySuccessfulRegister: boolean;
    let displayEmailTaken: boolean;
    let displayUsernameTaken: boolean;
    let displayBadPassword: boolean;
    let displayMissingFields: boolean;

    onMount(async () => {
        displayInvalidCreds = false;
        displaySuccessfulRegister = false;
        displayEmailTaken = false;
        displayUsernameTaken = false;
        displayBadPassword = false;
        displayMissingFields = false;
    })

    async function login(username, password){
        let credentials = {
            username,
            password
        }
        const loginResponseMessage = await attemptLogin(credentials)

        if(loginResponseMessage === "login successful"){
            goto("/surf-reports")
        } else if (loginResponseMessage === "No user found with those credentials"){
            displayInvalidCreds = true; 
            password = ""
        }    
    }

    async function attemptRegister(newUsername, newPassword, first_name, last_name, email){
        displayEmailTaken = false;
        displayUsernameTaken = false;
        displayBadPassword = false;
        displayMissingFields = false;
        
        let credentials = {
            username: newUsername,
            hashed_password: newPassword,
            first_name,
            last_name,
            email
        }
        try {
           const response = await axios.post(endpoint + "/createuser", credentials, axiosConfig) 
           if (response.data.message == "User registered successfully") {
                displaySuccessfulRegister = true;
                login(newUsername, newPassword);
            }
            if (response.data.message.toLowerCase() == "email already registered"){displayEmailTaken = true}
            if (response.data.message.toLowerCase() == "username already registered"){displayUsernameTaken = true}
            if (response.data.message.toLowerCase() == "invalid password"){displayBadPassword}
            if (response.data.message.toLowerCase() == "missing parameters"){displayMissingFields = true}
        } catch (error) {
            console.log(error)
        }
    }
</script>

<main>
    <Nav />
    <div class="container" in:fly="{{ delay: 400, duration: 200 }}">
        <div class="log-in-container">
            <div class="log-in-header">
                Log In
            </div>
            <div class="log-in-body">
                <form>
                    <div class="text-field">
                        <label for="un">Username</label>
                        <input type="text" bind:value={username} id="un">
                    </div>
                    <div class="text-field">
                        <label for="pw">Password</label>
                        <input type="password" bind:value={password} id="pw">
                    </div>
                </form>
                <button on:click={() => {
                    if(!$user){
                        login(username, password)
                    }
                }}>
                    Login
                </button>
                {#if $user}
                    <div class="err" in:fly="{{ delay: 400, duration: 200 }}">You are already logged in</div>
                {/if}
                {#if displayInvalidCreds}
                    <div class="err">Username or password invalid</div>
                {/if} 
            </div>
        </div>
        <div class="sign-up-container">
            <div class="sign-up-header">
                Sign Up
            </div>
            <div class="sign-up-body">
                <div class="text-field">
                    <label for="em">Email</label>
                    <input type="text" bind:value={email} id="em">
                </div>
                <div class="text-field">
                    <label for="un">Username</label>
                    <input type="text" bind:value={newUsername} id="un">
                </div>
                <div class="text-field">
                    <label for="pw">Password (at least five characters)</label>
                    <input type="password" bind:value={newPassword} id="pw">
                </div>
                <div class="text-field">
                    <label for="un">First Name</label>
                    <input type="text" bind:value={first_name} id="fn">
                </div>
                <div class="text-field">
                    <label for="ln">Last Name</label>
                    <input type="text" bind:value={last_name} id="ln">
                </div>
                <button on:click={() => attemptRegister(newUsername, newPassword, first_name, last_name, email)}>
                    Register
                </button>
                {#if displayUsernameTaken}
                    <div class="err">Username aready registered</div>
                {/if}
                {#if displayEmailTaken}
                    <div class="err">Email aready registered</div>
                {/if}
                {#if displayBadPassword}
                    <div class="err">Password must be more than five characters</div>
                {/if}
                {#if displayMissingFields}
                    <div class="err">Must enter all fields</div>
                {/if}
                
            </div>
        </div>
    </div>
</main>

<style>
    *, *::before, *::after{
        box-sizing: border-box;
        margin: 0;
        padding: 0;
    }

    :global(body) {
        background-color: #161616;
        transition: background-color 0.3s
    }

    .container {
        margin-top: 5em;
        display: flex;
        flex-direction: row;
        justify-content: space-evenly;
        align-items: flex-start;
    }

    .log-in-container, .sign-up-container {
        width: 370px;
    }

    .log-in-header, .sign-up-header {
        color: white;
        background-color: #313131;
        box-shadow: rgba(0, 0, 0, 0.3) 0px 4px 12px;
        text-align: center;
        padding: 0.5em 1em;
        font-size: 3em;
        font-family: Verdana, Geneva, Tahoma, sans-serif;
        margin-bottom: 0.5em;
    }

    .log-in-body, .sign-up-body {
        padding: 1em;
        background-color: #313131;
        box-shadow: rgba(0, 0, 0, 0.3) 0px 4px 12px;
    }

    label {
        color: white;
        margin-bottom: 0.25em;
        font-family: Helvetica, sans-serif;
        font-weight: lighter;
        font-size: 1.25em;
    }

    .text-field{
        display: flex;
        flex-direction: column;
        padding-bottom: 1em;
    }
    
    .err {
        margin-top: 00.75em;
        color: white;
        font-family: Helvetica, sans-serif;
        font-weight: lighter;
    }

    input {
        color: white;
        background: #111111;
        border: 0.5px solid grey;
        border-radius: 5px;
        font-size: 1.25em;
        padding: 00.25em;
        transition-duration: 300ms;
    }

    input:defined {
        border: none;
        outline: none;
        /* background: #030303; */
    }

    input:focus {
        background: #030303;
    }

    input:hover {
        background: #030303;
    }

    button {
        position: relative;
        background: linear-gradient(to right bottom, #6a37c2, #49329e);
        color: rgb(240, 234, 234);
        padding: 0.5em 1em;
        min-height: 5vh;
        width: 40%;
        text-align: center;
        font-size: 1.2em;
        font-family: Verdana, sans-serif;
        border: none;
        border-radius: 5px;
        cursor: pointer;
        box-shadow: rgba(0, 0, 0, 0.3) 0px 4px 12px;
        transition-duration: 300ms;
        z-index: 1;
    }

    button::before{
		position: absolute;
		content: "";
		top: 0;
		right: 0;
		bottom: 0;
		left: 0;
		border-radius: 5px;
		background: linear-gradient(to right bottom, rgb(59, 31, 110),rgb(57, 39, 122)0);
		transition: opacity 0.3s;
		z-index: -1;
		opacity: 0;
	}

    button:hover::before {
		opacity: 1;
	}

    button:hover {
        color: rgb(182, 182, 182);
    }

    input:selected {
        border: none;
        outline: none;
    }

    @media (max-width: 800px) {
        .container{
            flex-direction: column;
            justify-content: center;
            align-items: center;
        }

        .log-in-container, .sign-up-container{
            width: 80%;
            margin-bottom: 1.5em;
        }
    }

    @media (max-width: 575px) {
        .log-in-container, .sign-up-container{
            width: 90%;
        }
    }
</style>