/**
 *
 * Auth:Eric Shi
 * Mail:postmaster@pangu.cloud
 * QQ:155122504
 *
 */
(function(factory) {
    'use strict';
    if (typeof require === 'function' && typeof exports === 'object' && typeof module === 'object') {
        factory(module['exports'] || exports);
    } else if (typeof define === 'function' && (define.amd || define.cmd)) {
        define(['exports'], factory);
    } else {
        factory(window['rco'] = {});
    }
}(function(RealClouds) {
    'use strict';

    //ReadClouds plugins
    var rco = typeof RealClouds !== 'undefined' ? RealClouds : {};

    rco.version = '0.0.1';

    rco.navApp = function(elApp) {
        var navApp = new Vue({
            el: elApp,
            methods: {
                logout: function() {
                    location.href = "/logout?_=" + Math.random();
                }
            }
        });
        return navApp;
    };

    rco.loginApp = function(elApp, defaultUsername) {
        var loginApp = new Vue({
            el: elApp,
            data: {
                user: {
                    username: defaultUsername,
                    password: ""
                },
                loginBtnVal: "登陆",
                loginBtnDisabled: false
            },
            methods: {
                onLoginSubmit: function(e) {
                    this.loginBtnVal = "登陆中...";
                    this.loginBtnDisabled = true;

                    var $form = this;

                    $.ajax({
                        url: "/login" + "?_=" + Math.random(),
                        method: "POST",
                        data: $form.user,
                        dataType: "json",
                        cache: false,
                    }).done(function(data) {
                        var json = data;
                        if (typeof(data) != "object") {
                            json = $.parseJSON(data);
                        }
                        if (json.ok) {
                            location.href = json.data;
                        } else {
                            $.amaran({
                                'theme': 'awesome error',
                                'content': {
                                    title: '登陆失败',
                                    message: '请确认您的登陆信息。',
                                    info: 'ronglian.com',
                                    icon: 'fa fa-times-circle-o'
                                },
                                'position': 'bottom right',
                                'inEffect': 'slideBottom'
                            });
                            $form.loginBtnVal = "登陆";
                            $form.loginBtnDisabled = false;
                        }
                    }).fail(function() {
                        $form.loginBtnVal = "登陆";
                        $form.loginBtnDisabled = false;
                    });
                }
            }
        });
        return loginApp;
    };

    rco.msg = function(t, ico, msg, info) {
        $.amaran({
            'theme': 'awesome ' + t,
            'content': {
                title: '系统提示',
                message: msg,
                info: info,
                icon: 'fa ' + ico
            },
            'position': 'bottom right',
            'inEffect': 'slideBottom'
        });
    };

    rco.toJSON = function(data) {
        if (typeof(data) != "object") {
            return JSON.parse(data);
        } else {
            return data;
        }
    };

    rco.toString = function(data) {
        return JSON.stringify(data);
    };

    rco.b64encode = function(data) {
        return Base64.encode(data);
    };

    rco.fileSizeToUnit = function(bytes) {
        if (bytes === 0) return '0 B';
        var k = 1024,
            sizes = ['B', 'KB', 'MB', 'GB', 'TB', 'PB', 'EB', 'ZB', 'YB'],
            i = Math.floor(Math.log(bytes) / Math.log(k));
        return (bytes / Math.pow(k, i)).toPrecision(3) + ' ' + sizes[i];
    }

    rco.defautInit = function(elApp) {
        var resultScrollBottom = function() {
            var scrollHeight = $(resultDiv)[0].scrollHeight;
            var scrollTop = $(resultDiv)[0].scrollTop;
            $(resultDiv).scrollTop(scrollHeight + scrollTop);
        };
        var protocol = (location.protocol === "https:") ? "wss://" : "ws://";
        var sockURL = protocol + location.hostname + ((location.port) ? (":" + location.port) : "") + "/ws/image?_=" + Math.random();
        var sock = new Ws(sockURL);

        var ss = Base64.encode("sfsdfasdffasddafdgdfgsfgh==");

        sock.OnConnect(function() {
            rco.msg('ok', 'fa-check', '与服务器连接成功。', "ronglian.com");
            sock.Emit("loadImages", "ok");
        });

        sock.OnDisconnect(function() {
            rco.msg('error', 'fa-times-circle-o', '与服务器连接成功。', "ronglian.com");
        });

        var initApp = new Vue({
            el: elApp,
            data: {
                images: [],
                imageName: ""
            },
            computed: {
                inVisible: function() {
                    return this.images.length !== 0
                },
                filterImages: function() {
                    var searchName = this.imageName;
                    return this.images.filter(function(image) {
                        if (searchName.length == 0) {
                            return true;
                        } else {
                            return image.name.indexOf(searchName) > -1
                        }
                    });
                }
            },
            methods: {
                copyAddr: function(addr) {
                    var clipboard = new Clipboard(this.$el, {
                        text: function() {
                            return addr;
                        }
                    });
                    clipboard.on('success', function(e) {
                        rco.msg('ok', 'fa-check', '复制成功，请粘贴使用。', addr);
                        clipboard.destroy();
                    });
                    clipboard.on('error', function(e) {
                        rco.msg('error', 'fa-times-circle-o', '复制失败，请手动复制。', addr);
                        clipboard.destroy();
                    });
                },
                viewImage: function(id) {
                    rco.msg('warning', 'fa-info-circle', '功能未启用。', id);
                },
                invalidLink: function(id) {
                    rco.msg('warning', 'fa-info-circle', '功能未启用。', id);
                }
            },
            created: function() {
                var downAddr = location.protocol + location.hostname + ((location.port) ? (":" + location.port) : "")
                sock.On("imagesResult", function(data) {
                    var json = rco.toJSON(data);
                    if (json.ok) {
                        var imageArr = json.data;
                        for (var i in json.data) {
                            var img = imageArr[i];
                            img.size = rco.fileSizeToUnit(img.size);
                            img.path = downAddr + img.path
                            initApp.images.push(img);
                        }
                    }
                });
            }
        });
        return initApp;
    };

    //Jquery plugins
    $.extend({
        particles: function() {
            var particles_conf = {
                particles: {
                    number: { value: 20, density: { enable: true, value_area: 1E3 } },
                    color: { value: "#e1e1e1" },
                    shape: { type: "circle", stroke: { width: 0, color: "#000000" }, polygon: { nb_sides: 5 }, image: { src: "img/github.svg", width: 100, height: 100 } },
                    opacity: { value: .5, random: false, anim: { enable: false, speed: 1, opacity_min: .1, sync: false } },
                    size: {
                        value: 15,
                        random: true,
                        anim: {
                            enable: false,
                            speed: 180,
                            size_min: .1,
                            sync: false
                        }
                    },
                    line_linked: { enable: true, distance: 650, color: "#cfcfcf", opacity: .26, width: 1 },
                    move: { enable: true, speed: 2, direction: "none", random: true, straight: false, out_mode: "out", bounce: true, attract: { enable: false, rotateX: 600, rotateY: 1200 } }
                },
                interactivity: {
                    detect_on: "canvas",
                    events: { onhover: { enable: false, mode: "repulse" }, onclick: { enable: false, mode: "push" }, resize: true },
                    modes: {
                        grab: { distance: 400, line_linked: { opacity: 1 } },
                        bubble: { distance: 400, size: 40, duration: 2, opacity: 8, speed: 3 },
                        repulse: { distance: 200, duration: .4 },
                        push: { particles_nb: 4 },
                        remove: { particles_nb: 2 }
                    }
                },
                retina_detect: true
            };

            if (!!window.HTMLCanvasElement) {
                var particlesDiv = $('<div>', {
                    id: "particles"
                }).css({
                    'position': 'absolute',
                    'top': '0',
                    'z-index': '-1',
                    'width': '100%',
                    'height': '100%'
                }).appendTo("body");
                particlesJS("particles", particles_conf);
            }
        }
    });


    $.fn.loading = function(options) {
        if (options === undefined) {
            options = {};
        }
        var text, ico_class, _safe = this;
        text = options.text === undefined ? "Loading..." : options.text;
        ico_class = options.ico_class === undefined ? "fa fa-refresh fa-spin fa-3x fa-fw" : options.ico_class;
        var overlay = $('<div><i class="' + ico_class + '"></i> ' + text + '</div>');
        _safe.html("");
        _safe.css({
            "position": "relative"
        });
        _safe.append(overlay);
        $(overlay).css({
            "display": "inline-block",
            "position": "absolute",
            "z-index": 99999,
            "top": "40%",
            "left": "45%"
        });
    };

    rco.particles = function() {
        $.particles();
    };
}));