package models

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/coopernurse/gorp"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"reflect"
	"strings"
	"time"
)

//Public System constants
const (
	DbFormat   = "2006-01-02 15:04:05" 	//Database formatter template used by GO from date-n-time 
	JsonFormat = "2006-01-02"			//Database formatter template used by GO from date
)

//Public System variables
var (
	Dbm *gorp.DbMap //Database object public available
)

// Abstracts time.Time struct into JDate
type JDate time.Time

// whenever a type convertion is needed
type CustomTypeConverter struct{}


// It Initializes database connection and sets the connection as public
func init() {
	log.Println("Opening db...")
	var password = os.Getenv("WARNAPASS")
	path := []string{"root:", password, "@(localhost:3306)/warnabroda"}

	db, err := sql.Open("mysql", strings.Join(path, ""))
	checkErr(err, "opening db failed")
	Dbm = &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	Dbm.TypeConverter = CustomTypeConverter{}

	Dbm.AddTableWithName(DefaultStruct{}, "messages").SetKeys(true, "Id")

	Dbm.AddTableWithName(DefaultStruct{}, "contact_types").SetKeys(true, "Id")

	Dbm.AddTableWithName(DefaultStruct{}, "subjects").SetKeys(true, "Id")

	Dbm.AddTableWithName(Warning{}, "warnings").SetKeys(true, "Id")

	Dbm.AddTableWithName(Ignore_List{}, "ignore_list").SetKeys(true, "Id")

	Dbm.AddTableWithName(User{}, "users").SetKeys(true, "Id")

	//Dbm.TraceOn("[gorp]", log.New(os.Stdout, "###Warn A Broda LOG:", log.Lmicroseconds))
	err = Dbm.CreateTablesIfNotExists()
	checkErr(err, "create tables failed")	

	dbLoadDefaultData()

}

