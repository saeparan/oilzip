const webpack = require('webpack');

module.exports = {
  env: {
    baseUrl           : process.env.BASE_URL || 'http://localhost:3000',
    NAVERMapKey_DEV   : '7PdneeK65LEDuy4KOs2k',
    NAVERMapKey_PROD  : 'mjmuMDXe_DbbQM1H9eGA',
    apiUrl_DEV        : 'localhost:1323',
    apiUrl_PROD       : 'api.swingbylab.com'
  },
  /*
  ** Headers of the page
  */
  head: {
    title: '기름집',
    meta: [
      { charset: 'utf-8' },
      { name: 'viewport', content: 'width=device-width, initial-scale=1' },
      { name: 'description', hid: 'description', content: "주유소 가격비교, 주유소 최저가, 주유소 찾기" },
      { property: 'og:type', content: 'website'},
      { property: 'og:title', content: '기름집'},
      { property: 'og:description', content: "주유소 가격비교, 주유소 최저가, 주유소 찾기"},
      { property: 'og:url', content: "https://oilzip.saeparan.com"},
    ],
    link: [
      // { rel: 'stylesheet', href: 'https://fonts.googleapis.com/earlyaccess/notosanskr.css' },
      { rel: 'stylesheet', href: 'https://cdnjs.cloudflare.com/ajax/libs/pretty-checkbox/2.1.0/pretty.min.css' },
      { rel: 'stylesheet', href: 'https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap.min.css' },
      { rel: 'stylesheet', href: 'https://cdnjs.cloudflare.com/ajax/libs/bootflat/2.0.4/css/bootflat.min.css' },
      { rel: 'stylesheet', href: 'https://maxcdn.bootstrapcdn.com/font-awesome/4.7.0/css/font-awesome.min.css' }
    ]
  },

  /*
  ** Global CSS
  */
  css: [
    '~assets/css/main.css',
  ],

  /*
  ** Customize the progress-bar color
  */
  loading: { color: '#FFD200' },
  plugins: [
    '~plugins/ga.js'
  ],
  /*
  ** 추가 플러그인 로드
  */
  build: {
    vendor: ['axios', 'bootstrap', 'jquery', 'lodash'],
    plugins:[
      new webpack.ProvidePlugin({
          jQuery: 'jquery',
          $: 'jquery',
          jquery: 'jquery'
      }),
      new webpack.ProvidePlugin({
          _: 'lodash'
      })
    ]
  },
}
