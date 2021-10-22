package biligo

import (
	"encoding/json"
	"strconv"
	"testing"
)

var testBiliClient *BiliClient

func init() {
	testBiliClient = newTestBiliClient()
}
func newTestBiliClient() *BiliClient {
	c, _ := NewBiliClient(&BiliSetting{
		Auth: &CookieAuth{
			DedeUserID:      "",
			SESSDATA:        "",
			BiliJCT:         "",
			DedeUserIDCkMd5: "",
		},
		DebugMode: true,
	})
	return c
}

func TestNewBiliClient(t *testing.T) {

	if _, err := NewBiliClient(&BiliSetting{
		Auth: &CookieAuth{
			DedeUserID: "",
			SESSDATA:   "",
			BiliJCT:    "",
		},
	}); err != nil {
		t.Error(err)
		t.FailNow()
	}
}
func TestBili_GetMe(t *testing.T) {
	me, err := testBiliClient.GetMe()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("id: %d,sign: %s,uname: %s\n", me.MID, me.Sign, me.UName)
}

func TestBiliClient_GetNavInfo(t *testing.T) {
	info, err := testBiliClient.GetNavInfo()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf(
		"level: %d,curExp: %d,nextExp: %d,mobileVerified: %d,money: %.2f,role: %d,pendantID: %d,bcoin: %.2f",
		info.LevelInfo.CurrentLevel,
		info.LevelInfo.CurrentExp,
		info.LevelInfo.NextExp,
		info.MobileVerified,
		info.Money,
		info.Official.Role,
		info.Pendant.PID,
		info.Wallet.BcoinBalance,
	)
}
func TestBiliClient_GetNavStat(t *testing.T) {
	stat, err := testBiliClient.GetNavStat()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("following: %d,follower: %d,dynamic_count: %d", stat.Following, stat.Follower, stat.DynamicCount)
}
func TestBiliClient_GetExpRewardStat(t *testing.T) {
	stat, err := testBiliClient.GetExpRewardStat()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf(
		"login:%t,watch: %t,share: %t,coins: %d,tel: %t,email: %t,identifyCard: %t,safeQuestion: %t",
		stat.Login,
		stat.Watch,
		stat.Share,
		stat.Coins,
		stat.Tel,
		stat.Email,
		stat.IdentifyCard,
		stat.SafeQuestion,
	)
}
func TestBiliClient_GetExpCoinReward(t *testing.T) {
	n, err := testBiliClient.GetExpCoinReward()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("exp: %d", n)
}
func TestBiliClient_GetVipStat(t *testing.T) {
	info, err := testBiliClient.GetVipStat()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf(
		"mid: %d,vipType: %d,vipStatus: %d,vipDueDate: %d,vipPayType: %d,themeType: %d",
		info.MID,
		info.VipType,
		info.VipStatus,
		info.VipDueDate,
		info.VipPayType,
		info.ThemeType,
	)
}
func TestBiliClient_GetAccountSafetyStat(t *testing.T) {
	info, err := testBiliClient.GetAccountSafetyStat()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf(
		"mail: %s,tel: %s,hasPwd: %t,readNameVerified: %t,score: %d,pwdLevel: %d,qqBind: %d",
		info.AccountInfo.HideMail,
		info.AccountInfo.HideTel,
		!info.AccountInfo.UnneededCheck,
		info.AccountInfo.RealnameCertified,
		info.AccountSafe.Score,
		info.AccountSafe.PwdLevel,
		info.AccountSNS.QQBind,
	)
}
func TestBiliClient_GetMsgUnread(t *testing.T) {
	unread, err := testBiliClient.GetMsgUnread()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("like: %d,chat: %d,at: %d,reply: %d,sysMsg: %d,up: %d", unread.Like, unread.Chat, unread.At, unread.Reply, unread.SysMsg, unread.Up)
}
func TestBiliClient_GetRealNameStat(t *testing.T) {
	rn, err := testBiliClient.GetRealNameStat()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("realNameStatus: %t", rn)
}
func TestBiliClient_GetRealNameInfo(t *testing.T) {
	info, err := testBiliClient.GetRealNameInfo()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("status: %d,name: %s,remark: %s,card: %s,type: %d", info.Status, info.Realname, info.Remark, info.Card, info.CardType)
}
func TestBiliClient_GetCoinLogs(t *testing.T) {
	logs, err := testBiliClient.GetCoinLogs()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	for _, log := range logs {
		t.Logf("time: %s,delta: %.2f,reason: %s", log.Time, log.Delta, log.Reason)
	}

}
func TestBiliClient_GetRelationStat(t *testing.T) {
	stat, err := testBiliClient.GetRelationStat(546195)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("mid: %d,following: %d,follower: %d,black: %d,whisper: %d", stat.MID, stat.Following, stat.Follower, stat.Black, stat.Whisper)
}
func TestBiliClient_GetUpStat(t *testing.T) {
	stat, err := testBiliClient.GetUpStat(546195)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("archive: %d,article: %d,likes: %d", stat.Archive.View, stat.Article.View, stat.Likes)
}
func TestBili_SignUpdate(t *testing.T) {
	if err := testBiliClient.SignUpdate("除了自己的无知"); err != nil {
		t.Error(err)
		t.FailNow()
	}
}
func TestBiliClient_SpaceTagsSet(t *testing.T) {
	if err := testBiliClient.SpaceSetTags([]string{"test1", "test1", "test2", "test3", "test4", "test5"}); err != nil {
		t.Error(err)
		t.FailNow()
	}
}
func TestBiliClient_SpaceTagsSet2(t *testing.T) {
	// 测试TAG个数上限
	var tags []string
	for i := 0; i < 54; i++ {
		tags = append(tags, strconv.Itoa(i))
	}
	if err := testBiliClient.SpaceSetTags(tags); err != nil {
		t.Error(err)
		t.FailNow()
	}
}
func TestBiliClient_SpaceNoticeSet(t *testing.T) {
	if err := testBiliClient.SpaceSetNotice("testtesttesttesttest"); err != nil {
		t.Error(err)
		t.FailNow()
	}
}
func TestBili_LikeVideo(t *testing.T) {
	if err := testBiliClient.VideoAddLike(759937808, true); err != nil {
		t.Error(err)
		t.FailNow()
	}
}
func TestBili_VideoIsLiked(t *testing.T) {
	liked, err := testBiliClient.VideoIsLiked(759937808)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("liked: %t", liked)
}
func TestBili_VideoAddCoinsWithLike(t *testing.T) {
	if err := testBiliClient.VideoAddCoins(759937808, 1, true); err != nil {
		t.Error(err)
		t.FailNow()
	}
}
func TestBili_VideoIsAddedCoins(t *testing.T) {
	a, err := testBiliClient.VideoIsAddedCoins(759937808)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("coins: %d", a)
}
func TestBili_VideoFavourAdd(t *testing.T) {
	p, err := testBiliClient.VideoSetFavour(759937808, []int64{492144694, 213103794}, nil)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("prompt: %t", p)
}
func TestBili_VideoFavourDel(t *testing.T) {
	p, err := testBiliClient.VideoSetFavour(759937808, nil, []int64{492144694, 213103794})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("prompt: %t", p)
}
func TestBili_VideoIsFavoured(t *testing.T) {
	favoured, err := testBiliClient.VideoIsFavoured(759937808)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("favoured: %t", favoured)
}
func TestBili_VideoTriple(t *testing.T) {
	like, coin, favour, multiply, err := testBiliClient.VideoTriple(377389267)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("like: %t,coin: %t,favour: %t,multiply: %d", like, coin, favour, multiply)
}
func TestBili_VideoShare(t *testing.T) {
	share, err := testBiliClient.VideoShare(377389267)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("shareNum: %d", share)
}
func TestBiliClient_VideoProgressReport(t *testing.T) {
	if err := testBiliClient.VideoReportProgress(13502509, 66445301, 100); err != nil {
		t.Error(err)
		t.FailNow()
	}
}
func TestBiliClient_VideoTags(t *testing.T) {
	tags, err := testBiliClient.VideoGetTags(759949922)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	for _, tag := range tags {
		t.Logf("name: %s,use: %d,liked: %d,hated: %d,Isatten: %d", tag.TagName, tag.Count.Use, tag.Liked, tag.Hated, tag.IsAtten)
	}
}
func TestBiliClient_VideoTagLike(t *testing.T) {
	if err := testBiliClient.VideoLikeTag(759949922, 16230013); err != nil {
		t.Error(err)
		t.FailNow()
	}
}
func TestBiliClient_VideoTagHate(t *testing.T) {
	if err := testBiliClient.VideoHateTag(759949922, 16230013); err != nil {
		t.Error(err)
		t.FailNow()
	}
}
func TestBiliClient_DanmakuPost(t *testing.T) {
	result, err := testBiliClient.DanmakuPost(1, 292592903, 397011525, "bilitest6", 5000, 16777215, 25, 0, 1)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("dmid: %d", result.Dmid)
}
func TestBiliClient_DanmakuCommandPost(t *testing.T) {
	var d struct {
		Msg string `json:"msg,omitempty"`
	}
	d.Msg = "bili~"
	data, _ := json.Marshal(d)
	result, err := testBiliClient.DanmakuCommandPost(1, 292592903, 397011525, 10000, 1, string(data), 0)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("dmid: %d,json: %d", result.ID, result.Extra)
}
func TestBiliClient_DanmakuRecall(t *testing.T) {
	msg, err := testBiliClient.DanmakuRecall(397011525, 56335865687920640)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("msg: %s", msg)
}
func TestBiliClient_DanmakuGetLikes(t *testing.T) {
	result, err := testBiliClient.DanmakuGetLikes(397011525, []uint64{54109805459813888, 54109892081901568})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	for _, r := range result {
		t.Logf("likes: %d,isLiked: %d,dmid: %s", r.Likes, r.UserLike, r.IDStr)
	}
}
func TestBiliClient_DanmakuLike(t *testing.T) {
	err := testBiliClient.DanmakuLike(397011525, 54109805459813888, 1)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}
