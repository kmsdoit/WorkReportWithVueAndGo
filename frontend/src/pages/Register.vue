<template>
  <!-- <Header></Header> -->
  <div class="register-bg">
    <router-link to="/">
      <img class="logo-img" src="../assets/IDB_logo.png" alt="IDB_LOGO" />
    </router-link>
    <div class="register-container">
      <form @submit.prevent="submitForm">
        <h2>회원가입</h2>
        <!-- ID -->
        <div class="id-box register-box">
          <label for="id">아이디</label>
          <input
            type="text"
            id="id"
            autofocus
            v-model="id"
            @blur="IdValidation()"
          />
          <p class="idNone" id="idMsg">{{ idMsg }}</p>
        </div>
        <!-- PW -->
        <div class="pw-box register-box">
          <label for="pw">비밀번호</label>
          <input type="password" id="pw" v-model="pw" @blur="PwValidation()" />
          <p class="pwNone" id="pwMsg">{{ pwMsg }}</p>
        </div>
        <!-- PW확인 -->
        <div class="pw-check-box register-box">
          <label for="pwCheck">비밀번호 확인</label>
          <input
            type="password"
            id="pwCheck"
            v-model="pwCheck"
            @blur="PwCheckValidation()"
          />
          <p class="pwCheckNone" id="pwCheckMsg">{{ pwCheckMsg }}</p>
        </div>
        <!-- Name -->
        <div class="name-box register-box2">
          <p>이름</p>
          <input type="text" id="name" v-model="name" />
          <p class="nameNone" id="nameMsg">{{ nameMsg }}</p>
        </div>
        <!-- Email -->
        <div class="email-box register-box2">
          <p>이메일</p>
          <input type="email" id="email" v-model="email" />
          <p class="emailNone" id="emailMsg">{{ emailMsg }}</p>
        </div>
        <div style="clear: both"></div>
        <!-- 입사일 -->
        <div class="join-box register-box2">
          <p>입사일</p>
          <input type="text" id="joinCompany" v-model="joinCompany" />
          <p class="joinCompanyNone" id="joinCompanyMsg">
            {{ joinCompanyMsg }}
          </p>
        </div>
        <!-- 직책 -->
        <div class="position-box register-box2">
          <p>직책</p>
          <input type="text" id="position" v-model="position" />
          <p class="positionNone" id="positionMsg">{{ positionMsg }}</p>
        </div>
        <div style="clear: both"></div>
        <!-- 가입버튼 -->
        <button type="submit" class="register-btn">가입하기</button>
      </form>
    </div>
    <!-- </div> -->
  </div>
</template>

<script>
import axios from "axios";
// import Header from '..components/Header.vue';

export default {
  name: "Register",
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
      // 영문 대문자 또는 소문자 또는 숫자로 시작하는 아이디, 길이는 5~15자, 끝날때 제한 없음
      const idReg = /^[A-za-z0-9]{5,15}/g;
      const isValidId = idReg.test(this.id);

      const userId = document.getElementById("id");
      if (this.id.length == 0) {
        userId.classList.add("border-red");
        this.idMsg = "공백입니다.";
      } else if (isValidId !== true) {
        userId.classList.add("border-red");
        this.idMsg = "올바른 이메일 형식이 아닙니다.";
      } else if (isValidId == true) {
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
      } else if (isValidPw !== true) {
        userPw.classList.add("border-red");
        this.pwMsg =
          "비밀번호는 특수문자,문자,숫자를 포함한 형태의 8~15자리 이내로 해주세요.";
      } else if (isValidPw == true) {
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
      } else if (isValidPwCheck !== true) {
        userPwCheck.classList.add("border-red");
        this.pwCheckMsg = "비밀번호가 일치하지 않습니다.";
      } else if (isValidPwCheck == true) {
        userPwCheck.classList.remove("border-red");
        this.pwCheckMsg = "";
      }
    },
    emailValidation() {
      // 이메일은 공백이면 안되고, 이메일 형식으로 작성되어야 한다.
      const Reg =
        /^[0-9a-zA-Z]([-_.]?[0-9a-zA-Z])*@[0-9a-zA-Z]([-_.]?[0-9a-zA-Z])*.[a-zA-Z]{2,3}$/i;
      const isValidEmail = reg.test(this.id);
      console.log(isValidId); // 유효한 id를 입력하면 true가 나와야함
    },

    //^ 가입하기 클릭시 실행할 함수(form 전송시)
    submitForm() {
      // 폼 전송시 id, pw, pwCheck 모두 검사
      this.IdValidation();
      this.PwValidation();
      this.PwCheckValidation();
      // url 적고 post로 데이터 보내기
      axios
        .post("http://localhost:8081/api/user/register")
        .then((result) => {
          console.log(result);
          // if(status == 200) {
          // 		성공시 안의코드 실행
          // }
        })
        .catch((err) => {
          console.log(err);
        });
    },
  },
  // components: { Header }
};
</script>

<style>
div,
input,
form {
  box-sizing: border-box;
  margin: 0;
  padding: 0;
}
.register-bg {
  width: 100vw;
  height: 100vh;
  overflow-x: hidden;
  background-color: #0298db;
  font-size: 16px;
  font-weight: 400;
  line-height: 1.4;
  font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, Oxygen,
    Ubuntu, Cantarell, "Open Sans", "Helvetica Neue", sans-serif;
  position: relative;
}

.register-container {
  background: #fff;
  padding: 30px;
  width: 100%;
  max-width: 600px;
  margin: 0 auto 50px;
  border: none;
  border-radius: 10px;
  position: relative;
}

/* logo-img */
.logo-img {
  display: block;
  margin-left: auto;
  margin-right: auto;
  image-rendering: -webkit-optimize-contrast;
  transform: translateZ(0);
  backface-visibility: hidden;
}

.register-container form h2 {
  font-size: 26px;
  font-weight: 800;
  text-align: center;
}

/* 아이디, 비밀번호, 비밀번호 확인 */
.register-box {
  width: 96%;
  display: block;
  margin: 10px auto;
}

.register-box input,
.register-box2 input {
  width: 100%;
  padding: 10px;
  border: 2px solid #eee;
  border-radius: 5px;
  background-color: #fafafa;
}
.register-box label {
  width: 100%;
  display: block;
  margin: 10px auto;
}

.register-box p {
  width: 100%;
  display: block;
  margin: 8px auto;
  color: tomato;
}

/* 이름, 이메일, 입사일, 직책 */
.register-box2 {
  width: 50%;
  float: left;
  padding: 10px;
}

/* 클래스 추가시 붙이기 */
.register-box .border-red {
  border: none;
  border: 1px solid red;
  outline: none;
}

.join-box,
.position-box {
  margin-bottom: 50px;
}

.register-btn {
  display: block;
  width: 100%;
  max-width: 150px;
  padding: 10px 15px;
  font-size: 20px;
  text-decoration: none;
  letter-spacing: 3px;
  /* margin-top: 50px; */
  color: #fff;
  cursor: pointer;
  transition: 0.5s;
  border: none;
  border-radius: 4px;
  background: #0298db;
  position: absolute;
  right: 40px;
  bottom: 20px;
}
.register-btn:hover {
  background-color: #0276aa;
}
</style>
