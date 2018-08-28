<template>
  <section class="container">
    <LoadModal :message.sync="modalMessage"></LoadModal>

    <div class="filterContainer">
      <div class="close" v-on:click="closeFilter">&#215;</div>
      <div class="row">
        <!-- <div class="col-sm-12">
          <h5 class="mr-xs">지역검색</h5>
          <div class="btn-group">
            <input type="text" class="form-control input-sm">
          </div>
        </div> -->
        <div class="col-sm-12">
          <h5 class="mr-xs">유종</h5>
          <div class="btn-group">
            <button type="button" value="1" class="btn btn-sm btn-default raised" v-bind:class="fuelBtnClass(1)" v-on:click="fuelTypeChange">휘발유</button>
            <button type="button" value="2" class="btn btn-sm btn-default raised" v-bind:class="fuelBtnClass(2)" v-on:click="fuelTypeChange">고급휘발유</button>
            <button type="button" value="3" class="btn btn-sm btn-default raised" v-bind:class="fuelBtnClass(3)" v-on:click="fuelTypeChange">경유</button>
            <button type="button" value="4" class="btn btn-sm btn-default raised" v-bind:class="fuelBtnClass(4)" v-on:click="fuelTypeChange">LPG</button>
            <button type="button" value="5" class="btn btn-sm btn-default raised" v-bind:class="fuelBtnClass(5)" v-on:click="fuelTypeChange">전기충전소</button>
          </div>
        </div>

        <div class="col-sm-12">
          <h5 class="mr-xs">탐색반경</h5>
          <div class="btn-group">
            <button type="button" value="3" class="btn btn-sm btn-default raised" v-bind:class="distanceBtnClass(3)" v-on:click="distanceChange">3km</button>
            <button type="button" value="5" class="btn btn-sm btn-default raised" v-bind:class="distanceBtnClass(5)" v-on:click="distanceChange">5km</button>
            <button type="button" value="10" class="btn btn-sm btn-default raised" v-bind:class="distanceBtnClass(10)" v-on:click="distanceChange">10km</button>
          </div>
        </div>

        <div class="col-sm-12 checkfilter">
          <div class="pretty outline-success">
            <input v-on:change="washChange" type="checkbox"/>
            <label><i class="fa fa-2 fa-check"></i> 세차가능</label>
          </div>
          <div class="pretty outline-success">
            <input v-on:change="directStoreChange" type="checkbox"/>
            <label><i class="fa fa-2 fa-check"></i> 직영점</label>
          </div>
        </div>
      </div>
    </div>

    <div class="row">
      <div class="mapContainer col-sm-12 col-md-9">
        <div class="btn btn-sm btn-default ne" v-on:click="openFilter"><i class="fa fa-2 fa-filter"> 검색조건</i></div>
        <div id="map"></div>
      </div>
      <div class="col-md-3 list">
        <div>
          <div class="pull-left">
            <small><span>{{stationCount}}</span>개 찾음</small>
          </div>
          <div class="pull-right">
            <div class="btn-group">
              <button name="order" value="1" type="button" class="btn btn-default btn-xs" v-bind:class="orderBtnClass(1)" v-on:click="orderValueChange">가격순</button>
              <button name="order" value="2" type="button" class="btn btn-default btn-xs" v-bind:class="orderBtnClass(2)" v-on:click="orderValueChange">거리순</button>
            </div>
          </div>
          <div class="clearfix mb-sm"></div>

          <ul class="list-group list-stations">
            <li class="list-group-item text-center" v-if="stationCount == 0">
              <h6>해당 위치 주변에 주유소/충전소가 없습니다.</h6>
              <span>
                위치, 탐색반경을 변경해주세요.
              </span>
            </li>
            <li class="list-group-item" v-for="items in stations" :data-id="items.ID" v-on:click="mapFocus(items.ID, $event)">
              <div class="row" v-if="fuelType != 5">
                <div class="col-lg-2 col-md-3 vcenter">
                  <img :src="'/' + items.Vendor + '.png'" class="img-circle" width="40" height="40">
                </div>

                <div class="col-lg-7 col-md-9 pt-xs vcenter">
                  <div>
                    {{items.Name}}
                  </div>
                  <div>
                    <small>{{items.Address}} <span v-if="items.OldAddress">({{items.OldAddress}})</small>
                  </div>
                  <div>
                    <small class="text-primary">약 {{items.Distance}}km</small>
                  </div>
                  <div class="mt-sm">
                    <small>
                      <label class="label label-danger mr-xxs" v-if="items.Self === 1">셀프</label>
                      <label class="label label-primary mr-xxs" v-if="items.Wash === 1">세차</label>
                      <label class="label label-default mr-xxs" v-if="items.Cvs === 1">편의점</label>
                      <label class="label label-default mr-xxs" v-if="items.Garage === 1">정비소</label>
                      <label class="label label-default" v-if="items.Vendor == 'at_ex'">고속도로/휴게소</label>
                    </small>
                  </div>
                </div>

                <span class="col-lg-3 col-md-12 pr-xs vcenter text-center">
                  <h5>{{items.PriceStr}}</h5>
                </span>
              </div>
              <div v-else class="pt-xs">
                <div>
                  {{items.Name}}
                </div>
                <div>
                  <small>
                    <span v-if="items.Address">{{items.Address}}</span>
                    <span v-else-if="items.OldAddress">{{items.OldAddress}}</span>

                    <span v-if="items.Location">({{items.Location}})</span>
                  </small>
                </div>
                <div class="pt-xs pb-xs">
                  <small v-if="items.Restday">
                    <span class="mr-xs"><strong>쉬는날</strong></span>
                    <span>{{items.Restday}}</span>
                  </small>
                  <div></div>
                  <small>
                    <span class="mr-xs"><strong>운영시간</strong></span>
                    <span v-if="items.Starttime">{{items.Starttime}}시 ~ {{items.Endtime}}시</span>
                    <span v-else class="text-danger">24시간</span>
                  </small>
                </div>
                <div>
                  <small>
                    <label class="label label-mint mr-xs" v-if="items.Slow == 'Y'">완속</label>
                    <label class="label label-danger mr-xs" v-if="items.Quick == 'Y'">급속</label>
                  </small>
                </div>
              </div>

              <div class="clearfix">
            </li>
          </ul>
        </div>
      </div>
    </div>
  </section>
