// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/laclipasa/la-clipasa/ent/post"
	"github.com/laclipasa/la-clipasa/ent/predicate"
	"github.com/laclipasa/la-clipasa/ent/user"
)

// PostUpdate is the builder for updating Post entities.
type PostUpdate struct {
	config
	hooks    []Hook
	mutation *PostMutation
}

// Where appends a list predicates to the PostUpdate builder.
func (pu *PostUpdate) Where(ps ...predicate.Post) *PostUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetPinned sets the "pinned" field.
func (pu *PostUpdate) SetPinned(b bool) *PostUpdate {
	pu.mutation.SetPinned(b)
	return pu
}

// SetNillablePinned sets the "pinned" field if the given value is not nil.
func (pu *PostUpdate) SetNillablePinned(b *bool) *PostUpdate {
	if b != nil {
		pu.SetPinned(*b)
	}
	return pu
}

// SetUserID sets the "user_id" field.
func (pu *PostUpdate) SetUserID(u uuid.UUID) *PostUpdate {
	pu.mutation.SetUserID(u)
	return pu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (pu *PostUpdate) SetNillableUserID(u *uuid.UUID) *PostUpdate {
	if u != nil {
		pu.SetUserID(*u)
	}
	return pu
}

// SetTitle sets the "title" field.
func (pu *PostUpdate) SetTitle(s string) *PostUpdate {
	pu.mutation.SetTitle(s)
	return pu
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (pu *PostUpdate) SetNillableTitle(s *string) *PostUpdate {
	if s != nil {
		pu.SetTitle(*s)
	}
	return pu
}

// SetContent sets the "content" field.
func (pu *PostUpdate) SetContent(s string) *PostUpdate {
	pu.mutation.SetContent(s)
	return pu
}

// SetNillableContent sets the "content" field if the given value is not nil.
func (pu *PostUpdate) SetNillableContent(s *string) *PostUpdate {
	if s != nil {
		pu.SetContent(*s)
	}
	return pu
}

// ClearContent clears the value of the "content" field.
func (pu *PostUpdate) ClearContent() *PostUpdate {
	pu.mutation.ClearContent()
	return pu
}

// SetLink sets the "link" field.
func (pu *PostUpdate) SetLink(s string) *PostUpdate {
	pu.mutation.SetLink(s)
	return pu
}

// SetNillableLink sets the "link" field if the given value is not nil.
func (pu *PostUpdate) SetNillableLink(s *string) *PostUpdate {
	if s != nil {
		pu.SetLink(*s)
	}
	return pu
}

// SetModerationComment sets the "moderation_comment" field.
func (pu *PostUpdate) SetModerationComment(s string) *PostUpdate {
	pu.mutation.SetModerationComment(s)
	return pu
}

// SetNillableModerationComment sets the "moderation_comment" field if the given value is not nil.
func (pu *PostUpdate) SetNillableModerationComment(s *string) *PostUpdate {
	if s != nil {
		pu.SetModerationComment(*s)
	}
	return pu
}

// ClearModerationComment clears the value of the "moderation_comment" field.
func (pu *PostUpdate) ClearModerationComment() *PostUpdate {
	pu.mutation.ClearModerationComment()
	return pu
}

// SetIsModerated sets the "is_moderated" field.
func (pu *PostUpdate) SetIsModerated(b bool) *PostUpdate {
	pu.mutation.SetIsModerated(b)
	return pu
}

// SetNillableIsModerated sets the "is_moderated" field if the given value is not nil.
func (pu *PostUpdate) SetNillableIsModerated(b *bool) *PostUpdate {
	if b != nil {
		pu.SetIsModerated(*b)
	}
	return pu
}

// SetUpdatedAt sets the "updated_at" field.
func (pu *PostUpdate) SetUpdatedAt(t time.Time) *PostUpdate {
	pu.mutation.SetUpdatedAt(t)
	return pu
}

// SetCategories sets the "categories" field.
func (pu *PostUpdate) SetCategories(po post.Categories) *PostUpdate {
	pu.mutation.SetCategories(po)
	return pu
}

// SetNillableCategories sets the "categories" field if the given value is not nil.
func (pu *PostUpdate) SetNillableCategories(po *post.Categories) *PostUpdate {
	if po != nil {
		pu.SetCategories(*po)
	}
	return pu
}

// AddSavedByIDs adds the "saved_by" edge to the User entity by IDs.
func (pu *PostUpdate) AddSavedByIDs(ids ...int) *PostUpdate {
	pu.mutation.AddSavedByIDs(ids...)
	return pu
}

// AddSavedBy adds the "saved_by" edges to the User entity.
func (pu *PostUpdate) AddSavedBy(u ...*User) *PostUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return pu.AddSavedByIDs(ids...)
}

// AddLikedByIDs adds the "liked_by" edge to the User entity by IDs.
func (pu *PostUpdate) AddLikedByIDs(ids ...int) *PostUpdate {
	pu.mutation.AddLikedByIDs(ids...)
	return pu
}

// AddLikedBy adds the "liked_by" edges to the User entity.
func (pu *PostUpdate) AddLikedBy(u ...*User) *PostUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return pu.AddLikedByIDs(ids...)
}

