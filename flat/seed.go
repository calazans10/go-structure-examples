package main

import (
	"log"
	"time"

	"github.com/jmoiron/sqlx"
)

// Seed defines the seed structure
type Seed struct {
	DB *sqlx.DB
}

// NewSeed creates a new seed
func NewSeed(db *sqlx.DB) *Seed {
	return &Seed{
		DB: db,
	}
}

// Pollute prepares database with seed data
func (s *Seed) Pollute() {
	s.createTables()
	s.clearTables()
	s.populateMovies()
	s.populateReviews()
}

func (s *Seed) createTables() {
	var schema = `
	CREATE TABLE IF NOT EXISTS movie (
		id TEXT,
		title TEXT NOT NULL,
		release_year integer NOT NULL,
		duration integer NOT NULL,
		short_description text NOT NULL,
		created_at TIMESTAMPTZ NOT NULL,
		CONSTRAINT movie_pkey PRIMARY KEY (id)
	);

	CREATE TABLE IF NOT EXISTS review (
		id TEXT,
		first_name TEXT NOT NULL,
		last_name TEXT NOT NULL,
		score integer NOT NULL,
		text text NOT NULL,
		movie_id TEXT REFERENCES movie(id),
		created_at TIMESTAMPTZ NOT NULL,
		CONSTRAINT review_pkey PRIMARY KEY (id)
	)`
	s.DB.MustExec(schema)
}

func (s *Seed) clearTables() {
	s.DB.MustExec("DELETE FROM review")
	s.DB.MustExec("DELETE FROM movie")
}

