package service

import (
	"github.com/TremblingV5/DouTok/applications/comment/dal/model"
	"github.com/TremblingV5/DouTok/applications/comment/misc"
	"github.com/TremblingV5/DouTok/pkg/errno"
	"github.com/TremblingV5/DouTok/pkg/hbaseHandle"
	tools "github.com/TremblingV5/DouTok/pkg/misc"
)

func ListComment(video_id int64) ([]*model.CommentInHB, errno.ErrNo, error) {
	res, err := HBClient.Scan(
		"comment",
		hbaseHandle.GetFilterByRowKeyPrefix(misc.GetCommentQueryPrefix(video_id))...,
	)

	if err != nil {
		return nil, misc.QueryCommentListInHBErr, err
	}

	comment_list := []*model.CommentInHB{}

	for _, v := range res {
		temp := model.CommentInHB{}
		err := tools.Map2Struct(v, &temp)
		if err != nil {
			continue
		}
		comment_list = append(comment_list, &temp)
	}

	return comment_list, misc.Success, nil
}