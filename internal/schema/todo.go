package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Todo holds the schema definition for the Todo entity.
type Todo struct {
	ent.Schema
}

// Fields of the Todo.
func (Todo) Fields() []ent.Field {
	return []ent.Field{
		field.Text("text").NotEmpty().Optional().Comment("할일"),
		field.Bool("done").Default(false).Comment("처리 상태 "),
		field.Time("created_at").Default(time.Now).Immutable().Comment("생성된 시간"),
		field.Time("updated_at").Default(time.Now).Comment("수정된 시간"),
		field.Time("deleted_at").Optional().Nillable().Default(nil).Comment("삭제된 시간"),
		field.String("user_id").Unique(),
	}
}

// Edges of the Todo.
func (Todo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("todos").Field("user_id").Unique().Required(),
	}
}
