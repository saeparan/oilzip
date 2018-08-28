<template>
<div>
  <section class="container">
    <div class="row">
      <div class="int-bg text-center">
        <h1>Swing By</h1>
        <h3>'행성의 중력을 이용해 우주선을 가속시키는 항법'</h3>
        <h4>만약 당신이 과속 아닌 가속이 필요하다면.</h4>
      </div>
  </section>

  <section>
    <div class="container">
      <div class="row">
        <div class="col-lg-12 text-center">
          <h3>프로파일<hr></h3>
        </div>
        <div class="col-lg-12 text-center">
          <p><strong>이름</strong> 박영준</p>
          <p><strong>나이</strong> 29 <small>(2017년 기준)</small></p>
          <p><strong>거주지</strong> 서울</p>
          <p><strong>좌우명</strong> 아, 이거 왜안되지?</p>
          <p><strong>취미</strong> 자전거 라이딩, 자동차 드라이브, 위키읽기, 청소, 세차</p>
          <p><strong>관심사</strong> 웹, 우주, 과학(지구과학, 천문학, 우주 ...), 자동차, 아이돌그룹</p>
          <p><strong>개발을 시작하게 된 계기</strong> 초등학생 때 나모웹에디터로 장난을 치다가 발을 들이게 됨.</p>
        </div>
        <div class="col-lg-12">
          <div class="row">
              <div class="timeline-centered">
                <article class="timeline-entry">
                    <div class="timeline-entry-inner">
                        <div class="timeline-icon bg-success">
                            <i class="entypo-feather"></i>
                        </div>
                        <div class="timeline-label">
                            <h2>1989. 01. 출생</h2>
                            <p>1월의 새벽 청주에서 ?kg의 건강한 사내로 출생하였다.</p>
                        </div>
                    </div>
                </article>
                <article class="timeline-entry">
                    <div class="timeline-entry-inner">
                        <div class="timeline-icon bg-secondary">
                            <i class="entypo-suitcase"></i>
                        </div>
                        <div class="timeline-label">
                            <h2>2011. 02. 대학졸업</h2>
                            <p>
                              초, 중, 고, 군대, 대학까지 사회에 나갈 준비를 완료하였다.<br>
                              대학시절에는 외주를 받고 프로그램과 홈페이지를 만들어 팔아 돈을벌어 등록금을 내었다.
                            </p>
                        </div>
                    </div>
                </article>
                <article class="timeline-entry">
                    <div class="timeline-entry-inner">
                        <div class="timeline-icon bg-info">
                            <i class="entypo-suitcase"></i>
                        </div>
                        <div class="timeline-label">
                            <h2>2011. 02. 첫 회사 입사</h2>
                            <p>졸업을 하기도전에 운좋게 첫 회사에 입사하였는데, 입사하자마자 10만원의 설날 떡값을 받아서 기분이 좋았었다. (과거형)</p>
                        </div>
                    </div>
                </article>
                <article class="timeline-entry">
                    <div class="timeline-entry-inner">
                        <div class="timeline-icon bg-primary">
                            <i class="entypo-suitcase"></i>
                        </div>
                        <div class="timeline-label">
                            <h2>~ 현재. 개발자 생활 중</h2>
                            <p>개발이 재밌다, 무에서 유를 창조하는 중 ...</p>
                        </div>
                    </div>
                </article>
              </div>
        	</div>
        </div>
      </div>
    </div>
  </section>
  <section>
    <div class="container">
      <div class="row">
        <div class="col-lg-12 text-center">
          <h3 class="title">Tech Stack<hr></h3>
        </div>
        <div class="col-lg-12 div-skills">
          <p>
            웹과 관련된 모든 기술을 구사할 수 있으며, 언어는 도구일뿐이기에 모르는 부분은 공부하면 된다.<br>
            최근에는 자바스크립트 분야에 관심이 많다. <small>(Hybrid App, Electron, ES6, Fullstack framework, Promise ...)</small>
          </p>
          <p>
            <span class="label label-success">html</span>
            <span class="label label-danger">css</span>
            <span class="label label-info">javascript</span>
            <span class="label label-warning">node.js</span>
            <span class="label">php</span>
            <span class="label label-success">go</span>
            <span class="label label-danger">python</span>
            <span class="label label-info">mysql</span>
            <span class="label label-warning">mongodb</span>
            <span class="label">Linux(like Ubuntu)</span>
          </p>
        </div>
      </div>
    </div>
  </section>
</div>
</template>

<script>
export default {
  layout: 'default',
  methods: {
    alert(msg, event) {
      event.target.reset()
      if (event) event.preventDefault()
      alert(msg)
    }
  }
}
</script>

