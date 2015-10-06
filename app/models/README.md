# Data Models
------

Status: Stub

## Caching layer (per publisher)

The idea is to have a 'caching layer' that gets the data we need from the newsroom's systems.
We are going to have several strategies to do the translation/mapping from their systems to our own.

### NY Times Strategy

type: user, item, boolAction, associationAction, externalReference

Schema
```{
  database: nyt,

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
          local: id,

          relation: 'self',
          model: Comment
        },
        {},
        {
          name: UserID,
          type: int,
          local: id,

          relation: 'belongsTo',
          model: User
        },
        {
          name: AssetId,
          type: []byte,
          local: AssetId,

          relation: 'belongsTo',
          model: Asset
        }.
        {
          name: Comment,
          type: int,
          local: ParentId,

          relation: 'hasMany',
          model: Comment
        }
      ],    
    },

  "user": {

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
}```

#### NYT Comment's fields

1. commentID: comment's identifier. Numeric. - YES
2. assetID: Numeric. Article's identifier. - YES
3. statusID: Numeric. Options: [2, 3]  - moderated(1), rejected(2), flagged(3) - YES (ModerationStatus)
4. commentTitle: String. Maybe NULL  - deprecated
5. commentBody: String. Maybe NULL - YES (Content)
6. userID: user that commented. Numeric.  (registration id) - YES
7. createDate: TimeDate   - YES (sourceCreatedDate)
8. updateDate: TimeDate   - YES
9. approveDate: TimeDate  - YES
10. commentExcerpt: Summary? String. Maybe NULL  - deprectaed
11. editorsSelection: Boolean  - nyt picks - YES (this is a BooleanAction)
12. recommendationCount: Numeric  - how many people recommended the comment  - YES
13. replyCount: Numeric  - how many people reply to that comment - NO
14. isReply: Boolean  - if it is a reply or not - NO
15. commentSequence: Numeric  - NO
16. userDisplayName: The user's name. String.  Deprecated. - YES (into User model)
17. userURL: Deprecated.
18. userTitle:  Deprecated.
19. userLocation: City. String. Deprecated. - YES (into User model - json list.. )
20. showCommentExcerpt: Boolean  - NO
21. hideRegisteredUserName: Boolean  - Deprecated?? - NO (check if any is not zero)
22. commentType: Deprecated. - NO
23. parentID: The id of the parent's comment. - YES


#### Data Types

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
