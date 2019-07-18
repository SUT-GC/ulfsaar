var app = new Vue({
    el: '#app',
    data: {
        refreshMT: 5000,
        maxTimeCount: 30,
        addBlogShow: false,
        commitData: {
            columns: [],
            rows: []
        },
        commitSetting: {
            yAxisName: ['commit /次']
        },
        userInfo: [],
        newNickName: "",
        dataUrl: "http://haoyugg.cn:9090/"
    },
    methods: {
        asyncCommitBoard() {
            let getCommitBoardData = this.getCommitBoardData;
            getCommitBoardData();
            // setInterval(function () {
            //     getCommitBoardData()
            // }, this.refreshMT)
        },
        getCommitBoardData() {
            this.$http.get(this.dataUrl).then(function (response) {
                this.fillCommitBoard(response.body)

            }, function (response) {
                console.log('服务异常', response)
            });
        },
        fillCommitBoard(data) {
            let totalTime = [];
            let totalNick = [];
            this.commitData.columns = [];
            this.commitData.rows = [];
            this.userInfo = [];

            for (let t in data) {
                if (totalTime.indexOf(t) < 0) {
                    totalTime.push(t);
                }
                for (let n in data[t]) {
                    if (totalNick.indexOf(n) < 0) {
                        totalNick.push(n);
                    }
                }
            }

            this.commitData.columns = ['Time'];
            for (let index in totalNick) {
                this.commitData.columns.push(totalNick[index]);
            }

            totalTime.sort();
            this.commitData.rows = [];
            for (let i in totalTime) {
                let item = {};
                item['Time'] = timestampToDate(totalTime[i]);
                for (let j in totalNick) {
                    item[totalNick[j]] = data[totalTime[i]][totalNick[j]] === undefined ? 0 : data[totalTime[i]][totalNick[j]];
                }
                this.commitData.rows.push(item);
                this.commitData.rows = sliceArrayLeftEndPoints(this.commitData.rows, this.maxTimeCount)
            }

            for (let index in totalNick) {
                this.userInfo.push({"name": totalNick[index], "url": "https://leetcode-cn.com/u/" + totalNick[index]})
            }
            console.log(totalTime);
        },
        handleShow(index, row) {
            window.open(this.userInfo[index].url, '_blank');

            console.log(index, row);
        },
        handleDelete(index, row) {
            let deleteNickName = this.userInfo[index].name;

            console.log(deleteNickName);

            if (deleteNickName.length <= 0) {
                return
            }

            this.$http.post(this.dataUrl + "unregister", deleteNickName).then(function (response) {
                console.log(response.body)

            }, function (response) {
                console.log('服务异常', response)
            });

            this.getCommitBoardData()
        },

        addNewNick() {
            let newNick = this.newNickName;
            if (newNick.length > 0) {
                this.$http.post(this.dataUrl + "register", newNick).then(function (response) {
                    console.log(response.body)

                }, function (response) {
                    console.log('服务异常', response)
                });

                this.getCommitBoardData()
            }
            this.addBlogShow = false;

        },
        handleCloseAddNewNick() {

        },
        showAddBlog() {
            this.addBlogShow = true
        }
    }
});

app.asyncCommitBoard();