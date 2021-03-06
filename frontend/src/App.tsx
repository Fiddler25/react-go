import React, { Component } from "react";
import Signup from "./components/Signup";

interface AppState {
  email: string;
  password: string;
}

export default class App extends Component<{}, AppState> {
  constructor(props: {}) {
    super(props);
    this.state = {
      email: "",
      password: "",
    };
  }

  onSignup = async (email: string, password: string) => {
    const endpoint = "//localhost:3000/api/v1/signup";
    const options = {
      method: "POST",
      headers: {
        "Content-Type": "application/json; charset=utf-8",
      },
      body: JSON.stringify({ email, password }),
    };
    const { result } = await fetch(endpoint, options).then((res) => res.json());
    console.log(result);
  };

  handleChangeEmail = (email: string) => {
    this.setState({ email });
  };

  handleChangePassword = (password: string) => {
    this.setState({ password });
  };

  render() {
    const { email, password } = this.state;
    return (
      <div
        style={{
          width: "100vw",
          height: "100vh",
          display: "flex",
          justifyContent: "center",
          alignItems: "center",
        }}
      >
        <Signup
          onSignup={this.onSignup}
          onChangeEmail={this.handleChangeEmail}
          onChangePassword={this.handleChangePassword}
          email={email}
          password={password}
        />
      </div>
    );
  }
}