</template>

<script>
import axios from 'axios'
import LoadModal from '~components/LoadModal.vue'

export default {
  components: {LoadModal},
  data: function(context) {
    return {
      isDev: context.isDev,
      stations: [],
      fuelType: 1,
      listOrder: 1,
      distance: 5,
      wash: 0,
      initFinish: true,
      stationCount: 0,
      markers: [],
      currentMarker: null,
      highlightMarkerID: null,
      circle: null,
      loadStations: null,
      showStations: null,
      map: null,
      sortStations: null,
      modalMessage: '위치 가져오는 중 ...',
      mapOptions: {
        zoom: 8,
        minZoom: 4,
        maxZoom: 11
      }
    }
  },
  mounted: function() {
    const _this = this
    $('#modal-spinner').modal('hide')

    this.setCircle = (latlng) => {
      if( !_this.circle || _this.circle == null || typeof _this.circle == 'null' ) {
        _this.circle = new naver.maps.Circle({
            map: _this.map,
            center: latlng,
            radius: 5000,
            fillColor: '#3396EA',
            fillOpacity: 0.1
        })
      } else {
        _this.circle.setCenter(latlng)
      }
    }

    this.currentPosMarker = (position) => {
      if( this.currentMarker == null ) {
        this.currentMarker = new naver.maps.Marker({
          position: position,
          map: this.map,
          icon: {
                  content: '<div><img src="/placeholder.png" width="48" height="48" alt="현재 위치"></div>',
                  size: new naver.maps.Size(38, 58),
                  anchor: new naver.maps.Point(19, 58),
              }
        })
      } else {
        this.currentMarker.setPosition(position)
      }
    },

    this.markerUpdate = () => {
      var mapBounds = this.map.getBounds(), marker = null, position = null;
      for( var marker of this.markers ) {
        if( marker === undefined ) {
          continue
        }

        let setMapVal = null,
            position = marker.getPosition()

        if (mapBounds.hasLatLng(position)) {
            setMapVal = this.map
        }

        marker.setMap(setMapVal);
      }
    }

    this.loadStations = (latlng) => {
      if( latlng === undefined ) {
        latlng = _this.map.getCenter()
      }

      _this.highlightMarkerID = null
      _this.setCircle(latlng)
      _this.currentPosMarker(latlng)
      this.listSelected()

      let url = '//'+ apiUrl +'/stations/' + latlng._lat + '/' + latlng._lng + '/' + _this.fuelType + '/' + _this.distance + '/' + _this.wash
      if( _this.fuelType == 5 ) {
        url = '//'+ apiUrl +'/chargers/' + latlng._lat + '/' + latlng._lng + '/' + _this.distance + '/' + _this.wash
      }

      axios.get(url)
      .then(function (resp) {
        _this.showStations(resp.data)
      })
    }

    this.sortStations = () => {
      var iteratees = this.listOrder == 1 ? 'Price' : 'Distance'
      this.stations = _.orderBy(this.stations, [iteratees], ['asc'])
    }

    this.showStations = (data) => {
      for( var marker of this.markers ) {
        if( marker === undefined ) {
          continue
        }
        marker.setMap(null);
      }

      this.markers = []
      this.stations = []

      if( !data || data === null || typeof data == 'null' ) {
        this.stationCount = 0
        return
      }

      this.stationCount = data.length
      this.stations = data

      if( _this.fuelType == 5 ) {
        this.listOrder = 2
      }

      this.sortStations()

      for( var items of data ) {
        let position = new naver.maps.LatLng(items.X, items.Y),
            mapBounds = this.map.getBounds(),
            map = null;

        if( mapBounds.hasLatLng(position) ) {
          map = _this.map
        }

        let badge = '',
            markerContent = '';

        if( _this.fuelType == 5 ) {
          markerContent = '<div class="text">관공서<br>'
          markerContent += '<small class="text-success">완</small>'
          if( items.Quick == 'Y' ) {
            markerContent += ' / <small class="text-danger">급</small>'
          }
          markerContent += '</div>';
        } else {
          badge = '<span class="badge badge-default">'+items.PriceStr+'</span>'
          markerContent = '<img src="/'+items.Vendor+'.png">'
        }

        this.markers[items.ID] =
          new naver.maps.Marker({
            position: position,
            map: map,
            icon: {
                    content:[
                              '<div class="marker" data-id="'+items.ID+'">',
                              badge,
                              markerContent,
                              '</div>'
                            ].join(''),
                    size: new naver.maps.Size(38, 58),
                    anchor: new naver.maps.Point(19, 58),
                }
          })
      }
    }

    this.listFocusReset = () => {
      $('.list-group-item.selected').removeClass('selected')
    }

    this.listSelected = (id) => {
      let $list = $('.list-stations');
      let $target = $list.find("li[data-id='"+id+"']");
      let topOffSet = 0;

      $list.find('.selected').removeClass('selected')

      if( id !== undefined ) {
          topOffSet = ($target.offset().top + $('.list').scrollTop()) - 120
          $target.addClass('selected');
      }

      $('.list').animate({
        scrollTop: topOffSet
      })
    }

    const apiUrl = this.isDev ? process.env.apiUrl_DEV : process.env.apiUrl_PROD
    const apiKey = this.isDev ? process.env.NAVERMapKey_DEV : process.env.NAVERMapKey_PROD
    const scripts = document.createElement('SCRIPT')
    scripts.setAttribute('src', '//openapi.map.naver.com/openapi/v3/maps.js?clientId=' + apiKey)
    scripts.setAttribute('async', '')
    scripts.setAttribute('defer', '')
    document.body.appendChild(scripts)

    setTimeout(function() {
      var map = new naver.maps.Map('map', _this.mapOptions)

      _this.map = map

      function setLocation(pos) {
       let coords = pos.coords
       var location = new naver.maps.LatLng(coords.latitude, coords.longitude)

       map.setCenter(location)
       if( _this.initFinish === true ) {
         _this.initFinish = false
         _this.loadStations(location)

         $('#modal-spinner').modal('hide')
       }
      }

      function err(err) {
        _this.loadStations()

        _this.modalMessage = '위치를 가져올 수 없습니다.'
        setTimeout(function() {
          $('#modal-spinner').modal('hide')
        }, 1000)
      }

      if (navigator.geolocation) {
       navigator.geolocation.getCurrentPosition(setLocation, err, {timeout: 4000})
      }

      naver.maps.Event.addListener(map, 'dragend', function(e) {
        _this.loadStations()
      })
      naver.maps.Event.addListener(map, 'idle', function(e) {
        _this.markerUpdate()
      })

      $('body').on('click', '.marker', function() {
        let id = $(this).data('id');

        _this.mapFocus(id)
        _this.listSelected(id)
      })
    }, 500)
  },
  watch: {
    fuelType: function(val, oldVal){
      this.loadStations()
    }
  },
  methods: {
    mapFocus(markerID, e) {
      let marker = this.markers[markerID]
      let position = marker.getPosition()
      let highlightMarkerID = this.highlightMarkerID
      if( highlightMarkerID !== null ) {
        this.markerToggleEffect( this.markers[highlightMarkerID] )
      }

      if( e !== undefined ) {
        this.listSelected(markerID)
      }

      this.map.setCenter(position)
      if( this.map.zoom < this.mapOptions.maxZoom ) {
        this.map.setZoom(this.mapOptions.maxZoom-1)
      }

      this.markerToggleEffect(marker)
      this.highlightMarkerID = markerID
    },
    markerToggleEffect(marker) {
      let $html = $("<div>" + marker.icon.content + "</div>")
      $html.children().toggleClass('highlight')
      marker.setIcon(
        {
          content: $html.html(),
          size: new naver.maps.Size(38, 58),
          anchor: new naver.maps.Point(19, 58)
        }
      )
    },
    distanceBtnClass: function(val) {
      return this.distance == val ? 'active' : ''
    },
    orderBtnClass: function(val) {
      return this.listOrder == val ? 'active' : ''
    },
    fuelBtnClass: function(val) {
      return this.fuelType == val ? 'active' : ''
    },
    fuelTypeChange: function(e) {
      this.fuelType = e.target.value
      this.sortStations()
    },
    orderValueChange: function(e) {
      this.listOrder = e.target.value
      this.sortStations()
    },
    distanceChange: function(e) {
      this.distance = parseInt(e.target.value)
      this.circle.setRadius( this.distance * 1000 )
      this.loadStations()
    },
    washChange: function() {
      this.wash = !this.wash === false ? 0 : 1
      this.loadStations()
    },
    directStoreChange: function() {
      alert('직영점 찾기 옵션은 현재 준비 중 입니다.')
    },
    openFilter: function() {
      $('.filterContainer').show()
    },
    closeFilter: function() {
      $('.filterContainer').hide()
    }
  }
}
</script>

