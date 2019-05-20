import React from 'react'
import { Form, Icon, Input, Button, Checkbox, message } from 'antd'
import Link from 'umi/link';
import styles from './style.scss'

import api_sign from '@/services/sign'
import { SignUpReq } from '@/types/sign'
import router from 'umi/router'

class SignUp extends React.Component<any, any> {
	constructor(props: any) {
		super(props)
		this.state = {}
	}

	handleSubmit = (e: any) => {
		e.preventDefault()
		this.props.form.validateFields(async (err: any, values: SignUpReq) => {
			if (!err) {
				try {
					const res = await api_sign.signInByEmail(values)
					message.success('登录成功')
					router.push('/home')
					console.log('ressss', res)
				} catch (e) {
					console.log('error code:', e.code)
				}
			}
		})
	}

	render() {
		const { getFieldDecorator } = this.props.form
		return (
			<Form onSubmit={this.handleSubmit} className={styles['login-form']}>
				<Form.Item>
					{getFieldDecorator('email', {
						rules: [{ required: true, message: '请输入邮箱' }],
					})(
						<Input
							prefix={<Icon type="user" style={{ color: 'rgba(0,0,0,.25)' }} />}
							placeholder="用户名"
						/>,
					)}
				</Form.Item>
				<Form.Item>
					{getFieldDecorator('password', {
						rules: [{ required: true, message: '请输入密码' }],
					})(
						<Input
							prefix={<Icon type="lock" style={{ color: 'rgba(0,0,0,.25)' }} />}
							type="password"
							placeholder="密码"
						/>,
					)}
				</Form.Item>
				<Form.Item>
					{getFieldDecorator('remember', {
						valuePropName: 'checked',
						initialValue: true,
					})(<Checkbox>记住密码</Checkbox>)}
					<Link className={styles['login-form-forgot']} to="">
						忘记密码
					</Link>
					<Button type="primary" htmlType="submit" className={styles['login-form-button']}>
						登录
					</Button>
					<Link to="/sign/signup">立即注册</Link>
				</Form.Item>
			</Form>
		)
	}
}

const WrappedSignUpForm = Form.create({ name: 'normal_login' })(SignUp);

export default WrappedSignUpForm