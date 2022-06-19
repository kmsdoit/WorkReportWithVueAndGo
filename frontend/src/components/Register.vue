<template>
  <div class="register-container">
		<div class="w-30">
			<div>idb관련 이미지나 로고 넣기</div>
		</div>
		<div class="w-70">
			<form @submit.prevent="submitForm" class="register-form">
				<h2> 회원가입 </h2>
				<div class="email-box register-div">
					<label for="id">이메일</label>
					<input type="text" id="id" v-model="id" autofocus @blur="checkIdValidation"/>
					<p class="idNone" id="idMsg">{{idMsg}}</p>
				</div>
				<div class="pw-box register-div">
					<label for="pw">비밀번호</label>
					<input type="password" id="pw" v-model="pw" @blur="checkPwValidation"/>
					<p class="pwNone" id="pwMsg">{{pwMsg}}</p>
				</div>
				<div class="pw-check-box register-div">
					<label for="pwCheck">비밀번호 확인</label>
					<input type="password" id="pwCheck" v-model="pwCheck" @blur="checkPwCheckValidation"/>
					<p class="pwCheckNone" id="pwCheckMsg">{{pwCheckMsg}}</p>
				</div>
				<div class="name-box register-div">
					<label for="name">이름</label>
					<input type="text" id="name" v-model="name"  />
				</div>
				<!-- <div>
					<label for="position">직책</label>
					<input type="text" id="position" v-model="position" />
				</div> -->
				<button type="submit" class="register-btn">가입하기</button>
		</form>
	</div>
		
	</div>
</template>

<script>
import axios from 'axios';

export default {
	name: 'SignupForm',
	data() {
		return {
			id: '',
			pw: '',
			pwCheck: '',
			name: '',
			idMsg: '',
			pwMsg: '',
			pwCheckMsg: '',

		};
	},
	methods: {
		// id 유효성 검사
		checkIdValidation() {
			/* eslint-disable */ 
			// 아이디는 공백이면 안되고, 이메일 형식으로 작성되어야 한다.
			const reg = /^[0-9a-zA-Z]([-_.]?[0-9a-zA-Z])*@[0-9a-zA-Z]([-_.]?[0-9a-zA-Z])*.[a-zA-Z]{2,3}$/i;
			const isValidId = reg.test(this.id);
			console.log(isValidId);		// 유효한 id를 입력하면 true가 나와야함

			const userId = document.getElementById('id');


			if (this.id.length == 0) {
				userId.classList.add('border-red');
				this.idMsg = '공백입니다.';
			}
			else if (isValidId !== true) {
				userId.classList.add('border-red');
				this.idMsg = '올바른 이메일 형식이 아닙니다.';
			}
			else if (isValidId == true) {
				userId.classList.remove('border-red');
				this.idMsg = '';
			}
		},

		// 비밀번호 유효성 검사
		checkPwValidation() {
			// 특수문자 / 문자 / 숫자 포함 형태의 8~15자리 이내의 암호 정규식
			const reg = /^.*(?=^.{8,15}$)(?=.*\d)(?=.*[a-zA-Z])(?=.*[!@#$%^&+=]).*$/;
			const isValidPw = reg.test(this.pw);
			console.log(isValidPw);

			const userPw = document.getElementById('pw');


			if (this.pw.length == 0) {
				userPw.classList.add('border-red');
				this.pwMsg = '공백입니다.';
			}
			else if (isValidPw !== true) {
				userPw.classList.add('border-red');
				this.pwMsg = '비밀번호는 특수문자,문자,숫자를 포함한 형태의 8~15자리 이내로 해주세요.';
			}
			else if (isValidPw == true) {
				userPw.classList.remove('border-red');
				this.pwMsg = '';
			}
		},

		// 비밀번호 확인 유효성 검사
		checkPwCheckValidation() {
			// this.pw와 일치하는지만 확인
			const isValidPwCheck = this.pw === this.pwCheck;
			console.log(isValidPwCheck);

			const userPwCheck = document.getElementById('pwCheck');


			if (this.pwCheck.length == 0) {
				userPwCheck.classList.add('border-red');
				this.pwCheckMsg = '공백입니다.';
			}
			else if (isValidPwCheck !== true) {
				userPwCheck.classList.add('border-red');
				this.pwCheckMsg = '비밀번호가 일치하지 않습니다.';
			}
			else if (isValidPwCheck == true) {
				userPwCheck.classList.remove('border-red');
				this.pwCheckMsg = '';
			}
		},
		// 가입하기 클릭시 실행할 함수(form 전송시)
		submitForm() {
			// 폼 전송시 id, pw, pwCheck 모두 검사
			this.checkIdValidation();
			this.checkPwValidation();
			this.checkPwCheckValidation();


			// 
			axios.post('')
				.then(result => { 
					console.log(result);



				})
				.catch(err => { 
					console.log(err);



				})

			
		},
	},
};
</script>

<style>
* {
	box-sizing: border-box;
	font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, Oxygen, Ubuntu, Cantarell, "Open Sans", "Helvetica Neue", sans-serif;
	font-size: 16px;
	font-weight: 400; 
  line-height: 1.4; 
	margin: 0;
  padding: 0;
}
.register-container {
  width: 100vw;
  height: 100vh;
  overflow-x: hidden;
  overflow-y: auto;
	position: relative;

  /* background: linear-gradient(#141e30, #243b55); */
  /* color: white; */
  
}
.w-30 {
	width: 30vw;
	height: 100vh;
	display: flex;
	background-color: royalblue;
	position: absolute;
	top: 0;
	left: 0;
}
.w-70 {
	width: 70vw;
	height: 100vh;
	display: flex;
	/* background-color: orange; */
	position: absolute;
	top: 0;
	right: 0;
}
.register-form {
	display: block;
  margin: 150px auto;
  width: 600px;
  padding: 40px;
  background: rgba(0, 0, 0, 0.5); 
}
.register-form h2 {
	margin-bottom: 60px;
	font-size: 30px;
	font-weight: 600;
	color: #fff;
}
.register-div {
	/* width: 100%; */
	margin: 5px auto 60px;
	position: relative;
}

.register-div input {
	width: 100%;
	display: block;
	/* margin-bottom: 60px; */
	padding: 5px 0;
	border: none;
  border-bottom: 1px solid #fff;
  outline: none;
  background: transparent;
}
.register-div label {
	position: absolute;
  top: -30px;
  left: 0;
	/* padding: 5px; */
  font-size: 20px;
  color: #fff;
  pointer-events: none;
  transition: 0.5s;
}
.register-div p {
	text-align: left;
	margin-top: 5px;
	color: tomato;
}

/* 클래스 추가시 붙이기 */
.register-div .border-red {
	border: none;
	border-bottom: 1px solid red;
	outline: none;
}
/* .idNone, .pwNone, .pwCheckNone {
	display: none;
} */

.register-btn {
  display: inline-block;
  width: 100%;
  padding: 10px 15px;
  font-size: 20px;
  text-decoration: none;
  letter-spacing: 3px;
  margin-top: -20px;
	color: #fff;
  cursor: pointer;
  transition: 0.5s;
  border: 1px solid #fff;
	border-radius: 2px;
  background: transparent;
}
.register-btn:hover {
	color: #333;
  border-radius: 5px;
  background-color: #c8c8c8;
}
</style>