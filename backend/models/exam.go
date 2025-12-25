package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

// TODO: add latex and image support in problem and choice!
type Choice struct {
	Id   int    `json:"id"`
	Text string `json:"text"`
}

type Problem struct {
	Question string   `json:"question"`
	AnswerId int      `json:"answerId"`
	Choice   []Choice `json:"choice"`
}

type Exam struct {
	Id bson.ObjectID `bson:"_id,omitempty" json:"id"`

	Title       string    `bson:"title" json:"title"`
	ProblemList []Problem `bson:"problemList" json:"problemList"`

	OpenAt  time.Time `bson:"openAt" json:"openAt"`
	CloseAt time.Time `bson:"closeAt" json:"closeAt"`

	ExamDuration     time.Duration `bson:"examDuration" json:"examDuration"`
	AllowExceedClose bool          `bson:"allowExceedClose" json:"allowExceedClose"`

	ShuffleQuestion bool `bson:"shuffleQuestion" json:"shuffleQuestion"`
}

type ExamCreateDTO struct {
	Title       string    `bson:"title" json:"title" binding:"required"`
	ProblemList []Problem `bson:"problemList" json:"problemList"`

	OpenAt  time.Time `bson:"openAt" json:"openAt"`
	CloseAt time.Time `bson:"closeAt" json:"closeAt"`

	ExamDuration     time.Duration `bson:"examDuration" json:"examDuration"`
	AllowExceedClose bool          `bson:"allowExceedClose" json:"allowExceedClose"`

	ShuffleQuestion bool `bson:"shuffleQuestion" json:"shuffleQuestion"`
}