<style scoped>
  .div-skills .label {
    margin-left: 3px;
  }
  .int-bg {
    position: relative;
    left: 0px;
    width: 100%;
    background: url('~assets/img/intro.bg.jpg') 50% 50% no-repeat fixed;
    background-size: cover;
    padding-top: 11%;
    margin-bottom: 30px;
  }

  .int-bg h1 {
    font-weight: 500;
    font-fanily: Arial;
    animation: fadein 2s;
    -moz-animation: fadein 2s; /* Firefox */
    -webkit-animation: fadein 2s; /* Safari and Chome */
    -o-animation: fadein 2s; /* Opera */
  }

  .int-bg h1, .int-bg h2, .int-bg h3, .int-bg h4, .int-bg h5{
    color: #fff !important;
    font-weight: 300;
    text-shadow: #333 1px 1px 0;
  }

  div.container h3 {
    font-weight: 600;
    text-align: center;
  }

  @keyframes fadein {
      from {
          opacity:0;
      }
      to {
          opacity:3;
      }
  }
  @-moz-keyframes fadein { /* Firefox */
      from {
          opacity:0;
      }
      to {
          opacity:3;
      }
  }
  @-webkit-keyframes fadein { /* Safari and Chrome */
      from {
          opacity:0;
      }
      to {
          opacity:3;
      }
  }
  @-o-keyframes fadein { /* Opera */
      from {
          opacity:0;
      }
      to {
          opacity: 3;
      }
  }


  img {
      vertical-align: middle;
  }

  .img-responsive {
      display: block;
      height: auto;
      max-width: 100%;
  }

  .img-rounded {
      border-radius: 3px;
  }

  .img-thumbnail {
      background-color: #fff;
      border: 1px solid #ededf0;
      border-radius: 3px;
      display: inline-block;
      height: auto;
      line-height: 1.428571429;
      max-width: 100%;
      moz-transition: all .2s ease-in-out;
      o-transition: all .2s ease-in-out;
      padding: 2px;
      transition: all .2s ease-in-out;
      webkit-transition: all .2s ease-in-out;
  }

  .img-circle {
      border-radius: 50%;
  }

  .timeline-centered {
      position: relative;
      margin-bottom: 30px;
  }

      .timeline-centered:before, .timeline-centered:after {
          content: " ";
          display: table;
      }

      .timeline-centered:after {
          clear: both;
      }

      .timeline-centered:before, .timeline-centered:after {
          content: " ";
          display: table;
      }

      .timeline-centered:after {
          clear: both;
      }

      .timeline-centered:before {
          content: '';
          position: absolute;
          display: block;
          width: 4px;
          background: #f5f5f6;
          /*left: 50%;*/
          top: 20px;
          bottom: 20px;
          margin-left: 30px;
      }

      .timeline-centered .timeline-entry {
          position: relative;
          /*width: 50%;
          float: right;*/
          margin-top: 5px;
          margin-left: 30px;
          margin-bottom: 10px;
          clear: both;
      }

          .timeline-centered .timeline-entry:before, .timeline-centered .timeline-entry:after {
              content: " ";
              display: table;
          }

          .timeline-centered .timeline-entry:after {
              clear: both;
          }

          .timeline-centered .timeline-entry:before, .timeline-centered .timeline-entry:after {
              content: " ";
              display: table;
          }

          .timeline-centered .timeline-entry:after {
              clear: both;
          }

          .timeline-centered .timeline-entry.begin {
              margin-bottom: 0;
          }

          .timeline-centered .timeline-entry.left-aligned {
              float: left;
          }

              .timeline-centered .timeline-entry.left-aligned .timeline-entry-inner {
                  margin-left: 0;
                  margin-right: -18px;
              }

                  .timeline-centered .timeline-entry.left-aligned .timeline-entry-inner .timeline-time {
                      left: auto;
                      right: -100px;
                      text-align: left;
                  }

                  .timeline-centered .timeline-entry.left-aligned .timeline-entry-inner .timeline-icon {
                      float: right;
                  }

                  .timeline-centered .timeline-entry.left-aligned .timeline-entry-inner .timeline-label {
                      margin-left: 0;
                      margin-right: 70px;
                  }

                      .timeline-centered .timeline-entry.left-aligned .timeline-entry-inner .timeline-label:after {
                          left: auto;
                          right: 0;
                          margin-left: 0;
                          margin-right: -9px;
                          -moz-transform: rotate(180deg);
                          -o-transform: rotate(180deg);
                          -webkit-transform: rotate(180deg);
                          -ms-transform: rotate(180deg);
                          transform: rotate(180deg);
                      }

          .timeline-centered .timeline-entry .timeline-entry-inner {
              position: relative;
              margin-left: -20px;
          }

              .timeline-centered .timeline-entry .timeline-entry-inner:before, .timeline-centered .timeline-entry .timeline-entry-inner:after {
                  content: " ";
                  display: table;
              }

              .timeline-centered .timeline-entry .timeline-entry-inner:after {
                  clear: both;
              }

              .timeline-centered .timeline-entry .timeline-entry-inner:before, .timeline-centered .timeline-entry .timeline-entry-inner:after {
                  content: " ";
                  display: table;
              }

              .timeline-centered .timeline-entry .timeline-entry-inner:after {
                  clear: both;
              }

              .timeline-centered .timeline-entry .timeline-entry-inner .timeline-time {
                  position: absolute;
                  left: -100px;
                  text-align: right;
                  padding: 10px;
                  -webkit-box-sizing: border-box;
                  -moz-box-sizing: border-box;
                  box-sizing: border-box;
              }

                  .timeline-centered .timeline-entry .timeline-entry-inner .timeline-time > span {
                      display: block;
                  }

                      .timeline-centered .timeline-entry .timeline-entry-inner .timeline-time > span:first-child {
                          font-size: 15px;
                          font-weight: bold;
                      }

                      .timeline-centered .timeline-entry .timeline-entry-inner .timeline-time > span:last-child {
                          font-size: 12px;
                      }

              .timeline-centered .timeline-entry .timeline-entry-inner .timeline-icon {
                  background: #fff;
                  color: #737881;
                  display: block;
                  width: 40px;
                  height: 40px;
                  -webkit-background-clip: padding-box;
                  -moz-background-clip: padding;
                  background-clip: padding-box;
                  -webkit-border-radius: 20px;
                  -moz-border-radius: 20px;
                  border-radius: 20px;
                  text-align: center;
                  -moz-box-shadow: 0 0 0 5px #f5f5f6;
                  -webkit-box-shadow: 0 0 0 5px #f5f5f6;
                  box-shadow: 0 0 0 5px #f5f5f6;
                  line-height: 40px;
                  font-size: 15px;
                  float: left;
              }

                  .timeline-centered .timeline-entry .timeline-entry-inner .timeline-icon.bg-primary {
                      background-color: #303641;
                      color: #fff;
                  }

                  .timeline-centered .timeline-entry .timeline-entry-inner .timeline-icon.bg-secondary {
                      background-color: #ee4749;
                      color: #fff;
                  }

                  .timeline-centered .timeline-entry .timeline-entry-inner .timeline-icon.bg-success {
                      background-color: #00a651;
                      color: #fff;
                  }

                  .timeline-centered .timeline-entry .timeline-entry-inner .timeline-icon.bg-info {
                      background-color: #21a9e1;
                      color: #fff;
                  }

                  .timeline-centered .timeline-entry .timeline-entry-inner .timeline-icon.bg-warning {
                      background-color: #fad839;
                      color: #fff;
                  }

                  .timeline-centered .timeline-entry .timeline-entry-inner .timeline-icon.bg-danger {
                      background-color: #cc2424;
                      color: #fff;
                  }

              .timeline-centered .timeline-entry .timeline-entry-inner .timeline-label {
                  position: relative;
                  background: #f5f5f6;
                  padding: 1em;
                  margin-left: 60px;
                  -webkit-background-clip: padding-box;
                  -moz-background-clip: padding;
                  background-clip: padding-box;
                  -webkit-border-radius: 3px;
                  -moz-border-radius: 3px;
                  border-radius: 3px;
              }

                  .timeline-centered .timeline-entry .timeline-entry-inner .timeline-label:after {
                      content: '';
                      display: block;
                      position: absolute;
                      width: 0;
                      height: 0;
                      border-style: solid;
                      border-width: 9px 9px 9px 0;
                      border-color: transparent #f5f5f6 transparent transparent;
                      left: 0;
                      top: 10px;
                      margin-left: -9px;
                  }

                  .timeline-centered .timeline-entry .timeline-entry-inner .timeline-label h2, .timeline-centered .timeline-entry .timeline-entry-inner .timeline-label p {
                      color: #000;
                      font-size: 12px;
                      margin: 0;
                      line-height: 1.428571429;
                  }

                      .timeline-centered .timeline-entry .timeline-entry-inner .timeline-label p + p {
                          margin-top: 15px;
                      }

                  .timeline-centered .timeline-entry .timeline-entry-inner .timeline-label h2 {
                      font-size: 16px;
                      margin-bottom: 10px;
                  }

                      .timeline-centered .timeline-entry .timeline-entry-inner .timeline-label h2 a {
                          color: #303641;
                      }

                      .timeline-centered .timeline-entry .timeline-entry-inner .timeline-label h2 span {
                          -webkit-opacity: .6;
                          -moz-opacity: .6;
                          opacity: .6;
                          -ms-filter: alpha(opacity=60);
                          filter: alpha(opacity=60);
                      }

</style>