<style scoped>
th {
  text-align: center;
}

@media(max-width:767px){
  div.filterContainer {
    display: none;
    padding: 0px 20px 10px 20px;
  }

  div.filterContainer .close {
    position:absolute;
    top: 56px;
    right: 10px;
    font-size: 25px;
    z-index: 3;
  }

  div.filterContainer h5 {
    margin-bottom: 5px;
  }

  div.filterContainer .row .col-sm-12 {
    position: relative;
    padding-top: 10px;
  }

  div.filterContainer .checkfilter {
    margin-top: 30px;
  }

  .mapContainer .btn {
    position: absolute;
    top: 5px;
    right: 18px;
    z-index: 1;
  }
}

@media(min-width:768px){
  div.filterContainer {
    padding-left: 6px;
  }

  div.filterContainer .close {
    display: none;
  }

  div.filterContainer h5 {
    display: inline-block;
    font-size: 13px;
    font-weight: normal;
  }

  div.filterContainer .row .col-sm-12 {
    width: auto;
    margin-left: 20px;
  }

  div.filterContainer .pretty i {
    margin-top: 14px;
  }

  .mapContainer .btn {
    display: none;
  }
}

div.filterContainer {
  border-bottom: 1px solid #BDC3C7;
  background-color: #F5F7FA;
}