func TestBiliClient_DanmakuReport(t *testing.T) {
	err := testBiliClient.DanmakuReport(397011525, 54109805459813888, 11, "无效弹幕")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}
func TestBiliClient_DanmakuEditState(t *testing.T) {
	err := testBiliClient.DanmakuEditState(1, 397011525, []uint64{54109805459813888}, 1)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}
func TestBiliClient_DanmakuEditPool(t *testing.T) {
	err := testBiliClient.DanmakuEditPool(1, 397011525, []uint64{54109805459813888}, 1)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}
func TestBiliClient_DanmakuHistoryIndexGet(t *testing.T) {
	r, err := testBiliClient.DanmakuGetHistoryIndex(1176840, 2020, 5)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Log(r)
}
func TestBiliClient_DanmakuHistoryGet(t *testing.T) {
	r, err := testBiliClient.DanmakuGetHistory(1176840, "2020-05-05")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("num: %d", len(r.Danmaku))
}
func TestBiliClient_ChannelAdd(t *testing.T) {
	cid, err := testBiliClient.ChanAdd("test", "testtest")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("cid: %d", cid)
}
func TestBiliClient_ChannelEdit(t *testing.T) {
	err := testBiliClient.ChanEdit(200444, "test1", "1111")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}
func TestBiliClient_ChannelDel(t *testing.T) {
	err := testBiliClient.ChanDel(200444)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}
