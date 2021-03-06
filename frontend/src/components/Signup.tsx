import React, { Component } from "react";

interface SignupProps {
  email: string;
  password: string;
  onChangeEmail: (email: string) => void;
  onChangePassword: (password: string) => void;
  onSignup: (email: string, password: string) => void;
}

export default class Signup extends Component<SignupProps> {
  handleChangeEmail = (event: { target: { value: string } }) => {
    this.props.onChangeEmail(event.target.value);
  };

  handleChangePassword = (event: { target: { value: string } }) => {
    this.props.onChangePassword(event.target.value);
  };

  onSignup = () => {
    const { email, password } = this.props;
    this.props.onSignup(email, password);
  };

  render() {
    const { email, password } = this.props;
    return (
      <div style={{ textAlign: "center" }}>
        <table>
          <tbody>
            <tr>
              <th>メールアドレス</th>
              <td>
                <input
                  type="email"
                  value={email}
                  onChange={this.handleChangeEmail}
                />
              </td>
            </tr>
            <tr>
              <th>パスワード</th>
              <td>
                <input
                  type="password"
                  minLength={8}
                  value={password}
                  onChange={this.handleChangePassword}
                />
              </td>
            </tr>
          </tbody>
        </table>
        <button onClick={this.onSignup}>登録</button>
      </div>
    );
  }
}
