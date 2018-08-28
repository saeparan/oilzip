<template>
  <nav class="navbar navbar-default navbar-fixed-top navbar-custom affix-top">
    <div class="container">
      <div class="navbar-header page-scroll">
        <span class="navbar-toggle" v-on:click="toggleNavigation">
          <i class="fa fa-bars fa-5" aria-hidden="true"></i>
        </span>
        <nuxt-link to="/" class="navbar-brand">기름집</nuxt-link>
      </div>

      <div class="collapse navbar-collapse" v-bind:class="{ in: show }">
          <ul class="nav navbar-nav navbar-right">
            <li class="page-scroll" v-on:click="toggleNavigation">
                <nuxt-link to="/">주유소찾기</nuxt-link>
            </li>
            <li class="page-scroll">
                <a v-on:click="develMsg">내주유소</a>
            </li>
            <li class="page-scroll">
                <a v-on:click="develMsg">도움말</a>
            </li>
          </ul>
      </div>
    </div>
  </nav>
</template>

<script>
  export default {
    data: function() {
      return {
        show: false
      }
    },
    methods: {
      develMsg() {
        alert("해당 메뉴는 준비 중 입니다.")
      },
      toggleNavigation() {
        this.show = !this.show;
      }
    }
  }
</script>

<style>
.navbar-custom {
  background: #FFF;
  font-family: "Montserrat", "Helvetica Neue", Helvetica, Arial, sans-serif;
  text-transform: uppercase;
  font-weight: 500;
  border-bottom: 1px solid #BDC3C7;
}
.navbar-custom a:focus {
  outline: none;
}
.navbar-custom .navbar-brand {
  color: #ff5959;
  background: url('~assets/img/ci.png') no-repeat;
  background-size: 34px 47px;
  padding: 18px 48px;
}
.navbar-custom .navbar-brand:hover,
.navbar-custom .navbar-brand:focus,
.navbar-custom .navbar-brand:active,
.navbar-custom .navbar-brand.active {
  color: #ff5959;
}
.navbar-custom .navbar-nav {
  letter-spacing: 1px;
}
.navbar-custom .navbar-nav li a {
  color: #000;
}
.navbar-custom .navbar-nav li a:hover {
  color: #000;
  outline: none;
}
.navbar-custom .navbar-nav li.active a {
  color: white;
  background: #53DCFF;
}
.navbar-custom .navbar-nav li.active a:hover,
.navbar-custom .navbar-nav li.active a:focus,
.navbar-custom .navbar-nav li.active a:active {
  background: #53DCFF;
}
.navbar-default .navbar-collapse, .navbar-default .navbar-form {
  border-color: #000;
}
.navbar-custom .navbar-toggle {
  border-color: transparent;
  background-color: transparent;
  font-size: 20px;
  margin-top: 12px;
  padding: 0px;
  color: #434A54;
}
.navbar-custom .navbar-toggle:hover,
.navbar-custom .navbar-toggle:focus {
  border-color: transparent;
  background-color: transparent;
  color: #434A54;
  font-weight: normal;
}
@media(max-width:767px){
  .navbar-custom .navbar-brand {
    background-size: 29px 42px;
    background-position: 5px 3px;
    padding: 18px 48px;
  }
}
@media (min-width: 768px) {
  .navbar-custom {
    background: #FFF;
    height: 80px;
    padding: 15px 0;
    -webkit-transition: padding 0.3s;
    -moz-transition: padding 0.3s;
    transition: padding 0.3s;
  }
  .navbar-custom .navbar-brand {
    font-size: 1.7em;
    -webkit-transition: all 0.3s;
    -moz-transition: all 0.3s;
    transition: all 0.3s;
  }
  .navbar-custom.affix {
    padding: 10px 0;
  }
  .navbar-custom.affix .navbar-brand {
    font-size: 1.5em;
  }
}
</style>