func TestBiliClient_MyFavoritesList(t *testing.T) {
	list, err := testBiliClient.FavGetMy()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	for _, info := range list.List {
		t.Logf("mlid: %d,fid: %d,title: %s,count: %d", info.ID, info.FID, info.Title, info.MediaCount)
	}
}
func TestBiliClient_VideoHeartBeat(t *testing.T) {
	err := testBiliClient.VideoHeartBeat(13662970, 126654047, 1000)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}
func TestBiliClient_FavoritesAdd(t *testing.T) {
	d, err := testBiliClient.FavAdd("test", "ttttt", true, "")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("id: %d,title: %s,intro: %s,cover: %s", d.ID, d.Title, d.Intro, d.Cover)
}
func TestBiliClient_FavoritesEdit(t *testing.T) {
	d, err := testBiliClient.FavEdit(1342341894, "test222", "test333", true, "")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("id: %d,title: %s,intro: %s,cover: %s", d.ID, d.Title, d.Intro, d.Cover)
}
func TestBiliClient_FavoritesDel(t *testing.T) {
	err := testBiliClient.FavDel([]int64{1342341894})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}
func TestBiliClient_FavoritesResGet(t *testing.T) {
	d, err := testBiliClient.FavGetRes(504229694)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	for _, r := range d {
		t.Logf("id: %d,type: %d", r.ID, r.Type)
	}
}
func TestBiliClient_FavoritesResCopy(t *testing.T) {
	err := testBiliClient.FavCopyRes(504229694, 492144694, 25422594, []string{"626370388:2", "457308975:2", "90919730:2"})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}
