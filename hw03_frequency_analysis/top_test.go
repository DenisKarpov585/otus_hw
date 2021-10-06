package hw03frequencyanalysis

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	text1 = `Как видите, он  спускается  по  лестнице  вслед  за  своим
	другом   Кристофером   Робином,   головой   вниз,  пересчитывая
	ступеньки собственным затылком:  бум-бум-бум.  Другого  способа
	сходить  с  лестницы  он  пока  не  знает.  Иногда ему, правда,
		кажется, что можно бы найти какой-то другой способ, если бы  он
	только   мог   на  минутку  перестать  бумкать  и  как  следует
	сосредоточиться. Но увы - сосредоточиться-то ему и некогда.
		Как бы то ни было, вот он уже спустился  и  готов  с  вами
	познакомиться.
	- Винни-Пух. Очень приятно!
		Вас,  вероятно,  удивляет, почему его так странно зовут, а
	если вы знаете английский, то вы удивитесь еще больше.
		Это необыкновенное имя подарил ему Кристофер  Робин.  Надо
	вам  сказать,  что  когда-то Кристофер Робин был знаком с одним
	лебедем на пруду, которого он звал Пухом. Для лебедя  это  было
	очень   подходящее  имя,  потому  что  если  ты  зовешь  лебедя
	громко: "Пу-ух! Пу-ух!"- а он  не  откликается,  то  ты  всегда
	можешь  сделать вид, что ты просто понарошку стрелял; а если ты
	звал его тихо, то все подумают, что ты  просто  подул  себе  на
	нос.  Лебедь  потом  куда-то делся, а имя осталось, и Кристофер
	Робин решил отдать его своему медвежонку, чтобы оно не  пропало
	зря.
		А  Винни - так звали самую лучшую, самую добрую медведицу
	в  зоологическом  саду,  которую  очень-очень  любил  Кристофер
	Робин.  А  она  очень-очень  любила  его. Ее ли назвали Винни в
	честь Пуха, или Пуха назвали в ее честь - теперь уже никто  не
	знает,  даже папа Кристофера Робина. Когда-то он знал, а теперь
	забыл.
		Словом, теперь мишку зовут Винни-Пух, и вы знаете почему.
		Иногда Винни-Пух любит вечерком во что-нибудь поиграть,  а
	иногда,  особенно  когда  папа  дома,  он больше любит тихонько
	посидеть у огня и послушать какую-нибудь интересную сказку.
		В этот вечер...`

	text2 = `cat and dog, one dog,two cats and one man`

	text3 = `Жил на свете маленький цветок. Никто и не знал, что
	он есть на земле. Он рос один на пустыре; коровы и козы
	не  ходили  туда,  и  дети  из  пионерского  лагеря  там
	никогда не играли. На пустыре трава не росла, а лежали
	одни старые серые камни.
	Нечем  было  ему  питаться  в  камне;  капли  дождя,
	упавшие с неба, сходили по верху земли и не проникали
	до его корня, а цветок всё жил и жил и рос помаленьку
	выше.
	Днём  цветок  сторожил  ветер,  а  ночью  росу.
	Он трудился  день  и  ночь,  чтобы  жить  и  не  умереть.
	Он превозмогал  терпеньем  свою  боль  от  голода  и
	усталости. Лишь один раз в сутки цветок радовался: когда
	первый  луч  утреннего  солнца  касался  его  утомлённых
	листьев.
	Цветок,  однако,  не  хотел  жить  печально;  поэтому,
	когда ему бывало совсем горестно, он дремал. Всё же он
	постоянно старался расти.
	В  середине  лета  цветок  распустил
	венчик сверху. До этого он был похож на
	травку, а теперь стал настоящим цветком.
	Венчик у него был составлен из лепестков
	простого  светлого  цвета,  ясного  и
	сильного,  как  у  звезды.  И  как  звезда,  он
	светился живым мерцающим огнём, и его
	видно было даже в тёмную ночь. А когда ветер приходил
	на пустырь, он всегда касался цветка и уносил его запах с
	собою.`

	text4 = `Managing site deliveries is critical to the logistical planning and movement
	of your construction site, poorly managed deliveries can cause significant project
	delay and disruption.

	Munnelly Support Services co-ordinate all site deliveries from scheduling deliveries,
	identifying access routes and safely and securely offloading and storing site materials.
	Our Delivery Management System (DMS) ensures we coordinate all site deliveries in conjunction
	with the project programme, so as to not compromise the projects critical path.
	It enables real-time reservation scheduling, easy delivery dispatch and instantaneous
	delivery management through route optimization.
	
	Our process for managing site deliveries is centred around ensuring our sites are well resourced,
	our surrounding areas are not disrupted, and our projects are delivered on time.
	Our Delivery Management System ensures:
	
		A centralised cloud based system operated with a central administrator is used on the project
		All site deliveries are co-ordinated and assigned
		Project vehicle flow is well managed
		All site plant and equipment is accounted and uploaded into a centralised system
		Project deliveries are aligned and scheduled with the wider project programme
		CLOCS / Lifting Plans / CO2 Emissions all collated
		Disruption to the local area and the project is eliminated
		Comprehensive scheduling and planning for the project

		Email for cooperation: user@gmail.com
		Email for users: user@gmail.com
		Email for suggestions: user@gmail.com user@gmail.com user@gmail.com user@gmail.com

	`
)

