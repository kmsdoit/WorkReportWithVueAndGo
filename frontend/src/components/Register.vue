<template>
<!-- <Header></Header> -->
  <div class="register-container">
		<!-- <div class="w-30">
			<div>idb관련 이미지나 로고 넣기</div>
		</div> -->
		<!-- <div class="w-70"> -->
    <div class="logo-box">
      <img src="../assets/IDB_logo.png" alt="IDB_LOGO">
    </div>
			<form @submit.prevent="submitForm" class="register-form">
				<h2> 회원가입 </h2>
				<!-- ID -->
				<div class="id-box register-div">
					<label for="id">아이디</label>
					<input type="text" id="id" autofocus v-model="id" @blur="IdValidation()"/>
					<!-- <p class="idNone" id="idMsg">{{idMsg}}</p> -->
				</div>
				<!-- PW -->
				<div class="pw-box register-div">
					<label for="pw">비밀번호</label>
					<input type="password" id="pw" v-model="pw" @blur="PwValidation()"/>
					<p class="pwNone" id="pwMsg">{{pwMsg}}</p>
				</div>
				<!-- PW확인 -->
				<div class="pw-check-box register-div">
					<label for="pwCheck">비밀번호 확인</label>
					<input type="password" id="pwCheck" v-model="pwCheck" @blur="PwCheckValidation()"/>
					<p class="pwCheckNone" id="pwCheckMsg">{{pwCheckMsg}}</p>
				</div>
        <!-- ! 이메일이랑 네임 묶고 -->
        <!-- Name -->
        <div class="name-box register-div">
          <label for="name">이름</label>
          <input type="text" id="name" v-model="name" />
          <p class="nameNone" id="nameMsg">{{nameMsg}}</p>
        </div>
				<!-- Email -->
				<div class="email-box register-div">
					<label for="email">이메일</label>
					<input type="email" id="email" v-model="email" />
					<p class="emailNone" id="emailMsg">{{emailMsg}}</p>
				</div>
        <!-- ! 입사일이랑 직책 묶자 -->
				<!-- 입사일 -->
				<div class="join-box register-div">
					<label for="joinCompany">입사일</label>
					<input type="text" id="joinCompany" v-model="joinCompany"  />
					<p class="joinCompanyNone" id="joinCompanyMsg">{{joinCompanyMsg}}</p>
				</div>
        <!-- 직책 -->
				<div class="position-box register-div">
					<label for="position">직책</label>
					<input type="text" id="position" v-model="position" />
          <p class="positionNone" id="positionMsg">{{positionMsg}}</p>
				</div>
        <!-- 가입버튼 -->
				<button type="submit" class="register-btn">가입하기</button>
			</form>
		</div>
	<!-- </div> -->
</template>

<script>
import axios from 'axios';
import Header from './Header.vue';

