# TL;DR

この Actions は Issue / PR コメントに対して､編集可能なコメントを作成することができます｡
もうこれ以上､毎回過去のコメントを `outdate` する必要はありません｡

This Actions allows you to create editable comments on Issue / PR comments.
No more `outdating` past comments every time.
<img width="859" alt="スクリーンショット 2023-06-01 14 59 46 2" src="https://github.com/nakamuloud/actions-rewritable-comment/assets/25874783/c5df1b6d-db9a-4c22-85be-301f567958c3">



# Getting Started

## Works with on.pull_request (minimal option)

```yaml
- name: rewritable-issue-comment
  uses: nakamuloud/actions-rewritable-comment@v1.0.0
  env:
    GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
    VALUE: 'Hello World'
```

Will posts / edits below comment to the PR

```text
Hello World
<!-- actions-rewritable-comment:default -->
```

## Works with other triggers

```yaml
- name: rewritable-issue-comment
  uses: nakamuloud/actions-rewritable-comment@v1.0.0
  env:
    GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
    VALUE: 'Hello World'
    ISSUE_NUMBER: 1
```

Will posts / edits below comment to the Issue / PR (id=1)

```text
Hello World
<!-- actions-rewritable-comment:default -->
```

## Works with other triggers to another Repository

```yaml
- name: rewritable-issue-comment
  uses: nakamuloud/actions-rewritable-comment@v1.0.0
  env:
    GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
    VALUE: 'Hello World'
    ISSUE_NUMBER: 1
    OWNER: example
    REPOSITORY: dotfiles
```

Will posts / edits below comment to the Issue / PR (id=1) at example/dotfiles repository

```text
Hello World
<!-- actions-rewritable-comment:default -->
```

## Works with full options

```yaml
- name: rewritable-issue-comment
  uses: nakamuloud/actions-rewritable-comment@v1.0.0
  env:
    GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
    VALUE: 'Hello World'
    ISSUE_NUMBER: 1
    OWNER: example
    REPOSITORY: dotfiles
    KEY: key
    COMMENT_ID: comment_id
```

Will posts / edits below comment to the Issue / PR (id=1) at example/dotfiles repository with the annotation of `key:comment_id`

```text
Hello World
<!-- key:comment -->
```
