package entdb

import (
	"context"

	"github.com/YadaYuki/omochi/app/domain/entities"
	"github.com/YadaYuki/omochi/app/domain/repository"
	"github.com/YadaYuki/omochi/app/ent"
	"github.com/YadaYuki/omochi/app/errors"
	"github.com/YadaYuki/omochi/app/errors/code"
)

type InvertedIndexCompressedEntRepository struct {
	db *ent.Client
}

func NewInvertedIndexCompressedEntRepository(db *ent.Client) repository.InvertedIndexCompressedRepository {
	return &InvertedIndexCompressedEntRepository{db: db}
}

func (r *InvertedIndexCompressedEntRepository) BulkCreateInvertIndexesCompressed(ctx context.Context, invertIndexes *[]entities.InvertedIndexCompressed) (*[]entities.InvertedIndexCompressedDetail, *errors.Error) {
	bulk := make([]*ent.InvertIndexCompressedCreate, len(*invertIndexes))
	for i, invertIndex := range *invertIndexes {
		// TODO: Set Term Id
		bulk[i] = r.db.InvertIndexCompressed.Create().SetPostingListCompressed(invertIndex.PostingListCompressed)
	}
	invertIndexesCreated, err := r.db.InvertIndexCompressed.CreateBulk(bulk...).Save(ctx)
	if err != nil {
		return nil, errors.NewError(code.Unknown, err)
	}
	return convertInvertIndexesEntSchemaToEntity(invertIndexesCreated), nil
}

func convertInvertIndexesEntSchemaToEntity(t []*ent.InvertIndexCompressed) *[]entities.InvertedIndexCompressedDetail {
	invertIndexes := make([]entities.InvertedIndexCompressedDetail, len(t))
	for i, item := range t {
		invertIndexes[i] = entities.InvertedIndexCompressedDetail{
			Uuid: item.ID,
			// TODO: Set Term Id
			PostingListCompressed: item.PostingListCompressed,
			CreatedAt:             item.CreatedAt,
			UpdatedAt:             item.UpdatedAt,
		}
	}
	return &invertIndexes
}