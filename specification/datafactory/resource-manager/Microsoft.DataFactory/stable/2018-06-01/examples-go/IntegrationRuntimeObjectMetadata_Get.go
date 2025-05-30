package armdatafactory_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v10"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/8eb3f7a4f66d408152c32b9d647e59147172d533/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/IntegrationRuntimeObjectMetadata_Get.json
func ExampleIntegrationRuntimeObjectMetadataClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatafactory.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewIntegrationRuntimeObjectMetadataClient().Get(ctx, "exampleResourceGroup", "exampleFactoryName", "testactivityv2", &armdatafactory.IntegrationRuntimeObjectMetadataClientGetOptions{GetMetadataRequest: &armdatafactory.GetSsisObjectMetadataRequest{
		MetadataPath: to.Ptr("ssisFolders"),
	},
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SsisObjectMetadataListResponse = armdatafactory.SsisObjectMetadataListResponse{
	// 	Value: []armdatafactory.SsisObjectMetadataClassification{
	// 		&armdatafactory.SsisFolder{
	// 			Name: to.Ptr("TestFolder"),
	// 			Type: to.Ptr(armdatafactory.SsisObjectMetadataTypeFolder),
	// 			Description: to.Ptr(""),
	// 			ID: to.Ptr[int64](1),
	// 		},
	// 		&armdatafactory.SsisFolder{
	// 			Name: to.Ptr("EnvironmentFolder"),
	// 			Type: to.Ptr(armdatafactory.SsisObjectMetadataTypeFolder),
	// 			Description: to.Ptr(""),
	// 			ID: to.Ptr[int64](2),
	// 		},
	// 		&armdatafactory.SsisFolder{
	// 			Name: to.Ptr("ActivityTest"),
	// 			Type: to.Ptr(armdatafactory.SsisObjectMetadataTypeFolder),
	// 			Description: to.Ptr(""),
	// 			ID: to.Ptr[int64](3),
	// 		},
	// 		&armdatafactory.SsisFolder{
	// 			Name: to.Ptr("这是文件夹"),
	// 			Type: to.Ptr(armdatafactory.SsisObjectMetadataTypeFolder),
	// 			Description: to.Ptr(""),
	// 			ID: to.Ptr[int64](4),
	// 		},
	// 		&armdatafactory.SsisFolder{
	// 			Name: to.Ptr("1"),
	// 			Type: to.Ptr(armdatafactory.SsisObjectMetadataTypeFolder),
	// 			ID: to.Ptr[int64](5),
	// 		},
	// 		&armdatafactory.SsisFolder{
	// 			Name: to.Ptr("2"),
	// 			Type: to.Ptr(armdatafactory.SsisObjectMetadataTypeFolder),
	// 			ID: to.Ptr[int64](6),
	// 		},
	// 		&armdatafactory.SsisFolder{
	// 			Name: to.Ptr("3"),
	// 			Type: to.Ptr(armdatafactory.SsisObjectMetadataTypeFolder),
	// 			ID: to.Ptr[int64](7),
	// 		},
	// 		&armdatafactory.SsisFolder{
	// 			Name: to.Ptr("4"),
	// 			Type: to.Ptr(armdatafactory.SsisObjectMetadataTypeFolder),
	// 			ID: to.Ptr[int64](8),
	// 		},
	// 		&armdatafactory.SsisFolder{
	// 			Name: to.Ptr("5"),
	// 			Type: to.Ptr(armdatafactory.SsisObjectMetadataTypeFolder),
	// 			ID: to.Ptr[int64](9),
	// 		},
	// 		&armdatafactory.SsisFolder{
	// 			Name: to.Ptr("6"),
	// 			Type: to.Ptr(armdatafactory.SsisObjectMetadataTypeFolder),
	// 			ID: to.Ptr[int64](10),
	// 		},
	// 		&armdatafactory.SsisFolder{
	// 			Name: to.Ptr("7"),
	// 			Type: to.Ptr(armdatafactory.SsisObjectMetadataTypeFolder),
	// 			ID: to.Ptr[int64](11),
	// 		},
	// 		&armdatafactory.SsisFolder{
	// 			Name: to.Ptr("8"),
	// 			Type: to.Ptr(armdatafactory.SsisObjectMetadataTypeFolder),
	// 			ID: to.Ptr[int64](12),
	// 		},
	// 		&armdatafactory.SsisFolder{
	// 			Name: to.Ptr("9"),
	// 			Type: to.Ptr(armdatafactory.SsisObjectMetadataTypeFolder),
	// 			ID: to.Ptr[int64](13),
	// 		},
	// 		&armdatafactory.SsisFolder{
	// 			Name: to.Ptr("10"),
	// 			Type: to.Ptr(armdatafactory.SsisObjectMetadataTypeFolder),
	// 			ID: to.Ptr[int64](14),
	// 		},
	// 		&armdatafactory.SsisFolder{
	// 			Name: to.Ptr("11"),
	// 			Type: to.Ptr(armdatafactory.SsisObjectMetadataTypeFolder),
	// 			ID: to.Ptr[int64](15),
	// 		},
	// 		&armdatafactory.SsisFolder{
	// 			Name: to.Ptr("12"),
	// 			Type: to.Ptr(armdatafactory.SsisObjectMetadataTypeFolder),
	// 			ID: to.Ptr[int64](16),
	// 		},
	// 		&armdatafactory.SsisFolder{
	// 			Name: to.Ptr("13"),
	// 			Type: to.Ptr(armdatafactory.SsisObjectMetadataTypeFolder),
	// 			ID: to.Ptr[int64](17),
	// 		},
	// 		&armdatafactory.SsisFolder{
	// 			Name: to.Ptr("14"),
	// 			Type: to.Ptr(armdatafactory.SsisObjectMetadataTypeFolder),
	// 			ID: to.Ptr[int64](18),
	// 		},
	// 		&armdatafactory.SsisFolder{
	// 			Name: to.Ptr("15"),
	// 			Type: to.Ptr(armdatafactory.SsisObjectMetadataTypeFolder),
	// 			ID: to.Ptr[int64](19),
	// 		},
	// 		&armdatafactory.SsisFolder{
	// 			Name: to.Ptr("16"),
	// 			Type: to.Ptr(armdatafactory.SsisObjectMetadataTypeFolder),
	// 			ID: to.Ptr[int64](20),
	// 		},
	// 		&armdatafactory.SsisFolder{
	// 			Name: to.Ptr("17"),
	// 			Type: to.Ptr(armdatafactory.SsisObjectMetadataTypeFolder),
	// 			ID: to.Ptr[int64](21),
	// 		},
	// 		&armdatafactory.SsisFolder{
	// 			Name: to.Ptr("18"),
	// 			Type: to.Ptr(armdatafactory.SsisObjectMetadataTypeFolder),
	// 			ID: to.Ptr[int64](22),
	// 		},
	// 		&armdatafactory.SsisFolder{
	// 			Name: to.Ptr("19"),
	// 			Type: to.Ptr(armdatafactory.SsisObjectMetadataTypeFolder),
	// 			ID: to.Ptr[int64](23),
	// 		},
	// 		&armdatafactory.SsisFolder{
	// 			Name: to.Ptr("20"),
	// 			Type: to.Ptr(armdatafactory.SsisObjectMetadataTypeFolder),
	// 			ID: to.Ptr[int64](24),
	// 		},
	// 		&armdatafactory.SsisFolder{
	// 			Name: to.Ptr("21"),
	// 			Type: to.Ptr(armdatafactory.SsisObjectMetadataTypeFolder),
	// 			ID: to.Ptr[int64](25),
	// 		},
	// 		&armdatafactory.SsisFolder{
	// 			Name: to.Ptr("22"),
	// 			Type: to.Ptr(armdatafactory.SsisObjectMetadataTypeFolder),
	// 			ID: to.Ptr[int64](26),
	// 		},
	// 		&armdatafactory.SsisFolder{
	// 			Name: to.Ptr("23"),
	// 			Type: to.Ptr(armdatafactory.SsisObjectMetadataTypeFolder),
	// 			ID: to.Ptr[int64](27),
	// 		},
	// 		&armdatafactory.SsisFolder{
	// 			Name: to.Ptr("24"),
	// 			Type: to.Ptr(armdatafactory.SsisObjectMetadataTypeFolder),
	// 			ID: to.Ptr[int64](28),
	// 		},
	// 		&armdatafactory.SsisFolder{
	// 			Name: to.Ptr("25"),
	// 			Type: to.Ptr(armdatafactory.SsisObjectMetadataTypeFolder),
	// 			ID: to.Ptr[int64](29),
	// 		},
	// 		&armdatafactory.SsisFolder{
	// 			Name: to.Ptr("26"),
	// 			Type: to.Ptr(armdatafactory.SsisObjectMetadataTypeFolder),
	// 			ID: to.Ptr[int64](30),
	// 		},
	// 		&armdatafactory.SsisFolder{
	// 			Name: to.Ptr("27"),
	// 			Type: to.Ptr(armdatafactory.SsisObjectMetadataTypeFolder),
	// 			ID: to.Ptr[int64](31),
	// 		},
	// 		&armdatafactory.SsisFolder{
	// 			Name: to.Ptr("28"),
	// 			Type: to.Ptr(armdatafactory.SsisObjectMetadataTypeFolder),
	// 			ID: to.Ptr[int64](32),
	// 		},
	// 		&armdatafactory.SsisFolder{
	// 			Name: to.Ptr("29"),
	// 			Type: to.Ptr(armdatafactory.SsisObjectMetadataTypeFolder),
	// 			ID: to.Ptr[int64](33),
	// 		},
	// 		&armdatafactory.SsisFolder{
	// 			Name: to.Ptr("30"),
	// 			Type: to.Ptr(armdatafactory.SsisObjectMetadataTypeFolder),
	// 			ID: to.Ptr[int64](34),
	// 		},
	// 		&armdatafactory.SsisFolder{
	// 			Name: to.Ptr("31"),
	// 			Type: to.Ptr(armdatafactory.SsisObjectMetadataTypeFolder),
	// 			ID: to.Ptr[int64](35),
	// 		},
	// 		&armdatafactory.SsisFolder{
	// 			Name: to.Ptr("32"),
	// 			Type: to.Ptr(armdatafactory.SsisObjectMetadataTypeFolder),
	// 			ID: to.Ptr[int64](36),
	// 		},
	// 		&armdatafactory.SsisFolder{
	// 			Name: to.Ptr("33"),
	// 			Type: to.Ptr(armdatafactory.SsisObjectMetadataTypeFolder),
	// 			ID: to.Ptr[int64](37),
	// 		},
	// 		&armdatafactory.SsisFolder{
	// 			Name: to.Ptr("34"),
	// 			Type: to.Ptr(armdatafactory.SsisObjectMetadataTypeFolder),
	// 			ID: to.Ptr[int64](38),
	// 		},
	// 		&armdatafactory.SsisFolder{
	// 			Name: to.Ptr("35"),
	// 			Type: to.Ptr(armdatafactory.SsisObjectMetadataTypeFolder),
	// 			ID: to.Ptr[int64](39),
	// 		},
	// 		&armdatafactory.SsisFolder{
	// 			Name: to.Ptr("36"),
	// 			Type: to.Ptr(armdatafactory.SsisObjectMetadataTypeFolder),
	// 			ID: to.Ptr[int64](40),
	// 		},
	// 		&armdatafactory.SsisFolder{
	// 			Name: to.Ptr("37"),
	// 			Type: to.Ptr(armdatafactory.SsisObjectMetadataTypeFolder),
	// 			ID: to.Ptr[int64](41),
	// 		},
	// 		&armdatafactory.SsisFolder{
	// 			Name: to.Ptr("38"),
	// 			Type: to.Ptr(armdatafactory.SsisObjectMetadataTypeFolder),
	// 			ID: to.Ptr[int64](42),
	// 		},
	// 		&armdatafactory.SsisFolder{
	// 			Name: to.Ptr("39"),
	// 			Type: to.Ptr(armdatafactory.SsisObjectMetadataTypeFolder),
	// 			ID: to.Ptr[int64](43),
	// 		},
	// 		&armdatafactory.SsisFolder{
	// 			Name: to.Ptr("40"),
	// 			Type: to.Ptr(armdatafactory.SsisObjectMetadataTypeFolder),
	// 			ID: to.Ptr[int64](44),
	// 		},
	// 		&armdatafactory.SsisFolder{
	// 			Name: to.Ptr("41"),
	// 			Type: to.Ptr(armdatafactory.SsisObjectMetadataTypeFolder),
	// 			ID: to.Ptr[int64](45),
	// 		},
	// 		&armdatafactory.SsisFolder{
	// 			Name: to.Ptr("42"),
	// 			Type: to.Ptr(armdatafactory.SsisObjectMetadataTypeFolder),
	// 			ID: to.Ptr[int64](46),
	// 		},
	// 		&armdatafactory.SsisFolder{
	// 			Name: to.Ptr("43"),
	// 			Type: to.Ptr(armdatafactory.SsisObjectMetadataTypeFolder),
	// 			ID: to.Ptr[int64](47),
	// 		},
	// 		&armdatafactory.SsisFolder{
	// 			Name: to.Ptr("44"),
	// 			Type: to.Ptr(armdatafactory.SsisObjectMetadataTypeFolder),
	// 			ID: to.Ptr[int64](48),
	// 		},
	// 		&armdatafactory.SsisFolder{
	// 			Name: to.Ptr("45"),
	// 			Type: to.Ptr(armdatafactory.SsisObjectMetadataTypeFolder),
	// 			ID: to.Ptr[int64](49),
	// 		},
	// 		&armdatafactory.SsisFolder{
	// 			Name: to.Ptr("46"),
	// 			Type: to.Ptr(armdatafactory.SsisObjectMetadataTypeFolder),
	// 			ID: to.Ptr[int64](50),
	// 		},
	// 		&armdatafactory.SsisFolder{
	// 			Name: to.Ptr("47"),
	// 			Type: to.Ptr(armdatafactory.SsisObjectMetadataTypeFolder),
	// 			ID: to.Ptr[int64](51),
	// 		},
	// 		&armdatafactory.SsisFolder{
	// 			Name: to.Ptr("48"),
	// 			Type: to.Ptr(armdatafactory.SsisObjectMetadataTypeFolder),
	// 			ID: to.Ptr[int64](52),
	// 		},
	// 		&armdatafactory.SsisFolder{
	// 			Name: to.Ptr("49"),
	// 			Type: to.Ptr(armdatafactory.SsisObjectMetadataTypeFolder),
	// 			ID: to.Ptr[int64](53),
	// 		},
	// 		&armdatafactory.SsisFolder{
	// 			Name: to.Ptr("50"),
	// 			Type: to.Ptr(armdatafactory.SsisObjectMetadataTypeFolder),
	// 			ID: to.Ptr[int64](54),
	// 		},
	// 		&armdatafactory.SsisFolder{
	// 			Name: to.Ptr("51"),
	// 			Type: to.Ptr(armdatafactory.SsisObjectMetadataTypeFolder),
	// 			ID: to.Ptr[int64](55),
	// 		},
	// 		&armdatafactory.SsisFolder{
	// 			Name: to.Ptr("52"),
	// 			Type: to.Ptr(armdatafactory.SsisObjectMetadataTypeFolder),
	// 			ID: to.Ptr[int64](56),
	// 		},
	// 		&armdatafactory.SsisFolder{
	// 			Name: to.Ptr("53"),
	// 			Type: to.Ptr(armdatafactory.SsisObjectMetadataTypeFolder),
	// 			ID: to.Ptr[int64](57),
	// 		},
	// 		&armdatafactory.SsisFolder{
	// 			Name: to.Ptr("54"),
	// 			Type: to.Ptr(armdatafactory.SsisObjectMetadataTypeFolder),
	// 			ID: to.Ptr[int64](58),
	// 		},
	// 		&armdatafactory.SsisFolder{
	// 			Name: to.Ptr("55"),
	// 			Type: to.Ptr(armdatafactory.SsisObjectMetadataTypeFolder),
	// 			ID: to.Ptr[int64](59),
	// 		},
	// 		&armdatafactory.SsisFolder{
	// 			Name: to.Ptr("56"),
	// 			Type: to.Ptr(armdatafactory.SsisObjectMetadataTypeFolder),
	// 			ID: to.Ptr[int64](60),
	// 		},
	// 		&armdatafactory.SsisFolder{
	// 			Name: to.Ptr("57"),
	// 			Type: to.Ptr(armdatafactory.SsisObjectMetadataTypeFolder),
	// 			ID: to.Ptr[int64](61),
	// 		},
	// 		&armdatafactory.SsisFolder{
	// 			Name: to.Ptr("58"),
	// 			Type: to.Ptr(armdatafactory.SsisObjectMetadataTypeFolder),
	// 			ID: to.Ptr[int64](62),
	// 		},
	// 		&armdatafactory.SsisFolder{
	// 			Name: to.Ptr("59"),
	// 			Type: to.Ptr(armdatafactory.SsisObjectMetadataTypeFolder),
	// 			ID: to.Ptr[int64](63),
	// 		},
	// 		&armdatafactory.SsisFolder{
	// 			Name: to.Ptr("60"),
	// 			Type: to.Ptr(armdatafactory.SsisObjectMetadataTypeFolder),
	// 			ID: to.Ptr[int64](64),
	// 		},
	// 		&armdatafactory.SsisFolder{
	// 			Name: to.Ptr("61"),
	// 			Type: to.Ptr(armdatafactory.SsisObjectMetadataTypeFolder),
	// 			ID: to.Ptr[int64](65),
	// 		},
	// 		&armdatafactory.SsisFolder{
	// 			Name: to.Ptr("62"),
	// 			Type: to.Ptr(armdatafactory.SsisObjectMetadataTypeFolder),
	// 			ID: to.Ptr[int64](66),
	// 		},
	// 		&armdatafactory.SsisFolder{
	// 			Name: to.Ptr("63"),
	// 			Type: to.Ptr(armdatafactory.SsisObjectMetadataTypeFolder),
	// 			ID: to.Ptr[int64](67),
	// 		},
	// 		&armdatafactory.SsisFolder{
	// 			Name: to.Ptr("64"),
	// 			Type: to.Ptr(armdatafactory.SsisObjectMetadataTypeFolder),
	// 			ID: to.Ptr[int64](68),
	// 		},
	// 		&armdatafactory.SsisFolder{
	// 			Name: to.Ptr("65"),
	// 			Type: to.Ptr(armdatafactory.SsisObjectMetadataTypeFolder),
	// 			ID: to.Ptr[int64](69),
	// 		},
	// 		&armdatafactory.SsisFolder{
	// 			Name: to.Ptr("66"),
	// 			Type: to.Ptr(armdatafactory.SsisObjectMetadataTypeFolder),
	// 			ID: to.Ptr[int64](70),
	// 		},
	// 		&armdatafactory.SsisFolder{
	// 			Name: to.Ptr("67"),
	// 			Type: to.Ptr(armdatafactory.SsisObjectMetadataTypeFolder),
	// 			ID: to.Ptr[int64](71),
	// 		},
	// 		&armdatafactory.SsisFolder{
	// 			Name: to.Ptr("68"),
	// 			Type: to.Ptr(armdatafactory.SsisObjectMetadataTypeFolder),
	// 			ID: to.Ptr[int64](72),
	// 		},
	// 		&armdatafactory.SsisFolder{
	// 			Name: to.Ptr("69"),
	// 			Type: to.Ptr(armdatafactory.SsisObjectMetadataTypeFolder),
	// 			ID: to.Ptr[int64](73),
	// 		},
	// 		&armdatafactory.SsisFolder{
	// 			Name: to.Ptr("70"),
	// 			Type: to.Ptr(armdatafactory.SsisObjectMetadataTypeFolder),
	// 			ID: to.Ptr[int64](74),
	// 		},
	// 		&armdatafactory.SsisFolder{
	// 			Name: to.Ptr("71"),
	// 			Type: to.Ptr(armdatafactory.SsisObjectMetadataTypeFolder),
	// 			ID: to.Ptr[int64](75),
	// 		},
	// 		&armdatafactory.SsisFolder{
	// 			Name: to.Ptr("72"),
	// 			Type: to.Ptr(armdatafactory.SsisObjectMetadataTypeFolder),
	// 			ID: to.Ptr[int64](76),
	// 		},
	// 		&armdatafactory.SsisFolder{
	// 			Name: to.Ptr("73"),
	// 			Type: to.Ptr(armdatafactory.SsisObjectMetadataTypeFolder),
	// 			ID: to.Ptr[int64](77),
	// 		},
	// 		&armdatafactory.SsisFolder{
	// 			Name: to.Ptr("74"),
	// 			Type: to.Ptr(armdatafactory.SsisObjectMetadataTypeFolder),
	// 			ID: to.Ptr[int64](78),
	// 		},
	// 		&armdatafactory.SsisFolder{
	// 			Name: to.Ptr("75"),
	// 			Type: to.Ptr(armdatafactory.SsisObjectMetadataTypeFolder),
	// 			ID: to.Ptr[int64](79),
	// 		},
	// 		&armdatafactory.SsisFolder{
	// 			Name: to.Ptr("76"),
	// 			Type: to.Ptr(armdatafactory.SsisObjectMetadataTypeFolder),
	// 			ID: to.Ptr[int64](80),
	// 		},
	// 		&armdatafactory.SsisFolder{
	// 			Name: to.Ptr("77"),
	// 			Type: to.Ptr(armdatafactory.SsisObjectMetadataTypeFolder),
	// 			ID: to.Ptr[int64](81),
	// 		},
	// 		&armdatafactory.SsisFolder{
	// 			Name: to.Ptr("78"),
	// 			Type: to.Ptr(armdatafactory.SsisObjectMetadataTypeFolder),
	// 			ID: to.Ptr[int64](82),
	// 		},
	// 		&armdatafactory.SsisFolder{
	// 			Name: to.Ptr("79"),
	// 			Type: to.Ptr(armdatafactory.SsisObjectMetadataTypeFolder),
	// 			ID: to.Ptr[int64](83),
	// 		},
	// 		&armdatafactory.SsisFolder{
	// 			Name: to.Ptr("80"),
	// 			Type: to.Ptr(armdatafactory.SsisObjectMetadataTypeFolder),
	// 			ID: to.Ptr[int64](84),
	// 		},
	// 		&armdatafactory.SsisFolder{
	// 			Name: to.Ptr("81"),
	// 			Type: to.Ptr(armdatafactory.SsisObjectMetadataTypeFolder),
	// 			ID: to.Ptr[int64](85),
	// 		},
	// 		&armdatafactory.SsisFolder{
	// 			Name: to.Ptr("82"),
	// 			Type: to.Ptr(armdatafactory.SsisObjectMetadataTypeFolder),
	// 			ID: to.Ptr[int64](86),
	// 		},
	// 		&armdatafactory.SsisFolder{
	// 			Name: to.Ptr("83"),
	// 			Type: to.Ptr(armdatafactory.SsisObjectMetadataTypeFolder),
	// 			ID: to.Ptr[int64](87),
	// 		},
	// 		&armdatafactory.SsisFolder{
	// 			Name: to.Ptr("84"),
	// 			Type: to.Ptr(armdatafactory.SsisObjectMetadataTypeFolder),
	// 			ID: to.Ptr[int64](88),
	// 		},
	// 		&armdatafactory.SsisFolder{
	// 			Name: to.Ptr("85"),
	// 			Type: to.Ptr(armdatafactory.SsisObjectMetadataTypeFolder),
	// 			ID: to.Ptr[int64](89),
	// 		},
	// 		&armdatafactory.SsisFolder{
	// 			Name: to.Ptr("86"),
	// 			Type: to.Ptr(armdatafactory.SsisObjectMetadataTypeFolder),
	// 			ID: to.Ptr[int64](90),
	// 		},
	// 		&armdatafactory.SsisFolder{
	// 			Name: to.Ptr("87"),
	// 			Type: to.Ptr(armdatafactory.SsisObjectMetadataTypeFolder),
	// 			ID: to.Ptr[int64](91),
	// 		},
	// 		&armdatafactory.SsisFolder{
	// 			Name: to.Ptr("88"),
	// 			Type: to.Ptr(armdatafactory.SsisObjectMetadataTypeFolder),
	// 			ID: to.Ptr[int64](92),
	// 		},
	// 		&armdatafactory.SsisFolder{
	// 			Name: to.Ptr("89"),
	// 			Type: to.Ptr(armdatafactory.SsisObjectMetadataTypeFolder),
	// 			ID: to.Ptr[int64](93),
	// 		},
	// 		&armdatafactory.SsisFolder{
	// 			Name: to.Ptr("90"),
	// 			Type: to.Ptr(armdatafactory.SsisObjectMetadataTypeFolder),
	// 			ID: to.Ptr[int64](94),
	// 		},
	// 		&armdatafactory.SsisFolder{
	// 			Name: to.Ptr("91"),
	// 			Type: to.Ptr(armdatafactory.SsisObjectMetadataTypeFolder),
	// 			ID: to.Ptr[int64](95),
	// 		},
	// 		&armdatafactory.SsisFolder{
	// 			Name: to.Ptr("92"),
	// 			Type: to.Ptr(armdatafactory.SsisObjectMetadataTypeFolder),
	// 			ID: to.Ptr[int64](96),
	// 		},
	// 		&armdatafactory.SsisFolder{
	// 			Name: to.Ptr("93"),
	// 			Type: to.Ptr(armdatafactory.SsisObjectMetadataTypeFolder),
	// 			ID: to.Ptr[int64](97),
	// 		},
	// 		&armdatafactory.SsisFolder{
	// 			Name: to.Ptr("94"),
	// 			Type: to.Ptr(armdatafactory.SsisObjectMetadataTypeFolder),
	// 			ID: to.Ptr[int64](98),
	// 		},
	// 		&armdatafactory.SsisFolder{
	// 			Name: to.Ptr("95"),
	// 			Type: to.Ptr(armdatafactory.SsisObjectMetadataTypeFolder),
	// 			ID: to.Ptr[int64](99),
	// 		},
	// 		&armdatafactory.SsisFolder{
	// 			Name: to.Ptr("96"),
	// 			Type: to.Ptr(armdatafactory.SsisObjectMetadataTypeFolder),
	// 			ID: to.Ptr[int64](100),
	// 	}},
	// }
}