func dbLoadDefaultData(){
	_, err := Dbm.Exec("DELETE FROM contact_types")
    checkErr(err, "DELETE FROM contact_types failed")

    _, err = Dbm.Exec("DELETE FROM messages")
    checkErr(err, "DELETE FROM messages failed")

    _, err = Dbm.Exec("DELETE FROM subjects")
    checkErr(err, "DELETE FROM subjects failed")

    _, err = Dbm.Exec("DELETE FROM users")
    checkErr(err, "DELETE FROM users failed")

    _, err = Dbm.Exec("INSERT INTO contact_types VALUES (1,'E-mail','pt-br'),(2,'SMS','pt-br'),(3,'WhatsApp','pt-br')")
    checkErr(err, "INSERT INTO contact_types failed")

    _, err = Dbm.Exec("INSERT INTO messages VALUES (1,'Você está com Mau Hálito','pt-br'),(2,'Você está com odor de suor','pt-br'),(3,'Você tem sujeira nos dentes','pt-br'),(4,'Sua menstruação manchou sua roupa','pt-br'),(5,'O Seu vaso sanitário está sujo de cocô','pt-br'),(7,'Você está fazendo um barulho irritante com a boca','pt-br'),(8,'O barulho dos pés e/ou mãos incomodam','pt-br'),(9,'Você tá com um chulezinho hein','pt-br'),(10,'Tua roupa tá do lado avesso','pt-br'),(11,'Teu cofrinho tá aparecendo','pt-br'),(12,'Você tem caca no nariz','pt-br'),(13,'Mas que mesa bagunçada hein','pt-br'),(14,'Sou um(a) adimirador(a) secreto(a)','pt-br'),(15,'Relaxa, você está muito estressado(a)','pt-br'),(16,'Você está com o ziper aberto','pt-br'),(17,'Teu som tá muito alto. Poderia diminuir por favor?','pt-br'),(18,'Você tá muito linda, tá judiando do meu coração','pt-br'),(19,'Você tá um gato, to de olho hein, se liga','pt-br'),(20,'Your bad breath is noticeable','en'),(21,'Your sweat odor bothers','en'),(22,'You have got something stuck on your teeth','en'),(23,'Your flow leaked in your clothes','en'),(24,'You have some smelly shit stuck in your toilet','en'),(25,'Do you really need to eat so loud?','en'),(26,'Does your hands and/or feet really need to be this loud?','en'),(27,'Your feet are really smelly','en'),(28,'Your clothes are inside out','en'),(29,'Your butt crack is visible','en'),(30,'You have a booger in your nose','en'),(31,'Your desk is a complete mess','en'),(32,'You have a secret admirer, me =D','en'),(33,'Wow, you are too stressed, try to chill out','en'),(34,'Your zipper is down, dont let it run','en'),(35,'Your music/audio is too loud. Please turn it down','en'),(36,'You are so hot you hurt my heart, you are beautiful','en'),(37,'I Cant take my eyes off of you handsome!','en'),(38,'Usted está con mal aliento','es'),(39,'Su olor del sudor me molesta','es'),(40,'¿Tiene algo atorado en los dientes','es'),(41,'Tu menstruación te manchó la ropa','es'),(42,'Usted tiene mierda pegado en su sanitário','es'),(43,'Usted está haciendo ruido molesto con la boca','es'),(44,'Usted está haciendo ruido molesto con los pies u manos','es'),(45,'Usted tiene mal olor en los pies','es'),(46,'Usted está con la ropa al revés','es'),(47,'Tu grieta del extremo está muy visible','es'),(48,'Tu tiene un notable mocos en la nariz','es'),(49,'Tu escritorio está un desastre','es'),(50,'Usted tiene un admirador secreto','es'),(51,'Usted está muy estresado, relajarse','es'),(52,'Tu cremallera está abierta','es'),(53,'Su música/audio es demasiado alto. Por favor, baje el volumen','es'),(54,'Usted es tan caliente que daño a mi corazón, eres muy hermosa','es'),(55,'!No puedo despegar mis ojos de tí, guapo!','es'),(56,'Please respect, no whistle','en'),(57,'Te acho muito metido(a)','pt-br'),(58,'You are so full of yourself','en'),(59,'Eres muy presumido','es'),(60,'Você tem um odiador(a) secreto(a)','pt-br'),(63,'Você não é um bom chefe','pt-br'),(64,'You are not a good boss','en'),(65,'Usted no es un buen jefe','es'),(66,'Não gostei da sua atitude','pt-br'),(67,'I dont like your attitude','en'),(68,'No me gustó tu actitud','es'),(69,'Eu sei que você peidou','pt-br'),(70,'I know you have farted','en'),(71,'Sé que usted ha tirado un pedo','es'),(72,'Alguém lhe recomendou usar o serviço de aviso anônimo','pt-br'),(73,'Someone recommended this anonymous warning service to you','en'),(74,'Álguien te recomienda este servicio de aviso anónimo','es'), (75,'Você é um péssimo motorista, pelo menos aprenda a estacionar direito','pt-br'),(76,'You are such a terrible driver, you should at least learn how to park','en'),(77,'Usted es un pésimo conductor, aprender a aparcar','es')")
    checkErr(err, "INSERT INTO messages failed")

    _, err = Dbm.Exec("INSERT INTO subjects VALUES (1,'Pô parceiro(a)','pt-br'),(2,'Um amigo(a) pediu para avisar','pt-br'),(3,'Só um toque','pt-br'),(4,'Ow, se liga ae','pt-br'),(5,'Olá, venho atraves desse informar','pt-br'),(6,'É para o teu bem','pt-br'),(7,'Quem avisa amigo é','pt-br'),(8,'Não é por mal é só um aviso','pt-br'),(9,'Só estás sendo avisado porque te acham legal','pt-br'),(10,'Só estás sendo avisando porque se importam contigo','pt-br'),(11,'Um amigo acaba de lhe dar um toque','pt-br'),(12,'Los verdaderos amigos te abren los ojos','es'),(13,'Un amigo le advierte suavemente','es'),(15,'Para que lo sepas','es'),(16,'Hey man, escucha me','es'),(17,'Hola, creo que esto puede afectar a usted','es'),(18,'Esto es para su propio bien','es'),(19,'Solo los buenos amigos te abren los ojos','es'),(20,'Yo no quiero ser malo, es sólo un advertirle','es'),(21,'Usted ha sido advertido porque alguien pensará que usted es agradable','es'),(22,'Usted ha sido advertido porque alguien se preocupa por ti','es'),(23,'Un amigo le dio un codazo','es'),(24,'Hey Broda, come on','en'),(25,'A friend gently warns you','en'),(26,'Just so you know','en'),(27,'Hey man, listen up','en'),(28,'Hi, I believe this may concern you','en'),(29,'This is for your own good','en'),(30,'Who warns friend is','en'),(31,'I do not want to be mean, it is just a warn','en'),(32,'You have been warned because someone think you are nice','en'),(33,'You have been warned because someone cares about you','en'),(34,'A friend just poked you','en')")
    checkErr(err, "INSERT INTO subjects failed")

    _, err = Dbm.Exec("INSERT INTO users VALUES (1,'ademarizu','ademarizu@gmail.com','a8df37059a2e9915141c92730242a18a9d85c06c','Ademar Izu','2015-01-02 23:04:46',0,'ROLE_ADMIN'),(2,'hbt.vieira','hbt.vieira@gmail.com','ce9cdc7e751033862c9a453e931cbdce11c7852c','Herbert Smith','2015-01-22 16:28:01',0,'ROLE_ADMIN'),(3,'raissaesther','raissaesther@gmail.com','a8df37059a2e9915141c92730242a18a9d85c06c','Raissa Esther','2015-01-02 23:08:31',0,'ROLE_ADMIN')")
    checkErr(err, "INSERT INTO users failed")
}

func (d JDate) MarshalJSON() ([]byte, error) {
	return json.Marshal((*time.Time)(&d).Format(DbFormat))
}

func (d *JDate) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	t, err := time.Parse(DbFormat, s)
	if err != nil {
		return err
	}
	*d = JDate(t)
	return nil
}

func (me CustomTypeConverter) ToDb(val interface{}) (interface{}, error) {
	switch t := val.(type) {
	case JDate:
		return time.Time(t), nil
	}
	return val, nil
}

func (me CustomTypeConverter) FromDb(target interface{}) (gorp.CustomScanner, bool) {
	switch target.(type) {
	case *JDate:
		binder := func(holder, target interface{}) error {
			// time.Time is returned from db as string
			s, ok := holder.(*string)
			if !ok {
				return errors.New("FromDb: Unable to convert JDate to *string")
			}
			st, ok := target.(*JDate)
			if !ok {
				return errors.New(fmt.Sprint("FromDb: Unable to convert target to *JDate: ", reflect.TypeOf(target)))
			}
			t, _ := time.Parse(DbFormat, *s)
			*st = JDate(t)
			return nil
		}		
		return gorp.CustomScanner{new(string), target, binder}, true
	}
	return gorp.CustomScanner{}, false
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
