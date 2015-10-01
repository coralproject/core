== Data Models

Status: Stub

==== Caching layer (per publisher)

The idea is to have a 'caching layer' that gets the data we need from the newsroom's systems.
We are going to have several strategies to do the translation/mapping from their systems to our own.

======= JSON File with the NY Times Strategy

type: user, item, boolAction, associationAction, externalReference

Schema
{

  "article": {
    type: externalReference,
    table: {
      local: articles
    },
    fields: [
      {
        local: id,
        type: int
      },
      {
        name: url,
        type: url
      },
      {
        name: publisher
        type: string
      }   
    ]
  },

  "comment": {
      type: item,
      table: {
        foreign: comments,
        local: comments
      },
      fields: [
        {
          name: CommentID,
          type: int,
          primaryKey: true,
          local: id
        },
        {
          name: UserID,
          type: int,
          local: userId
        },
        {
          name: Content,
          type: []byte,
          local: content
        }
      ],    
    },

  "comment_recommend": {
      type: boolAction,
      table: {
        foreign: recommendations
        local: recommendations
      },
      fields: [
        {
          name: UserID,
          type: int,
          local: userid
        },
        {
          name: CommentID,
          type: int,
          local: commentid
        },
      ]
    }
}


"relation": [
  {
    name: 'reply',
    type: hasMany,
    table: comment,
    fKey: parent
  }
]

===== Type of Tables

Users: people
  Authors
  Moderators
  Contributors

Item: a piece of content/information that we pull into our databases.
  Comment

BoolAction: boolean (yes/no) actions between two items or Users
  User recommends Comment
  User likes Article
  User shares Article
  Author writes Article
  User mutes User
  User mutes Article
  User follows User

AssociationAction or ¿Relations?
  Comment hasMany Comments
  Article hasMany Comments
  Section hasMany Articles


ExternalReference: a reference to a piece of content/information
  Article
  Section