div.filterContainer .pretty i {
  margin-right: 2px;
  font-size: 17px !important;
}

section.container, section.container .row {
  height: 100%;
  padding-left: 0px;
  padding-right: 0px;
}

section.container {
  overflow:hidden;
}

.mapContainer, #map {
  height: 100%;
}

section.container div .list {
  overflow-y: scroll;
  padding: 10px;
  height: 100%;
  background-color: #fafafa;
  box-shadow: -8px 0 10px -8px #858585, 0 0 0 0 black;
}

section.container div .list > div {
  margin-right: 15px;
}

.col-sm-12 {
  padding-right: 0px;
}

h6 {
  font-weight: bold;
}

li.list-group-item {
  cursor: pointer;
}

li.list-group-item:hover {
  background-color: #ECF0F1;
}

li.list-group-item.selected {
  color: #000;
  background-color: #FCEBED;
}

li.list-group-item label {
  font-weight: normal;
}

.circle{
  border-radius: 200px;
  background-color: #fff;
  padding: 5px;
  text-align: center;
  font-weight: 600;
}
</style>

<style>
  #map .marker {
    position: relative;
    display: inline-block;
    background-color: #FFF;
    padding: 6px;
    border-radius: 50%;
    border: 2px solid #2c3e50;
    color: #FFF;
  }
  #map .marker img {
    width: 30px;
    height: 30px;
  }
  #map .marker:before {
    content: "";
    position: absolute;
    bottom: -11px;
    left: calc(50% - 14px);
    border-style: solid;
    border-width: 14px 14px 0;
    border-color: #2c3e50 transparent;
    display: block;
    width: 0;
    z-index: -1;
  }
  #map .marker.highlight {
    border: 5px solid #CF000F;
  }
  #map .marker.highlight:before {
    border-color: #CF000F transparent;
  }
  #map .marker span.badge {
    position: absolute;
    color: #000;
    top: -27px;
    left: -1px;
  }
  #map .marker div.text {
    text-align: center;
    width: 30px;
    height: 30px;
    font-size: 11px;
    font-weight: 700;
    overflow: hidden;
    color: #000;
  }
  .label.label-mint {
    background-color: #26A65B !important;
  }
  .vcenter {
    display: inline-block;
    vertical-align: middle;
    float: none;
}
</style>