export default {
    name: "SignupForm",
    data() {
      return {
        id: "",
        pw: "",
        pwCheck: "",
        email: "",
        name: "",
        joinCompany: "",
        position: "",
          
        idMsg: "",
        pwMsg: "",
        pwCheckMsg: "",
        emailMsg: "",
        nameMsg: "",
        joinCompanyMsg: "",
        positionMsg: "",
      };
    },
    methods: {
        //^ id 유효성 검사 (id 형식으로 변경해야함)
        IdValidation() {
            /* eslint-disable */
            // 아이디는 공백이면 안되고, 이메일 형식으로 작성되어야 한다.
            const reg = /^[0-9a-zA-Z]([-_.]?[0-9a-zA-Z])*@[0-9a-zA-Z]([-_.]?[0-9a-zA-Z])*.[a-zA-Z]{2,3}$/i;
            const isValidId = reg.test(this.id);
            console.log(isValidId); // 유효한 id를 입력하면 true가 나와야함
            const userId = document.getElementById("id");
            if (this.id.length == 0) {
                userId.classList.add("border-red");
                this.idMsg = "공백입니다.";
            }
            else if (isValidId !== true) {
                userId.classList.add("border-red");
                this.idMsg = "올바른 이메일 형식이 아닙니다.";
            }
            else if (isValidId == true) {
                userId.classList.remove("border-red");
                this.idMsg = "";
            }
        },
        //^ 비밀번호 유효성 검사
        PwValidation() {
            // 특수문자 / 문자 / 숫자 포함 형태의 8~15자리 이내의 암호 정규식
            const reg = /^.*(?=^.{8,15}$)(?=.*\d)(?=.*[a-zA-Z])(?=.*[!@#$%^&+=]).*$/;
            const isValidPw = reg.test(this.pw);
            console.log(isValidPw);
            const userPw = document.getElementById("pw");
            if (this.pw.length == 0) {
                userPw.classList.add("border-red");
                this.pwMsg = "공백입니다.";
            }
            else if (isValidPw !== true) {
                userPw.classList.add("border-red");
                this.pwMsg = "비밀번호는 특수문자,문자,숫자를 포함한 형태의 8~15자리 이내로 해주세요.";
            }
            else if (isValidPw == true) {
                userPw.classList.remove("border-red");
                this.pwMsg = "";
            }
        },
        //^ 비밀번호 확인 유효성 검사
        PwCheckValidation() {
            // this.pw와 일치하는지만 확인
            const isValidPwCheck = this.pw === this.pwCheck;
            console.log(isValidPwCheck);
            const userPwCheck = document.getElementById("pwCheck");
            if (this.pwCheck.length == 0) {
                userPwCheck.classList.add("border-red");
                this.pwCheckMsg = "공백입니다.";
            }
            else if (isValidPwCheck !== true) {
                userPwCheck.classList.add("border-red");
                this.pwCheckMsg = "비밀번호가 일치하지 않습니다.";
            }
            else if (isValidPwCheck == true) {
                userPwCheck.classList.remove("border-red");
                this.pwCheckMsg = "";
            }
        },
        //^ 가입하기 클릭시 실행할 함수(form 전송시)
        submitForm() {
            // 폼 전송시 id, pw, pwCheck 모두 검사
            this.IdValidation();
            this.PwValidation();
            this.PwCheckValidation();
            // url 적고 post로 데이터 보내기
            axios.post("http://localhost:8081/api/user/register")
                .then(result => {
                console.log(result);
                // if(status == 200) {
                // 		성공시 안의코드 실행
                // }
            })
                .catch(err => {
                console.log(err);
            });
        },
    },
    components: { Header }
};
</script>

<style>
.register-container {
  width: 100vw;
  height: 100vh;
  overflow: hidden;
  background-image: linear-gradient( rgba(0,0,0,0.3), rgba(0,0,0,0.3) ), url("../assets/architecture.jpg");
  background-size: cover;
  background-repeat: no-repeat;
  background-position: center; 
  color: #fff;
  font-size: 16px;
  font-weight: 400;
  line-height: 1.4;
  font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, Oxygen, Ubuntu, Cantarell, "Open Sans", "Helvetica Neue", sans-serif;
  margin: 0;
  padding: 0;
  box-sizing: border-box;
  text-align: center;
  position: relative;
}

/* 화면 왼쪽 30% */
/* .w-30 {
	width: 30vw;
	height: 100vh;
	display: flex;
	background-color: royalblue;
	position: absolute;
	top: 0;
	left: 0;
} */
/* 화면 오른쪽 70% */
/* .w-70 {
	width: 70vw;
	height: 100vh;
	display: flex;
	justify-content: center;
	align-items: center;
} */

/* logo-box */
.logo-box {
  width: 400px;
  display: flex;
  justify-content: center;
  align-items: center;
  margin: 50px auto 0;
}
.logo-box img {
  width: 200px;
  height: 50px;
  image-rendering: -webkit-optimize-contrast;
  transform: translateZ(0); 
  backface-visibility: hidden;
}

/* form */
.register-form {
	display: block;
  margin: 0px auto;
  width: 500px;
  padding: 40px;
  background: rgba(0, 0, 0, 0.6);
  border-radius: 10px;
  box-shadow: 0 15px 25px rgba(0, 0, 0, 0.6);
}
.register-form h2 {
	margin: 0 0 40px;
	font-size: 30px;
	font-weight: 600;
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