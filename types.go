package main


type Project struct {
  //struct for todoist project
  ID string `json:"id"`
  Name string `json:"name"`
  Order int `json:"order"`
  Color string `json:"color"`
  Shared bool `json:"is_shared"`
  Favorite bool `json:"is_favorite"`
  ParentID int `json:"parent_id"`
  ViewStyle string `json:"view_style"`
  URL string `json:"url"`
  Tasks Tasks
}

type Projects []Project

type Task struct {
  //struct for todoist item
  ID string `json:"id"`
  ProjectID string `json:"project_id"`
  SectionId string `json:"section_id"`
  Content string `json:"content"`
  Completed bool `json:"is_completed"`
  Order int `json:"order"`
  Priority int `json:"priority"`
  URL string `json:"url"`
  Due Due `json:"due"`
  LabelIDs []int `json:"label_ids"`
  CommentCount int `json:"comment_count"`
}

type Due struct {
  String string `json:"string"`
  Date string `json:"date"`
  Recurring bool `json:"recurring"`
  DateTime string `json:"datetime"`
  TimeZone string `json:"timezone"`
}

type Tasks []Task

type Label struct {
  //struct for todoist label
  ID string `json:"id"`
  Name string `json:"name"`
  Color int `json:"color"`
  Order int `json:"order"`
  Favorite bool `json:"is_favorite"`
}

type Labels []Label

type model struct {
    projects  Projects           // items on the to-do list
    cursor   int                // which to-do list item our cursor is pointing at
}