func (s *Seed) populateMovies() {
	defaultMovies := []Movie{
		{
			ID:          "197451da-86fa-49d0-9702-817e8c885fb8",
			Title:       "Roma",
			ReleaseYear: 2018,
			Duration:    135,
			ShortDesc: "Cidade do México, 1970. A rotina de uma família de classe " +
				"média é controlada de maneira silenciosa por uma mulher " +
				"(Yalitza Aparicio), que trabalha como babá e empregada " +
				"doméstica. Durante um ano, diversos acontecimentos inesperados " +
				"começam a afetar a vida de todos os moradores da casa, dando " +
				"origem a uma série de mudanças, coletivas e pessoais.",
		},
		{
			ID:          "755671e8-7d2b-491a-ac8f-d147e56bf127",
			Title:       "Caixa de Pássaros",
			ReleaseYear: 2018,
			Duration:    125,
			ShortDesc: "Num cenário pós-apocalíptico onde o simples olhar pode te " +
				"levar à morte, uma mãe e seus dois filhos atravessam um rio de " +
				"olhos vendados em um barco, em busca de um lugar seguro.",
		},
		{
			ID:          "98e9542f-c215-499c-bbb6-24cbcf3ea7f1",
			Title:       "Monstros S.A.",
			ReleaseYear: 2001,
			Duration:    92,
			ShortDesc: "O astro do susto, Sulley, e seu falante assistente, Mike, " +
				"trabalham na Monstros S.A., a maior fábrica de processamento de " +
				"gritos da cidade de Monstrópolis. A principal fonte de energia " +
				"do mundo dos monstros provém da coleta dos gritos das crianças " +
				"humanas. Os monstros acreditam que as crianças são tóxicas, e " +
				"entram em Pânico quando uma menininha invade seu mundo. Sulley " +
				"e Mike fazem de tudo para levar a garota de volta para casa, " +
				"mas enfrentam desafios monstruosos e algumas situações hilárias " +
				"em suas atrapalhadas aventuras.",
		},
		{
			ID:          "10d58f60-5341-4f1d-b2a6-b8d7033f3c67",
			Title:       "Shrek",
			ReleaseYear: 2001,
			Duration:    90,
			ShortDesc: "Em um pântano distante vive Shrek (Mike Myers), um ogro " +
				"solitário que vê, sem mais nem menos, sua vida ser invadida por " +
				"uma série de personagens de contos de fada, como três ratos " +
				"cegos, um grande e malvado lobo e ainda três porcos que não têm " +
				"um lugar onde morar. Todos eles foram expulsos de seus lares " +
				"pelo maligno Lorde Farquaad (John Lithgow). Determinado a " +
				"recuperar a tranquilidade de antes, Shrek resolve encontrar " +
				"Farquaad e com ele faz um acordo: todos os personagens poderão " +
				"retornar aos seus lares se ele e seu amigo Burro (Eddie Murphy) " +
				"resgatarem uma bela princesa (Cameron Diaz), que é prisioneira " +
				"de um dragão. Porém, quando Shrek e o Burro enfim conseguem " +
				"resgatar a princesa logo eles descobrem que seus problemas " +
				"estão apenas começando.",
		},
		{
			ID:          "a45834f0-5b31-479e-89ae-5eb812158e8f",
			Title:       "Moana: Um Mar de Aventuras",
			ReleaseYear: 2016,
			Duration:    113,
			ShortDesc: "Moana Waialiki é uma corajosa jovem, filha do chefe de uma " +
				"tribo na Oceania, vinda de uma longa linhagem de navegadores, " +
				"que é seu maior hobbie e, também, trabalho. Querendo descobrir " +
				"mais sobre seu passado e ajudar sua família, ela resolve partir " +
				"em busca de seus ancestrais, habitantes de uma ilha mítica que " +
				"ninguém sabe onde é. Com a ajuda do lendário semideus Maui, " +
				"Moana começa sua jornada pelo mar aberto, onde vai enfrentar " +
				"criaturas marinhas e descobrir antigas histórias do submundo.",
		},
		{
			ID:          "724b4e9d-ab80-45e5-a167-6b64e7091d63",
			Title:       "Homem-Aranha no Aranhaverso",
			ReleaseYear: 2018,
			Duration:    120,
			ShortDesc: "Miles Morales é um jovem negro do Brooklyn que se tornou o " +
				"Homem-Aranha inspirado no legado de Peter Parker, já falecido. " +
				"Entretanto, ao visitar o túmulo de seu ídolo em uma noite " +
				"chuvosa, ele é surpreendido com a presença do próprio Peter, " +
				"vestindo o traje do herói aracnídeo sob um sobretudo. A surpresa" +
				" fica ainda maior quando Miles descobre que ele veio de uma " +
				"dimensão paralela, assim como outras versões do Homem-Aranha.",
		},
		{
			ID:          "cae41533-d480-4f12-9230-048c6d0dd005",
			Title:       "Onde os Fracos Não Têm Vez",
			ReleaseYear: 2007,
			Duration:    122,
			ShortDesc: "Texas, década de 80. Um traficante de drogas é encontrado " +
				"no deserto por um caçador pouco esperto, Llewelyn Moss (Josh " +
				"Brolin), que pega uma valise cheia de dinheiro mesmo sabendo que" +
				" em breve alguém irá procurá-lo devido a isso. Logo Anton " +
				"Chigurh (Javier Bardem), um assassino psicótico sem senso de " +
				"humor e piedade, é enviado em seu encalço. Porém para alcançar " +
				"Moss ele precisará passar pelo xerife local, Ed Tom Bell " +
				"(Tommy Lee Jones).",
		},
		{
			ID:          "08b771f1-89ea-4cfb-9b45-35daa794ceec",
			Title:       "Gran Torino",
			ReleaseYear: 2008,
			Duration:    117,
			ShortDesc: "Walt Kowalski (Clint Eastwood) é um inflexível veterano " +
				"da Guerra da Coréia, que está agora aposentado. Para passar o " +
				"tempo ele faz consertos em casa, bebe cerveja e vai mensalmente " +
				"ao barbeiro (John Carroll Lynch). Sua vida é alterada quando " +
				"passa a ter como vizinhos imigrantes hmong, vindos do Laos, os " +
				"quais Walt despreza. Ressentido e desconfiando de todos, Walt " +
				"apenas deseja passar o tempo que lhe resta de vida. Até que " +
				"Thao (Bee Vang), seu tímido vizinho adolescente, é obrigado por " +
				"uma gangue a roubar o carro de Walt, um Gran Torino retirado da " +
				"linha de montagem pelo próprio. Walt consegue impedir o roubo, o" +
				" que faz com que se torne uma espécie de herói local, " +
				"especialmente para Sue (Ahney Her), irmã de Thao, que insiste " +
				"que deve trabalhar para Walt como forma de recompensá-lo.",
		},
		{
			ID:          "9182a3d6-15e1-4777-9e2e-380bd5b7167d",
			Title:       "O Poderoso Chefão",
			ReleaseYear: 1972,
			Duration:    177,
			ShortDesc: "Em 1945, Don Corleone (Marlon Brando) é o chefe de uma " +
				"mafiosa família italiana de Nova York. Ele costuma apadrinhar " +
				"várias pessoas, realizando importantes favores para elas, em " +
				"troca de favores futuros. Com a chegada das drogas, as famílias " +
				"começam uma disputa pelo promissor mercado. Quando Corleone se " +
				"recusa a facilitar a entrada dos narcóticos na cidade, não " +
				"oferecendo ajuda política e policial, sua família começa a " +
				"sofrer atentados para que mudem de posição. É nessa complicada " +
				"época que Michael (Al Pacino), um herói de guerra nunca " +
				"envolvido nos negócios da família, vê a necessidade de proteger " +
				"o seu pai e tudo o que ele construiu ao longo dos anos.",
		},
		{
			ID:          "ba313e4f-43ad-4704-bfcf-0b3561338ed0",
			Title:       "Titanic",
			ReleaseYear: 1997,
			Duration:    194,
			ShortDesc: "Uma expedição aos destroços do Titanic leva uma " +
				"sobrevivente do naufrágio a relembrar uma grande história de " +
				"amor que viveu no navio. Em 1912, na única viagem do que então " +
				"era o maior navio já construído, Rose (Winslet) é uma jovem da " +
				"alta sociedade prestes a se casar com seu rico noivo. Mas a " +
				"bordo do Titanic ela conhece Jack Dawson (DiCaprio), um jovem " +
				"simples e aventureiro, e se apaixona pelo rapaz. As diferenças " +
				"sociais fazem com que muitos se oponham ao relacionamento que " +
				"surge. Em meio ao intenso romance e à rebeldia dos dois, " +
				"acontece o trágico acidente, que eles enfrentam juntos.",
		},
	}

	for _, movie := range defaultMovies {
		movie.CreatedAt = time.Now()
		tx := s.DB.MustBegin()
		tx.NamedExec("INSERT INTO movie (id, title, release_year, duration, short_description, created_at) VALUES (:id, :title, :release_year, :duration, :short_description, :created_at)", &movie)
		err := tx.Commit()
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (s *Seed) populateReviews() {
	defaultReviews := []Review{
		{
			ID:        "07cb68b0-1ce9-4a9c-9be9-35bee798799d",
			MovieID:   "ba313e4f-43ad-4704-bfcf-0b3561338ed0",
			FirstName: "Sheldon",
			LastName:  "Cooper",
			Score:     5,
			Text:      "Bazinga",
		},
		{
			ID:        "d76a9ba8-9498-4923-9405-43c16cd1e80b",
			MovieID:   "9182a3d6-15e1-4777-9e2e-380bd5b7167d",
			FirstName: "Leonard",
			LastName:  "Hofstadter",
			Score:     1,
			Text:      "Achei o filme muito violento.",
		},
		{
			ID:        "5b887d05-27d5-4419-956e-61f14bcc53df",
			MovieID:   "9182a3d6-15e1-4777-9e2e-380bd5b7167d",
			FirstName: "Howard",
			LastName:  "Wolowitz",
			Score:     3,
			Text:      "Infelizmente não tinha comédia no filme.",
		},
		{
			ID:        "333e1c08-cb34-4b50-b925-f5433919e393",
			MovieID:   "ba313e4f-43ad-4704-bfcf-0b3561338ed0",
			FirstName: "Raj",
			LastName:  "Koothrappali",
			Score:     4,
			Text:      "Chorei horrores com a morte do Jack.",
		},
		{
			ID:        "f2a17805-2b44-43a4-873e-5ca4784342b9",
			MovieID:   "ba313e4f-43ad-4704-bfcf-0b3561338ed0",
			FirstName: "Stuart",
			LastName:  "Bloom",
			Score:     2,
			Text:      "Não era um filme de super-heróis.",
		},
		{
			ID:        "62291783-9b26-47df-a840-767e44688971",
			MovieID:   "9182a3d6-15e1-4777-9e2e-380bd5b7167d",
			FirstName: "Penny",
			LastName:  "Hofstadter",
			Score:     5,
			Text:      "Eu sou de Nebraska .",
		},
	}

	for _, review := range defaultReviews {
		review.CreatedAt = time.Now()
		tx := s.DB.MustBegin()
		tx.NamedExec("INSERT INTO review (id, movie_id, first_name, last_name, score, text, created_at) VALUES (:id, :movie_id, :first_name, :last_name, :score, :text, :created_at)", &review)
		err := tx.Commit()
		if err != nil {
			log.Fatal(err)
		}
	}
}
