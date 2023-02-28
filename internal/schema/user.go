package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").DefaultFunc(func() string {
			return uuid.New().String()
		}).Comment("유저 아이디"),
		field.String("name").NotEmpty().Optional().Comment("이름"),
		field.Time("created_at").Default(time.Now).Immutable().Comment("생성된 시간"),
		field.Time("updated_at").Default(time.Now).Comment("수정된 시간"),
		field.Time("deleted_at").Optional().Nillable().Default(nil).Comment("삭제된 시간"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("todos", Todo.Type),
	}
}
