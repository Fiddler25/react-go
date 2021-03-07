import React, { Component } from "react";
import TextField from "@material-ui/core/TextField";
import Button from "@material-ui/core/Button";

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
                <TextField
                  label="email"
                  margin="dense"
                  required
                  size="small"
                  type="email"
                  value={email}
                  onChange={this.handleChangeEmail}
                />
              </td>
            </tr>
            <tr>
              <th>パスワード</th>
              <td>
                <TextField
                  autoComplete="current-password"
                  inputProps={{ minLength: 8 }}
                  label="password"
                  margin="dense"
                  required
                  size="small"
                  type="password"
                  value={password}
                  onChange={this.handleChangePassword}
                />
              </td>
            </tr>
          </tbody>
        </table>
        <Button color="inherit" variant="outlined" onClick={this.onSignup}>
          登録
        </Button>
      </div>
    );
  }
}
