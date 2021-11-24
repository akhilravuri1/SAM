package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"SAM/graph/generated"
	"SAM/graph/model"
	"context"
	"encoding/csv"
	"os"
)

var SAM_db = []*model.Sam{
	{
		ID:       "1",
		Name:     "SAM-1",
		Email:    "SAM1@gmail.com",
		Category: "Gold",
	},
	{
		ID:       "2",
		Name:     "SAM-2",
		Email:    "SAM2@gmail.com",
		Category: "Electronics",
	},
}

var Sellers_db = []*model.Seller{
	{
		ApplicationNo: "1231",
		Email:         "Seller1@gmail.com",
		Name:          "Seller-1",
		Status:        "PFA",
		SamID:         "1",
	},
	{
		ApplicationNo: "1232",
		Email:         "Seller2@gmail.com",
		Name:          "Seller-2",
		Status:        "PFA",
		SamID:         "1",
	},
	{
		ApplicationNo: "1233",
		Email:         "Seller3@gmail.com",
		Name:          "Seller-3",
		Status:        "PFA",
		SamID:         "1",
	},
	{
		ApplicationNo: "1234",
		Email:         "Seller4@gmail.com",
		Name:          "Seller-4",
		Status:        "PFA",
		SamID:         "2",
	},
	{
		ApplicationNo: "1235",
		Email:         "Seller5@gmail.com",
		Name:          "Seller-5",
		Status:        "PFA",
		SamID:         "2",
	},
}

func (r *mutationResolver) AcceptSeller(ctx context.Context, input *model.ChangeIn) (*model.Seller, error) {
	var res int
	for i, seller := range Sellers_db {
		if seller.ApplicationNo == input.ID {
			Sellers_db[i].Status = "Accepted"
			res = i
		}
	}
	return Sellers_db[res], nil
}

func (r *mutationResolver) RejectSeller(ctx context.Context, input *model.ChangeIn) (*model.Seller, error) {
	//panic(fmt.Errorf("not implemented"))
	var res int
	for i, seller := range Sellers_db {
		if seller.ApplicationNo == input.ID {
			Sellers_db[i].Status = "Rejected"
			res = i
		}
	}
	return Sellers_db[res], nil
}

func (r *mutationResolver) TransferSeller(ctx context.Context, input *model.TransferIn) (*model.Seller, error) {
	//panic(fmt.Errorf("not implemented"))
	var res int
	for i, seller := range Sellers_db {
		if seller.ApplicationNo == input.SellerID {
			Sellers_db[i].SamID = input.SamID
			res = i
		}
	}
	return Sellers_db[res], nil
}

func (r *mutationResolver) AcceptBulk(ctx context.Context, input []*model.ChangeIn) ([]*model.Seller, error) {
	//panic(fmt.Errorf("not implemented"))
	var res []*model.Seller
	for _, in := range input {
		for i, seller := range Sellers_db {
			if seller.ApplicationNo == in.ID {
				Sellers_db[i].Status = "Accepted"
				res = append(res, Sellers_db[i])
			}
		}
	}
	return res, nil
}

func (r *mutationResolver) UpdateSellerData(ctx context.Context, input *model.SellerDataIn) (*model.Seller, error) {
	//panic(fmt.Errorf("not implemented"))
	var res int
	for i, seller := range Sellers_db {
		if seller.ApplicationNo == input.SellerID {
			if input.Email != nil {
				Sellers_db[i].Email = *input.Email
			}
			if input.Name != nil {
				Sellers_db[i].Name = *input.Name
			}
			if input.SamID != nil {
				Sellers_db[i].SamID = *input.SamID
			}
			if input.Status != nil {
				Sellers_db[i].Status = *input.Status
			}
			res = i
		}
	}
	return Sellers_db[res], nil
}

func (r *queryResolver) GetAllSAMs(ctx context.Context) ([]*model.Sam, error) {
	//panic(fmt.Errorf("not implemented"))
	return SAM_db, nil
}

func (r *queryResolver) GetAllSellers(ctx context.Context) ([]*model.Seller, error) {
	//panic(fmt.Errorf("not implemented"))
	return Sellers_db, nil
}

func (r *queryResolver) SellersByStatus(ctx context.Context, input *model.StatusIn) ([]*model.Seller, error) {
	//panic(fmt.Errorf("not implemented"))
	var SellerByStatusTemp []*model.Seller
	for _, seller := range Sellers_db {
		if seller.Status == input.Status {
			SellerByStatusTemp = append(SellerByStatusTemp, seller)
		}
	}
	return SellerByStatusTemp, nil
}

func (r *queryResolver) SellersBySam(ctx context.Context, input *model.SamIn) ([]*model.Seller, error) {
	//panic(fmt.Errorf("not implemented"))
	var SellersBySamTemp []*model.Seller
	for _, seller := range Sellers_db {
		if seller.SamID == input.ID {
			SellersBySamTemp = append(SellersBySamTemp, seller)
		}
	}
	return SellersBySamTemp, nil
}

func (r *queryResolver) DownloadData(ctx context.Context) ([]*model.Seller, error) {
	//panic(fmt.Errorf("not implemented"))
	file, err := os.Create("records.csv")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	w := csv.NewWriter(file)
	defer w.Flush()

	header := []string{"ApplicationNo", "NAME", "EMAIL", "Status", "SamID"}
	w.Write(header)

	// Using WriteAll
	var data [][]string
	for _, record := range Sellers_db {
		row := []string{record.ApplicationNo, record.Name, record.Email, record.Status, record.SamID}
		data = append(data, row)
	}
	w.WriteAll(data)
	return Sellers_db, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