func TestTop10(t *testing.T) {
	t.Run("no words in empty string", func(t *testing.T) {
		require.Len(t, Top10(""), 0)
	})

	t.Run("no words in empty string (custom)", func(t *testing.T) {
		require.Len(t, Top10custom(""), 0)
	})

	t.Run("positive test", func(t *testing.T) {
		expected := []string{
			"он",        // 8
			"а",         // 6
			"и",         // 6
			"ты",        // 5
			"что",       // 5
			"-",         // 4
			"Кристофер", // 4
			"если",      // 4
			"не",        // 4
			"то",        // 4
		}
		require.Equal(t, expected, Top10(text1))
	})

	t.Run("positive test (custom)", func(t *testing.T) {
		expectedcstm := []string{
			"а",         // 8
			"он",        // 8
			"и",         // 6
			"ты",        // 5
			"что",       // 5
			"в",         // 4
			"его",       // 4
			"если",      // 4
			"кристофер", // 4
			"не",        // 4
		}
		require.Equal(t, expectedcstm, Top10custom(text1))
	})

	t.Run("less 10 words", func(t *testing.T) {
		expected := []string{
			"and",
			"one",
			"cat",
			"cats",
			"dog,",
			"dog,two",
			"man",
		}
		require.Equal(t, expected, Top10(text2))
	})

	t.Run("less 10 words (custom)", func(t *testing.T) {
		expectedcstm := []string{
			"and",
			"one",
			"cat",
			"cats",
			"dog",
			"dog,two",
			"man",
		}
		require.Equal(t, expectedcstm, Top10custom(text2))
	})
	t.Run("english words + email", func(t *testing.T) {
		expected := []string{
			"and",
			"the",
			"deliveries",
			"site",
			"is",
			"project",
			"user@gmail.com",
			"are",
			"for",
			"Email",
		}
		require.Equal(t, expected, Top10(text4))
	})

	t.Run("english words + email (custom)", func(t *testing.T) {
		expectedcstm := []string{
			"and",
			"deliveries",
			"project",
			"site",
			"the",
			"is",
			"our",
			"user@gmail.com",
			"all",
			"are",
		}
		require.Equal(t, expectedcstm, Top10custom(text4))
	})
	t.Run("more words", func(t *testing.T) {
		expected := []string{
			"и",
			"не",
			"он",
			"на",
			"а",
			"его",
			"цветок",
			"Он",
			"в",
			"когда",
		}
		require.Equal(t, expected, Top10(text3))
	})

	t.Run("more words (custom)", func(t *testing.T) {
		expectedcstm := []string{
			"и",
			"он",
			"не",
			"на",
			"цветок",
			"а",
			"в",
			"его",
			"жил",
			"когда",
		}
		require.Equal(t, expectedcstm, Top10custom(text3))
	})
}