func TestBiliClient_FavoritesResMove(t *testing.T) {
	err := testBiliClient.FavMoveRes(492144694, 213103794, 25422594, []string{"626370388:2", "457308975:2", "90919730:2"})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}
func TestBiliClient_FavoritesResDel(t *testing.T) {
	err := testBiliClient.FavDelRes(213103794, []string{"626370388:2", "457308975:2", "90919730:2"})
	if err != nil {
		t.FailNow()
	}
}
func TestBiliClient_FavoritesResClean(t *testing.T) {
	err := testBiliClient.FavCleanRes(213103794)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}
func TestBiliClient_VideoPlayURLGet(t *testing.T) {
	// flv
	r, err := testBiliClient.VideoGetPlayURL(99999999, 171776208, 112, 128)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("qn: %d,duration: %d", r.Quality, r.TimeLength)
	t.Logf("acceptDesc: %v", r.AcceptDescription)
	for _, u := range r.DURL {
		t.Logf("order: %d,size: %d,url: %s", u.Order, u.Size, u.URL)
	}
}
func TestBiliClient_VideoPlayURLGet2(t *testing.T) {
	// dash
	r, err := testBiliClient.VideoGetPlayURL(717935322, 406422412, 112, 128|16)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("qn: %d,duration: %d", r.Quality, r.TimeLength)
	t.Logf("acceptDesc: %v", r.AcceptDescription)
	for _, v := range r.Dash.Video {
		t.Logf("id: %d,codecs: %s,baseURL: %s", v.ID, v.Codecs, v.BaseURL)
	}
	for _, a := range r.Dash.Audio {
		t.Logf("id: %d,codecs: %s,baseURL: %s", a.ID, a.Codecs, a.BaseURL)
	}
}
func TestBiliClient_EmotePackageAll(t *testing.T) {
	packs, err := testBiliClient.EmotePackGetAll("reply")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	for _, pack := range packs {
		t.Logf("pack id: %d,text: %s,url: %s", pack.ID, pack.Text, pack.URL)
		for i := 0; i < 5; i++ {
			t.Logf("  emote id: %d,text: %s", pack.Emote[i].ID, pack.Emote[i].Text)
		}
	}
}
func TestBiliClient_EmotePackageAll2(t *testing.T) {
	packs, err := testBiliClient.EmotePackGetAll("dynamic")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	for _, pack := range packs {
		t.Logf("pack id: %d,text: %s,url: %s", pack.ID, pack.Text, pack.URL)
		for i := 0; i < 5; i++ {
			t.Logf("  emote id: %d,text: %s", pack.Emote[i].ID, pack.Emote[i].Text)
		}
	}
}
func TestBiliClient_EmotePackageAdd(t *testing.T) {
	err := testBiliClient.EmotePackAdd(1, "reply")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}