// Mutation returns the PostMutation object of the builder.
func (pu *PostUpdate) Mutation() *PostMutation {
	return pu.mutation
}

// ClearSavedBy clears all "saved_by" edges to the User entity.
func (pu *PostUpdate) ClearSavedBy() *PostUpdate {
	pu.mutation.ClearSavedBy()
	return pu
}

// RemoveSavedByIDs removes the "saved_by" edge to User entities by IDs.
func (pu *PostUpdate) RemoveSavedByIDs(ids ...int) *PostUpdate {
	pu.mutation.RemoveSavedByIDs(ids...)
	return pu
}

// RemoveSavedBy removes "saved_by" edges to User entities.
func (pu *PostUpdate) RemoveSavedBy(u ...*User) *PostUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return pu.RemoveSavedByIDs(ids...)
}

// ClearLikedBy clears all "liked_by" edges to the User entity.
func (pu *PostUpdate) ClearLikedBy() *PostUpdate {
	pu.mutation.ClearLikedBy()
	return pu
}

// RemoveLikedByIDs removes the "liked_by" edge to User entities by IDs.
func (pu *PostUpdate) RemoveLikedByIDs(ids ...int) *PostUpdate {
	pu.mutation.RemoveLikedByIDs(ids...)
	return pu
}

// RemoveLikedBy removes "liked_by" edges to User entities.
func (pu *PostUpdate) RemoveLikedBy(u ...*User) *PostUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return pu.RemoveLikedByIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PostUpdate) Save(ctx context.Context) (int, error) {
	pu.defaults()
	return withHooks(ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PostUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PostUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PostUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pu *PostUpdate) defaults() {
	if _, ok := pu.mutation.UpdatedAt(); !ok {
		v := post.UpdateDefaultUpdatedAt()
		pu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pu *PostUpdate) check() error {
	if v, ok := pu.mutation.Link(); ok {
		if err := post.LinkValidator(v); err != nil {
			return &ValidationError{Name: "link", err: fmt.Errorf(`ent: validator failed for field "Post.link": %w`, err)}
		}
	}
	if v, ok := pu.mutation.Categories(); ok {
		if err := post.CategoriesValidator(v); err != nil {
			return &ValidationError{Name: "categories", err: fmt.Errorf(`ent: validator failed for field "Post.categories": %w`, err)}
		}
	}
	return nil
}

func (pu *PostUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := pu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(post.Table, post.Columns, sqlgraph.NewFieldSpec(post.FieldID, field.TypeInt))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Pinned(); ok {
		_spec.SetField(post.FieldPinned, field.TypeBool, value)
	}
	if value, ok := pu.mutation.UserID(); ok {
		_spec.SetField(post.FieldUserID, field.TypeUUID, value)
	}
	if value, ok := pu.mutation.Title(); ok {
		_spec.SetField(post.FieldTitle, field.TypeString, value)
	}
	if value, ok := pu.mutation.Content(); ok {
		_spec.SetField(post.FieldContent, field.TypeString, value)
	}
	if pu.mutation.ContentCleared() {
		_spec.ClearField(post.FieldContent, field.TypeString)
	}
	if value, ok := pu.mutation.Link(); ok {
		_spec.SetField(post.FieldLink, field.TypeString, value)
	}
	if value, ok := pu.mutation.ModerationComment(); ok {
		_spec.SetField(post.FieldModerationComment, field.TypeString, value)
	}
	if pu.mutation.ModerationCommentCleared() {
		_spec.ClearField(post.FieldModerationComment, field.TypeString)
	}
	if value, ok := pu.mutation.IsModerated(); ok {
		_spec.SetField(post.FieldIsModerated, field.TypeBool, value)
	}
	if value, ok := pu.mutation.UpdatedAt(); ok {
		_spec.SetField(post.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := pu.mutation.Categories(); ok {
		_spec.SetField(post.FieldCategories, field.TypeEnum, value)
	}
	if pu.mutation.SavedByCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   post.SavedByTable,
			Columns: post.SavedByPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedSavedByIDs(); len(nodes) > 0 && !pu.mutation.SavedByCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   post.SavedByTable,
			Columns: post.SavedByPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.SavedByIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   post.SavedByTable,
			Columns: post.SavedByPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.LikedByCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   post.LikedByTable,
			Columns: post.LikedByPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedLikedByIDs(); len(nodes) > 0 && !pu.mutation.LikedByCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   post.LikedByTable,
			Columns: post.LikedByPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.LikedByIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   post.LikedByTable,
			Columns: post.LikedByPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{post.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// PostUpdateOne is the builder for updating a single Post entity.
type PostUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PostMutation
}

// SetPinned sets the "pinned" field.
func (puo *PostUpdateOne) SetPinned(b bool) *PostUpdateOne {
	puo.mutation.SetPinned(b)
	return puo
}

// SetNillablePinned sets the "pinned" field if the given value is not nil.
func (puo *PostUpdateOne) SetNillablePinned(b *bool) *PostUpdateOne {
	if b != nil {
		puo.SetPinned(*b)
	}
	return puo
}

// SetUserID sets the "user_id" field.
func (puo *PostUpdateOne) SetUserID(u uuid.UUID) *PostUpdateOne {
	puo.mutation.SetUserID(u)
	return puo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (puo *PostUpdateOne) SetNillableUserID(u *uuid.UUID) *PostUpdateOne {
	if u != nil {
		puo.SetUserID(*u)
	}
	return puo
}

// SetTitle sets the "title" field.
func (puo *PostUpdateOne) SetTitle(s string) *PostUpdateOne {
	puo.mutation.SetTitle(s)
	return puo
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (puo *PostUpdateOne) SetNillableTitle(s *string) *PostUpdateOne {
	if s != nil {
		puo.SetTitle(*s)
	}
	return puo
}

// SetContent sets the "content" field.
func (puo *PostUpdateOne) SetContent(s string) *PostUpdateOne {
	puo.mutation.SetContent(s)
	return puo
}

// SetNillableContent sets the "content" field if the given value is not nil.
func (puo *PostUpdateOne) SetNillableContent(s *string) *PostUpdateOne {
	if s != nil {
		puo.SetContent(*s)
	}
	return puo
}

// ClearContent clears the value of the "content" field.
func (puo *PostUpdateOne) ClearContent() *PostUpdateOne {
	puo.mutation.ClearContent()
	return puo
}

// SetLink sets the "link" field.
func (puo *PostUpdateOne) SetLink(s string) *PostUpdateOne {
	puo.mutation.SetLink(s)
	return puo
}

// SetNillableLink sets the "link" field if the given value is not nil.
func (puo *PostUpdateOne) SetNillableLink(s *string) *PostUpdateOne {
	if s != nil {
		puo.SetLink(*s)
	}
	return puo
}

// SetModerationComment sets the "moderation_comment" field.
func (puo *PostUpdateOne) SetModerationComment(s string) *PostUpdateOne {
	puo.mutation.SetModerationComment(s)
	return puo
}

// SetNillableModerationComment sets the "moderation_comment" field if the given value is not nil.
func (puo *PostUpdateOne) SetNillableModerationComment(s *string) *PostUpdateOne {
	if s != nil {
		puo.SetModerationComment(*s)
	}
	return puo
}

// ClearModerationComment clears the value of the "moderation_comment" field.
func (puo *PostUpdateOne) ClearModerationComment() *PostUpdateOne {
	puo.mutation.ClearModerationComment()
	return puo
}

// SetIsModerated sets the "is_moderated" field.
func (puo *PostUpdateOne) SetIsModerated(b bool) *PostUpdateOne {
	puo.mutation.SetIsModerated(b)
	return puo
}

// SetNillableIsModerated sets the "is_moderated" field if the given value is not nil.
func (puo *PostUpdateOne) SetNillableIsModerated(b *bool) *PostUpdateOne {
	if b != nil {
		puo.SetIsModerated(*b)
	}
	return puo
}

// SetUpdatedAt sets the "updated_at" field.
func (puo *PostUpdateOne) SetUpdatedAt(t time.Time) *PostUpdateOne {
	puo.mutation.SetUpdatedAt(t)
	return puo
}

// SetCategories sets the "categories" field.
func (puo *PostUpdateOne) SetCategories(po post.Categories) *PostUpdateOne {
	puo.mutation.SetCategories(po)
	return puo
}

// SetNillableCategories sets the "categories" field if the given value is not nil.
func (puo *PostUpdateOne) SetNillableCategories(po *post.Categories) *PostUpdateOne {
	if po != nil {
		puo.SetCategories(*po)
	}
	return puo
}

// AddSavedByIDs adds the "saved_by" edge to the User entity by IDs.
func (puo *PostUpdateOne) AddSavedByIDs(ids ...int) *PostUpdateOne {
	puo.mutation.AddSavedByIDs(ids...)
	return puo
}

// AddSavedBy adds the "saved_by" edges to the User entity.
func (puo *PostUpdateOne) AddSavedBy(u ...*User) *PostUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return puo.AddSavedByIDs(ids...)
}

// AddLikedByIDs adds the "liked_by" edge to the User entity by IDs.
func (puo *PostUpdateOne) AddLikedByIDs(ids ...int) *PostUpdateOne {
	puo.mutation.AddLikedByIDs(ids...)
	return puo
}

// AddLikedBy adds the "liked_by" edges to the User entity.
func (puo *PostUpdateOne) AddLikedBy(u ...*User) *PostUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return puo.AddLikedByIDs(ids...)
}

// Mutation returns the PostMutation object of the builder.
func (puo *PostUpdateOne) Mutation() *PostMutation {
	return puo.mutation
}

// ClearSavedBy clears all "saved_by" edges to the User entity.
func (puo *PostUpdateOne) ClearSavedBy() *PostUpdateOne {
	puo.mutation.ClearSavedBy()
	return puo
}

// RemoveSavedByIDs removes the "saved_by" edge to User entities by IDs.
func (puo *PostUpdateOne) RemoveSavedByIDs(ids ...int) *PostUpdateOne {
	puo.mutation.RemoveSavedByIDs(ids...)
	return puo
}

// RemoveSavedBy removes "saved_by" edges to User entities.
func (puo *PostUpdateOne) RemoveSavedBy(u ...*User) *PostUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return puo.RemoveSavedByIDs(ids...)
}

// ClearLikedBy clears all "liked_by" edges to the User entity.
func (puo *PostUpdateOne) ClearLikedBy() *PostUpdateOne {
	puo.mutation.ClearLikedBy()
	return puo
}

// RemoveLikedByIDs removes the "liked_by" edge to User entities by IDs.
func (puo *PostUpdateOne) RemoveLikedByIDs(ids ...int) *PostUpdateOne {
	puo.mutation.RemoveLikedByIDs(ids...)
	return puo
}

// RemoveLikedBy removes "liked_by" edges to User entities.
func (puo *PostUpdateOne) RemoveLikedBy(u ...*User) *PostUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return puo.RemoveLikedByIDs(ids...)
}

// Where appends a list predicates to the PostUpdate builder.
func (puo *PostUpdateOne) Where(ps ...predicate.Post) *PostUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PostUpdateOne) Select(field string, fields ...string) *PostUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Post entity.
func (puo *PostUpdateOne) Save(ctx context.Context) (*Post, error) {
	puo.defaults()
	return withHooks(ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PostUpdateOne) SaveX(ctx context.Context) *Post {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PostUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PostUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (puo *PostUpdateOne) defaults() {
	if _, ok := puo.mutation.UpdatedAt(); !ok {
		v := post.UpdateDefaultUpdatedAt()
		puo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (puo *PostUpdateOne) check() error {
	if v, ok := puo.mutation.Link(); ok {
		if err := post.LinkValidator(v); err != nil {
			return &ValidationError{Name: "link", err: fmt.Errorf(`ent: validator failed for field "Post.link": %w`, err)}
		}
	}
	if v, ok := puo.mutation.Categories(); ok {
		if err := post.CategoriesValidator(v); err != nil {
			return &ValidationError{Name: "categories", err: fmt.Errorf(`ent: validator failed for field "Post.categories": %w`, err)}
		}
	}
	return nil
}

func (puo *PostUpdateOne) sqlSave(ctx context.Context) (_node *Post, err error) {
	if err := puo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(post.Table, post.Columns, sqlgraph.NewFieldSpec(post.FieldID, field.TypeInt))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Post.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, post.FieldID)
		for _, f := range fields {
			if !post.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != post.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.Pinned(); ok {
		_spec.SetField(post.FieldPinned, field.TypeBool, value)
	}
	if value, ok := puo.mutation.UserID(); ok {
		_spec.SetField(post.FieldUserID, field.TypeUUID, value)
	}
	if value, ok := puo.mutation.Title(); ok {
		_spec.SetField(post.FieldTitle, field.TypeString, value)
	}
	if value, ok := puo.mutation.Content(); ok {
		_spec.SetField(post.FieldContent, field.TypeString, value)
	}
	if puo.mutation.ContentCleared() {
		_spec.ClearField(post.FieldContent, field.TypeString)
	}
	if value, ok := puo.mutation.Link(); ok {
		_spec.SetField(post.FieldLink, field.TypeString, value)
	}
	if value, ok := puo.mutation.ModerationComment(); ok {
		_spec.SetField(post.FieldModerationComment, field.TypeString, value)
	}
	if puo.mutation.ModerationCommentCleared() {
		_spec.ClearField(post.FieldModerationComment, field.TypeString)
	}
	if value, ok := puo.mutation.IsModerated(); ok {
		_spec.SetField(post.FieldIsModerated, field.TypeBool, value)
	}
	if value, ok := puo.mutation.UpdatedAt(); ok {
		_spec.SetField(post.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := puo.mutation.Categories(); ok {
		_spec.SetField(post.FieldCategories, field.TypeEnum, value)
	}
	if puo.mutation.SavedByCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   post.SavedByTable,
			Columns: post.SavedByPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedSavedByIDs(); len(nodes) > 0 && !puo.mutation.SavedByCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   post.SavedByTable,
			Columns: post.SavedByPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.SavedByIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   post.SavedByTable,
			Columns: post.SavedByPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.LikedByCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   post.LikedByTable,
			Columns: post.LikedByPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedLikedByIDs(); len(nodes) > 0 && !puo.mutation.LikedByCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   post.LikedByTable,
			Columns: post.LikedByPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.LikedByIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   post.LikedByTable,
			Columns: post.LikedByPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Post{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{post.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}