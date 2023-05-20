export interface Author {
    name: string;
  }
  
export interface Article {
id: number;
author: Author;
title: string;
description: string;
body: string;
tagList: string[];
}

export interface Comment {
id: number;
articleId: number;
comment: string;
}

export interface State {
articles: Article[];
comments: Comment[];
}