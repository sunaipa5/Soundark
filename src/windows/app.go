package main

import (
	"context"
	"strings"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

}

func (a *App) ondomready(ctx context.Context) {
	js := `
	document.addEventListener('DOMContentLoaded', function() {
		let style = '
html{color-scheme:dark}*::-webkit-scrollbar{width:0 !important}*::-webkit-scrollbar-thumb{display:none}*{border-color:#27272a !important}body{background-color:#000 !important}#app{background-color:#000 !important;color:#999 !important}.dashbox,.frontHero__content,.l-footer,.l-front-mobile-teaser,.l-front-osc-teaser,.l-front-signup-teaser,.mobileApps{display:none !important}.frontHero__container{background-color:#000;max-height:70px}.frontHero{background-color:#000;max-height:70px}.modalWhiteout{background-color:rgba(0, 0, 0, 0.6) !important}.g-z-index-modal-content{border-radius:10px !important}.l-front-footer{margin-top:15vh}.commentForm__inputWrapper input,.headerSearch input,.textfield input{background-color:#111 !important;color:#999 !important}.frontHero{border-top:none !important}.playControls{margin-top:100px !important}.playControls__inner{backdrop-filter: blur(10px) !important;background-color:rgba(0, 0, 0, 0.6) !important;min-height:80px}.mixedSelectionModule{border-radius:10px;padding:10px 20px;border:1px solid #27272a !important}.header__inner,.header__logo{background-color:transparent !important}.header{backdrop-filter: blur(10px) !important;background-color:rgba(0, 0, 0, 0.6) !important;border-bottom:1px solid #27272a}.header__navMenu .selected{background-color:rgba(25, 25, 25, 0.6) !important}.l-main{border:none !important}.sidebarModule{border:1px solid #27272a !important;padding:10px;border-radius:10px !important}.sc-artwork{border-radius:10px !important}.soundBadge__actions,.soundBadge__additional{background-color:#000 !important;background:none !important}.moreActions,.moreActions__group,.queue,.sc-button{background-color:transparent !important;border:none !important}.dialog,.dropdownContent__container,.dropdownMenu,.dropdownMenu,.g-z-index-header-menu,.gritter-item-wrapper,.modal__modal,.playControls__queue,.readMoreTile__countWrapper,.sound__soundActions,.streamContext,.volume__sliderWrapper{backdrop-filter: blur(10px) !important;background-color:rgba(0, 0, 0, 0.6) !important;border-radius:10px !important;border:1px solid #27272a !important;color:#fff !important}.playControl,.sc-button,.sc-button-like{filter: grayscale(100%) brightness(400%)}.sc-button-selected{filter: none}.commentForm__wrapper,.commentItem__timestampLink,.dropbar__content,.headerMenu,.header__left div a:active,.header__left div a:visited,.listenContent__inner,.listenContext,.m-active,.playControls__bg,.queue__itemWrapper,.queue__itemsHeight,.queue__scrollable,.uiEvoButton{background:transparent !important}.m-active,.m-active:hover,.m-dimmed:hover{background-color:rgba(25, 25, 25, 0.6) !important}.queueItemView__details *{color:#fff !important}.playbackTimeline__progressBackground,.playbackTimeline__progressBar{height:4px !important;border-radius:10px}.playbackTimeline__progressHandle{min-width:12px;min-height:12px}.streamSidebar{min-width:322px !important}.headerMenu__list a:hover,.trackList__item .active{background-color:#111 !important}.trackList__item .hover{background-color:#222 !important}#gritter-notice-wrapper,.trackItem__additional{background:transparent !important}.l-about-top{border-radius:10px;padding:10px 20px;border:1px solid #27272a !important}.listenEngagemen,.profileUploadFooter{background-color:transparent !important;border:none !important}.gritter-image{margin:10px}.gritter-item-wrapper{min-height:65px !important;color:#fff !important}.gritter-with-image *{padding:5px 7px}.headerMenu{border-radius:10px !important;overflow:hidden !important;border:none !important}.tileGallery__sliderButton{background-color:#151515 !important;color:red !important;border-radius:100% !important;border:1px solid #27272a !important;filter: none;filter: grayscale(100%) brightness(200%)}.compactTrackList{overflow:hidden !important;border-radius:10px !important}.soundActions__purchaseLink{color:#fff !important}.streamContext{padding:10px}.headerMenu__list a,.l-nav a,.volume__button{color:#999 !important;filter: grayscale(100%) brightness(400%);border:none !important}.l-nav .active{filter: none !important;color:#ff5500 !important}.l-nav{border-bottom:1px solid #27272a}.collectionNav{margin:auto;display:table;border:none !important}.commentItem__content *{color:#fff !important}.fullListenHero,.profileHeaderBackground,.visuals__container{border-radius:0 0 10px 10px !important;overflow:hidden !important}.sc-button-play{filter: none !important;border:1px solid #fff !important}.g-opacity-transition-500{border-radius:10px !important}.soundBadge__actions{background-color:#111 !important;border-radius:10px !important;border:1px solid #27272a}.audibleTilePlaceholder::before{border:1px solid #27272a !important;border-radius:10px !important}.dropdownMenu a{background-color:transparent !important;background:transparent !important}.otFlat{backdrop-filter: blur(10px) !important;background-color:rgba(0, 0, 0, 0.6) !important}#onetrust-policy *,.playableTile__description *{color:#fff !important}#onetrust-button-group button{margin:5px !important;width:214px !important}.webAuthContainerWrapper iframe{filter: invert(100%)}.sc-border-box{border-radius:10px !important;border:1px solid #27272a !important}.header__left div a{border:none !important}.dropbar{backdrop-filter: blur(10px) !important;background-color:rgba(0, 0, 0, 0.6) !important;border:1px solid #27272a !important;border-radius:0 0 10px 10px}.dropbar *{border:none !important}.g-modal-section,.tabs__tabs{background:none !important}.addToPlaylistTabs,.createPlaylistSuggestion__addContainer{background:none !important}.modal__modal *:not(.active){color:#fff !important}.modal__modal a{border:none !important}.modal__modal::before,.searchOptions__scrollableInner a{filter: grayscale(100%) brightness(400%)}.searchTitle{backdrop-filter: blur(10px) !important;background-color:rgba(0, 0, 0, 0.6) !important}.searchItem__trackItem,.sound__artwork,.sound__content{margin:10px !important;overflow:hidden !important}.searchItem__trackItem .visuals__container{border-radius:10px !important}.searchOptions__scrollableInner *{color:#fff !important;border-radius:10px !important}.searchOptions__scrollableInner:after{display:none !important}.createPlaylistSuggestion__add{filter: none !important;background-color:#ff5500 !important;border-radius:10px !important}.createPlaylistSounds,.createPlaylistSuggestions,.streamContext{background-color:#111;border-radius:10px !important;overflow:hidden !important}.createPlaylistSuggestions{border:1px solid #27272a !important;padding:10px !important}.createPlaylistSuggestions .createPlaylistSuggestions__item{border:none !important}.createPlaylistSuggestions .createPlaylistSuggestions__item:not(:first-child):not(:last-child){border-top:1px solid #27272a !important;border-bottom:1px solid #27272a !important}.playControls__elements button{filter: grayscale(100%) brightness(400%)}.g-tabs-link:not(.active),.playbackTimeline__duration{color:#fff !important}.g-tabs-link{border:none !important}.queueItemView{border:none !important;border-radius:0 !important;border-bottom:1px solid #27272a !important}.uploadMain div{background:transparent !important;border-radius:10px !important;color:#fff}.compactTrackListItem:hover,.compactTrackList__moreLink:hover{background:#222 !important}.compactTrackListItem *{color:#fff !important}.compactTrackListItem__additional{background:transparent !important}.header__logoLink{background-color:transparent !important}.sc-truncate{color:#fff !important}.streamSidebar,.userSidebar{height:80vh !important;overflow:scroll !important;padding-right:10px !important}
	'; 

	
	var head = document.head || document.getElementsByTagName('head')[0];

	var styleElement = document.createElement('style');
	styleElement.type = 'text/css';
	styleElement.appendChild(document.createTextNode(style));
	head.appendChild(styleElement);

	if (window.location.href == "https://soundcloud.com/") {



setTimeout(function() {
	var buttonWithTitle = document.querySelector('button[title="Sign in"]');
	if (buttonWithTitle) {
		buttonWithTitle.click();
	}
}, 1200); 

			 
	
	}


	});
		
    `
	modified := strings.ReplaceAll(js, "'", "`")
	runtime.WindowExecJS(ctx, modified)

}
