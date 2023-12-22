package main

import (
	"fmt"

	"github.com/brianvoe/gofakeit/v6"
	pb "github.com/yawkar/wbl0/pkg/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func mkGen(seed int64) func() *pb.Order {
	faker := gofakeit.New(seed)
	return func() *pb.Order {
		return genOrder(faker)
	}
}

func genOrder(faker *gofakeit.Faker) *pb.Order {
	orderUuid := faker.UUID()
	trackNumber := faker.Sentence(3)
	entry := faker.RandomString([]string{
		"WBIL",
		"WAUTO",
		"WHELM",
		"WIO",
	})
	locale := faker.RandomString([]string{
		"en",
		"ru",
		"br",
		"kz",
		"jp",
		"de",
	})
	items := genItems(faker, trackNumber)
	goodsTotal := func() int64 {
		var s int64 = 0
		for _, item := range items {
			s += item.TotalPrice
		}
		return s
	}()
	deliveryCost := int64(faker.IntRange(100, 2000))
	customFee := int64(faker.IntRange(0, 200))
	person := faker.Person()
	return &pb.Order{
		OrderUid:    orderUuid,
		TrackNumber: trackNumber,
		Entry:       entry,
		Delivery: &pb.Delivery{
			Name:    fmt.Sprintf("%s %s", person.FirstName, person.LastName),
			Phone:   person.Contact.Phone,
			Zip:     person.Address.Zip,
			City:    person.Address.City,
			Address: person.Address.Address,
			Region:  person.Address.Street,
			Email:   person.Contact.Email,
		},
		Payment: &pb.Payment{
			Transaction:  orderUuid,
			RequestId:    faker.UUID(),
			Currency:     faker.RandomString([]string{"USD", "RUB", "JPY", "EUR", "GBP", "CAD", "CHF", "CNY"}),
			Provider:     faker.RandomString([]string{"wbpay", "ekassa", "onlinekassa", "yamarkt"}),
			Amount:       deliveryCost + goodsTotal + customFee,
			PaymentDt:    int64(faker.NanoSecond()),
			Bank:         faker.RandomString([]string{"alpha", "sber", "tinkoff", "yabank", "otkrytie", "tochka"}),
			DeliveryCost: deliveryCost,
			GoodsTotal:   goodsTotal,
			CustomFee:    customFee,
		},
		Items:             items,
		Locale:            locale,
		InternalSignature: faker.HexUint256(),
		CustomerId:        faker.Username(),
		DeliveryService:   faker.Company(),
		ShardKey:          fmt.Sprint(faker.UintRange(1, 50)),
		SmId:              int64(faker.IntRange(1, 1e4)),
		DateCreated:       timestamppb.New(faker.FutureDate()),
		OofShard:          fmt.Sprint(faker.UintRange(1, 50)),
	}
}

func genItems(faker *gofakeit.Faker, trackNumber string) []*pb.Item {
	itemsCount := faker.IntRange(1, 4)
	items := make([]*pb.Item, itemsCount)
	for i := range items {
		items[i] = new(pb.Item)
		it := items[i]
		product := faker.Product()
		it.ChrtId = faker.Int64()
		it.TrackNumber = trackNumber
		it.Price = int64(product.Price)
		it.Rid = faker.UUID()
		it.Name = product.Name
		it.Sale = int64(faker.IntRange(0, 99))
		it.Size = fmt.Sprint(faker.IntRange(1, 4444))
		it.TotalPrice = int64(product.Price * (1.0 - float64(it.Sale) / 100.0))
		it.NmId = int64(faker.Uint32())
		it.Brand = faker.Company()
		it.Status = int32(faker.HTTPStatusCode())
	}
	return items
}