func TestBiliClient_AudioMyCollections(t *testing.T) {
	coll, err := testBiliClient.AudioGetMyFavLists(1, 2)
	// coll, err := testBiliClient.AudioGetMyFavLists(1, 5)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("total: %d,pageCount: %d,cur: %d,pageSize: %d", coll.TotalSize, coll.PageCount, coll.CurPage, coll.PageSize)
	for _, a := range coll.Data {
		t.Logf("\tid: %d,uid: %d,uname: %s", a.ID, a.UID, a.Uname)
		t.Logf("\tsid: %d,play: %d,collect: %d,comment: %d,share: %d", a.Statistic.SID, a.Statistic.Play, a.Statistic.Collect, a.Statistic.Comment, a.Statistic.Share)
		t.Logf("\ttitle: %s,ctime: %d,type: %d,published: %d", a.Title, a.Ctime, a.Type, a.Published)
		t.Logf("\tcover: %s", a.Cover)
		for _, id := range a.Sids {
			t.Logf("\t\t%d", id)
		}
	}
}
func TestBiliClient_AudioPlayURLGet(t *testing.T) {
	data, err := testBiliClient.AudioGetPlayURL(2478206, 3)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("auid: %d,type: %d,size: %d", data.SID, data.Type, data.Size)
	t.Logf("timeout: %d,title: %s,cover: %s", data.Timeout, data.Title, data.Cover)
	for _, u := range data.CDNs {
		t.Logf("\turl: %s", u)
	}
	for _, q := range data.Qualities {
		t.Logf(
			"type: %d,size: %d,bps: %s,desc: %s,tag: %s,require: %d,requireDesc: %s",
			q.Type,
			q.Size,
			q.Bps,
			q.Desc,
			q.Tag,
			q.Require,
			q.RequireDesc,
		)
	}
}
func TestBiliClient_AudioIsCollected(t *testing.T) {
	is, err := testBiliClient.AudioIsFavored(2478206)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Log(is)
}
func TestBiliClient_AudioIsCoined(t *testing.T) {
	coin, err := testBiliClient.AudioIsCoined(2478206)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Log(coin)
}
func TestBiliClient_ChargeTrade(t *testing.T) {
	// 没钱，穷，没测试
	r, err := testBiliClient.ChargeTradeCreateBp(20, 293793435, "up", 293793435)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf(
		"mid: %d,upMID: %d,orderNo: %s,bpNum: %s,exp: %d,status: %d,msg: %s",
		r.MID,
		r.UpMID,
		r.OrderNo,
		r.BpNum,
		r.Exp,
		r.Status,
		r.Msg,
	)
}
func TestBiliClient_ChargeTradeQrCodeCreate(t *testing.T) {
	// 空间充电
	qr, err := testBiliClient.ChargeTradeCreateQrCode(2, false, 23215368, "up", 23215368)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("url: %s", qr.QrCodeURL)
	t.Logf("exp: %d", qr.Exp)
	t.Logf("token: %s", qr.QrToken)
}
func TestBiliClient_ChargeTradeQrCodeCreate2(t *testing.T) {
	// 视频充电
	qr, err := testBiliClient.ChargeTradeCreateQrCode(2, true, 480366389, "archive", 378015891)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("url: %s", qr.QrCodeURL)
	t.Logf("exp: %d", qr.Exp)
	t.Logf("token: %s", qr.QrToken)
}
func TestBiliClient_ChargeTradeQrCodeCheck(t *testing.T) {
	stat, err := testBiliClient.ChargeTradeCheckQrCode("efe50b495b864c3e9cf5b74b0ae4c482")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("status: %d,mid: %d,order: %s,token: %s", stat.Status, stat.MID, stat.OrderNo, stat.QrToken)
}