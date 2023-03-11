package handler

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"litemall-srvs/goods_srv/global"
	"litemall-srvs/goods_srv/model"
	"litemall-srvs/goods_srv/proto"
)

//品牌和轮播图

func (s *GoodsServer) BrandList(ctx context.Context, request *proto.BrandFilterRequest) (*proto.BrandListResponse, error) {
	//TODO implement me
	brandListResponse := proto.BrandListResponse{}

	var brands []model.Brand
	result := global.DB.Scopes(Paginate(int(request.Pages), int(request.PagePerNums))).Find(&brands)
	if result.Error != nil {
		return nil, result.Error
	}

	var total int64
	global.DB.Model(&model.Brand{}).Count(&total)
	brandListResponse.Total = int32(total)

	var brandResponses []*proto.BrandInfoResponse
	for _, brand := range brands {
		brandResponses = append(brandResponses, &proto.BrandInfoResponse{
			Id:   brand.ID,
			Name: brand.Name,
			Logo: brand.PicUrl,
		})
	}
	brandListResponse.Data = brandResponses
	return &brandListResponse, nil
}

//新建品牌

func (s *GoodsServer) CreateBrand(ctx context.Context, request *proto.BrandRequest) (*proto.BrandInfoResponse, error) {
	//TODO implement me
	if result := global.DB.Where("name=?", request.Name).First(&model.Brand{}); result.RowsAffected == 1 {
		return nil, status.Errorf(codes.InvalidArgument, "品牌已存在")
	}

	brand := &model.Brand{
		Name:   request.Name,
		PicUrl: request.Logo,
	}
	global.DB.Save(brand)

	return &proto.BrandInfoResponse{Id: brand.ID}, nil
}

func (s *GoodsServer) DeleteBrand(ctx context.Context, request *proto.BrandRequest) (*emptypb.Empty, error) {
	//TODO implement me
	if result := global.DB.Delete(&model.Brand{}, request.Id); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "品牌不存在")
	}
	return &emptypb.Empty{}, nil
}

func (s *GoodsServer) UpdateBrand(ctx context.Context, request *proto.BrandRequest) (*emptypb.Empty, error) {
	//TODO implement me
	brand := model.Brand{}
	if result := global.DB.First(&brand); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "品牌不存在")
	}

	if request.Name != "" {
		brand.Name = request.Name
	}
	if request.Logo != "" {
		brand.PicUrl = request.Logo
	}

	global.DB.Save(&brand)

	return &emptypb.Empty{}, nil
}
