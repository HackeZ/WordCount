package main

import "fmt"

func main() {
	page := "某高中同学在网上告诉我小樱领证结婚了，还给我发来他们的结婚照。红底照片上的一对新人好似金童玉女，十分合衬。我盯着小樱光洁的额头看，突然好奇，新郎知道小樱曾经额头上都是毛发吗？你才不是一个没有故事的女同学　　小樱是我的高中舍友。那时我们很要好，手机还不普及，我的床位靠近固定电话机，她经常躺在我的床上用IC卡打电话，巴掌大的小脸，露出满额头的毛发。我第一次看到时吓了一跳，“你的额头怎么都是毛？”她拿开话筒，嘘声说：“没进化好！”原来她毛发旺盛，额头都是毛发，发际线模糊，撩起来怪吓人，平时留着齐刘海，轻易不让人看到额头。就算没进化好，她依然是一个美女。其实她算不上真正意义的美，我们家的男同胞全见过她的人，都说一般，没有我说的那么美。是，严格算起来，她根本不美，小眼睛，单眼皮，塌鼻梁，尤其一口参差不齐的烂牙，实在算不上美女。但她依然被很多人评为班花，这大概和她在深圳长大有关，她小时候和家人在深圳生活上学，因为户籍问题不能在深圳异地高考才回到老家上高中。见过世面的她气质上较之课余还要下田干活的农村女孩自然更出尘优雅一些，而且她的皮肤极好，雪白细腻，招人喜爱。我很喜欢她，她性格活泼热情，声音甜美清脆，我们很要好，几乎形影不离。和她走在路上，常有男生看她，我跟她开玩笑：“和你在一起，我也享受了回头率。”我们去超市买东西，连我的东西老板也因为喜欢她而省去零头。因为喜欢，我待她很好，大早起床去食堂帮她买早餐。当然她也对我不错，怒骂对她不错但欺负过我的男生，帮我洗衣服，和我讲她初中男友的故事。我知道有很多男生喜欢他，但很少人明儿地追她，因为和她走得近的男生大都知道她有个初中同学的男朋友。他们分分合合好几年，整个高中，她都在我的床上和他通电话。他在深圳，她在深圳上初中认识了他。在她的描述里，他是一个清瘦挺拔的男生，气质忧郁，但聪敏有趣。每一段潮湿惨绿的青春期，都有一颗无处盛放的灵魂和一份叛逆不安的爱情。小樱说，初中时他们的恋爱，曾经惊动得老师把家长都叫来。一次，班主任听说有女生跳楼殉情，急不可待，还没有确认真假就把她家长喊来学校，因为她是班里唯一公开恋爱的女生。当然不是她，她说她是自爱的人，再大的事也不会做自伤的事。但经班主任一搅，父母全知道，她虽然认为恋爱天经地义但让父母操心还是很难受。恋爱的问题，亲子的问题，在叛逆的青春期里，成为少女心中一刀又一刀的伤痕，她最不开心的时候，曾一个人在坟地改造的学校后山呆一整夜，也曾将男朋友送的东西扔到大海里。恋爱中的女生自有一种对异性的排他，所以追她的人不多，她和大部分男生是一种哥们儿关系。高一下学期有个班级被拆分，分来一个胆小内向的男生，喜欢上了小樱。因为胆小内向，他倒也没敢有什么行动，但不知为何，他喜欢小樱的事还是被很多男生看出来，大概是因为年少单纯的人很容易真情流露吧。令人讨厌的男生们看到他写满小樱名字的作业本，大肆嘲笑，广而告之。那是一个很怯弱内向的男生，并不起眼，面对自己的心情大概也慌张，话也不敢跟小樱说，在其他人的嘲笑下日渐退缩，独来独往，更见孤单。整整大半个学期，嘲笑声不间断。"

	epage := "There are moments in life when you miss someone so much that you just want to pick them from your dreams and hug them for real! Dream what you want to dream;go where you want to go;be what you want to be,because you have only one life and one chance to do all the things you want to do.May you have enough happiness to make you sweet,enough trials to make you strong,enough sorrow to keep you human,enough hope to make you happy? Always put yourself in others’shoes.If you feel that it hurts you,it probably hurts the other person, too.The happiest of people don’t necessarily have the best of everything;they just make the most of everything that comes along their way.Happiness lies for those who cry,those who hurt, those who have searched,and those who have tried,for only they can appreciate the importance of peoplewho have touched their lives.Love begins with a smile,grows with a kiss and ends with a tear.The brightest future will always be based on a forgotten past, you can’t go on well in lifeuntil you let go of your past failures and heartaches.When you were born,you were crying and everyone around you was smiling.Live your life so that when you die,you're the one who is smiling and everyoene around you is crying.Please send this message to those people who mean something to you,to those who have touched your life in one way or another,to those who make you smile when you really need it,to those that make you see the brighter side of things when you are really down,to those who you want to let them know that you appreciate their friendship.And if you don’t, don’t worry,nothing bad will happen to you,you will just miss out on the opportunity to brighten someone’s day with this message."

	MaxC, MaxW := searchEnglishWord(epage)
	fmt.Println("MaxW in English word --> ", MaxC, MaxW, "\n\n")
	MaxC, MaxW = searchChineseWord(page)
	fmt.Println("MaxW in Chinese word --> ", MaxC, MaxW, "\n\n")

}

func searchChineseWord(page string) (int, []string) {

	countWord := make(map[string]int)
	for _, v := range page {
		// 不是汉字抛弃
		if v < '\u4E00' || v > '\u9FFF' {
			continue
		}
		// 统计汉字信息
		sv := string(v)
		if _, ok := countWord[sv]; !ok {
			countWord[sv] = 1
		}
		countWord[sv]++
	}

	fmt.Println("word -->", countWord)
	return PickMaxValueMap(countWord)

}

func searchEnglishWord(page string) (int, []string) {
	countWord := make(map[string]int)
	var singleWord string
	for _, v := range page {
		// 拼接单字
		sv := string(v)
		if !(sv >= "a" && sv <= "z") || (sv >= "A" && sv <= "Z") {
			singleWord += sv
			continue
		}
		// Map 字符管理
		if _, ok := countWord[singleWord]; !ok {
			countWord[singleWord] = 1
		}
		countWord[singleWord]++
		// 别忘记清空状态，也方便GC
		singleWord = ""
	}

	fmt.Println("eword -->", countWord)

	return PickMaxValueMap(countWord)
}

// PickMaxValueMap ...
func PickMaxValueMap(countWord map[string]int) (int, []string) {
	result := make([]string, len(countWord))
	var maxC int
	for k, v := range countWord {
		if maxC == v {
			result = append(result, k)
		}
		if maxC < v {
			// 清空 Map
			maxC = v
			result = nil
			result = append(result, k)
		}
	}
	return maxC, result
}
