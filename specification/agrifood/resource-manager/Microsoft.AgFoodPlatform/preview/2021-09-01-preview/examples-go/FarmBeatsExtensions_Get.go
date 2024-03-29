package armagrifood_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/agrifood/armagrifood"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a7af6049f4b4743ef3b649f3852bcc7bd9a43ee0/specification/agrifood/resource-manager/Microsoft.AgFoodPlatform/preview/2021-09-01-preview/examples/FarmBeatsExtensions_Get.json
func ExampleFarmBeatsExtensionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armagrifood.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFarmBeatsExtensionsClient().Get(ctx, "DTN.ContentServices", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.FarmBeatsExtension = armagrifood.FarmBeatsExtension{
	// 	Name: to.Ptr("DTN.ContentServices"),
	// 	Type: to.Ptr("Microsoft.AgFoodPlatform/farmBeatsExtensionDefinitions"),
	// 	ID: to.Ptr("Microsoft.AgFoodPlatform/farmBeatsExtensionDefinitions/DTN.ContentServices"),
	// 	SystemData: &armagrifood.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-12T15:28:06.000Z"); return t}()),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-12T15:30:01.000Z"); return t}()),
	// 	},
	// 	Properties: &armagrifood.FarmBeatsExtensionProperties{
	// 		DetailedInformation: []*armagrifood.DetailedInformation{
	// 			{
	// 				APIInputParameters: []*string{
	// 					to.Ptr("stationId"),
	// 					to.Ptr("lat"),
	// 					to.Ptr("lon"),
	// 					to.Ptr("days"),
	// 					to.Ptr("units"),
	// 					to.Ptr("precision"),
	// 					to.Ptr("sector")},
	// 					APIName: to.Ptr("GetDailyObservations"),
	// 					CustomParameters: []*string{
	// 						to.Ptr("stationId"),
	// 						to.Ptr("stationLatitude"),
	// 						to.Ptr("stationLongitude"),
	// 						to.Ptr("timeZone"),
	// 						to.Ptr("sunrise"),
	// 						to.Ptr("sunset"),
	// 						to.Ptr("weatherCode"),
	// 						to.Ptr("weatherDescription"),
	// 						to.Ptr("maxTemperature"),
	// 						to.Ptr("minTemperature"),
	// 						to.Ptr("avgHeatIndex"),
	// 						to.Ptr("maxHeatIndex"),
	// 						to.Ptr("minHeatIndex"),
	// 						to.Ptr("maxWindChill"),
	// 						to.Ptr("minWindChill"),
	// 						to.Ptr("maxFeelsLike"),
	// 						to.Ptr("minFeelsLike"),
	// 						to.Ptr("avgFeelsLike"),
	// 						to.Ptr("maxWindSpeed"),
	// 						to.Ptr("avgWetBulbGlobeTemp"),
	// 						to.Ptr("maxWetBulbGlobeTemp"),
	// 						to.Ptr("minWetBulbGlobeTemp"),
	// 						to.Ptr("minutesOfSunshine"),
	// 						to.Ptr("cornHeatUnit"),
	// 						to.Ptr("evapotranspiration")},
	// 						PlatformParameters: []*string{
	// 							to.Ptr("cloudCover"),
	// 							to.Ptr("dewPoint"),
	// 							to.Ptr("growingDegreeDay"),
	// 							to.Ptr("precipitation"),
	// 							to.Ptr("pressure"),
	// 							to.Ptr("relativeHumidity"),
	// 							to.Ptr("temperature"),
	// 							to.Ptr("wetBulbTemperature"),
	// 							to.Ptr("dateTime"),
	// 							to.Ptr("windChill"),
	// 							to.Ptr("windSpeed"),
	// 							to.Ptr("windDirection")},
	// 							UnitsSupported: &armagrifood.UnitSystemsInfo{
	// 								Key: to.Ptr("units"),
	// 								Values: []*string{
	// 									to.Ptr("us"),
	// 									to.Ptr("si")},
	// 								},
	// 							},
	// 							{
	// 								APIInputParameters: []*string{
	// 									to.Ptr("stationId"),
	// 									to.Ptr("lat"),
	// 									to.Ptr("lon"),
	// 									to.Ptr("hours"),
	// 									to.Ptr("units"),
	// 									to.Ptr("precision"),
	// 									to.Ptr("sector")},
	// 									APIName: to.Ptr("GetHourlyObservations"),
	// 									CustomParameters: []*string{
	// 										to.Ptr("stationId"),
	// 										to.Ptr("stationLatitude"),
	// 										to.Ptr("stationLongitude"),
	// 										to.Ptr("timeZone"),
	// 										to.Ptr("weatherCode"),
	// 										to.Ptr("weatherDescription"),
	// 										to.Ptr("feelsLike"),
	// 										to.Ptr("visibilityWeatherCode"),
	// 										to.Ptr("visibilityWeatherDescription"),
	// 										to.Ptr("minutesOfSunshine")},
	// 										PlatformParameters: []*string{
	// 											to.Ptr("cloudCover"),
	// 											to.Ptr("dewPoint"),
	// 											to.Ptr("precipitation"),
	// 											to.Ptr("pressure"),
	// 											to.Ptr("relativeHumidity"),
	// 											to.Ptr("temperature"),
	// 											to.Ptr("wetBulbTemperature"),
	// 											to.Ptr("dateTime"),
	// 											to.Ptr("visibility"),
	// 											to.Ptr("windChill"),
	// 											to.Ptr("windSpeed"),
	// 											to.Ptr("windDirection"),
	// 											to.Ptr("windGust")},
	// 											UnitsSupported: &armagrifood.UnitSystemsInfo{
	// 												Key: to.Ptr("units"),
	// 												Values: []*string{
	// 													to.Ptr("us"),
	// 													to.Ptr("si")},
	// 												},
	// 											},
	// 											{
	// 												APIInputParameters: []*string{
	// 													to.Ptr("stationId"),
	// 													to.Ptr("lat"),
	// 													to.Ptr("lon"),
	// 													to.Ptr("days"),
	// 													to.Ptr("units"),
	// 													to.Ptr("precision"),
	// 													to.Ptr("sector")},
	// 													APIName: to.Ptr("GetHourlyForecasts"),
	// 													CustomParameters: []*string{
	// 														to.Ptr("stationId"),
	// 														to.Ptr("stationLatitude"),
	// 														to.Ptr("stationLongitude"),
	// 														to.Ptr("timeZone"),
	// 														to.Ptr("weatherCode"),
	// 														to.Ptr("weatherDescription"),
	// 														to.Ptr("feelsLike"),
	// 														to.Ptr("visibilityWeatherCode"),
	// 														to.Ptr("visibilityWeatherDescription"),
	// 														to.Ptr("minutesOfSunshine")},
	// 														PlatformParameters: []*string{
	// 															to.Ptr("cloudCover"),
	// 															to.Ptr("dewPoint"),
	// 															to.Ptr("precipitation"),
	// 															to.Ptr("pressure"),
	// 															to.Ptr("relativeHumidity"),
	// 															to.Ptr("temperature"),
	// 															to.Ptr("wetBulbTemperature"),
	// 															to.Ptr("dateTime"),
	// 															to.Ptr("visibility"),
	// 															to.Ptr("windChill"),
	// 															to.Ptr("windSpeed"),
	// 															to.Ptr("windDirection"),
	// 															to.Ptr("windGust")},
	// 															UnitsSupported: &armagrifood.UnitSystemsInfo{
	// 																Key: to.Ptr("units"),
	// 																Values: []*string{
	// 																	to.Ptr("us"),
	// 																	to.Ptr("si")},
	// 																},
	// 															},
	// 															{
	// 																APIInputParameters: []*string{
	// 																	to.Ptr("stationId"),
	// 																	to.Ptr("lat"),
	// 																	to.Ptr("lon"),
	// 																	to.Ptr("days"),
	// 																	to.Ptr("units"),
	// 																	to.Ptr("precision"),
	// 																	to.Ptr("sector")},
	// 																	APIName: to.Ptr("GetDailyForecasts"),
	// 																	CustomParameters: []*string{
	// 																		to.Ptr("stationId"),
	// 																		to.Ptr("stationLatitude"),
	// 																		to.Ptr("stationLongitude"),
	// 																		to.Ptr("timeZone"),
	// 																		to.Ptr("sunrise"),
	// 																		to.Ptr("sunset"),
	// 																		to.Ptr("weatherCode"),
	// 																		to.Ptr("weatherDescription"),
	// 																		to.Ptr("maxTemperature"),
	// 																		to.Ptr("minTemperature"),
	// 																		to.Ptr("avgHeatIndex"),
	// 																		to.Ptr("maxHeatIndex"),
	// 																		to.Ptr("minHeatIndex"),
	// 																		to.Ptr("maxWindChill"),
	// 																		to.Ptr("minWindChill"),
	// 																		to.Ptr("maxFeelsLike"),
	// 																		to.Ptr("minFeelsLike"),
	// 																		to.Ptr("avgFeelsLike"),
	// 																		to.Ptr("maxWindSpeed"),
	// 																		to.Ptr("avgWetBulbGlobeTemp"),
	// 																		to.Ptr("maxWetBulbGlobeTemp"),
	// 																		to.Ptr("minWetBulbGlobeTemp"),
	// 																		to.Ptr("minutesOfSunshine"),
	// 																		to.Ptr("cornHeatUnit"),
	// 																		to.Ptr("evapotranspiration")},
	// 																		PlatformParameters: []*string{
	// 																			to.Ptr("cloudCover"),
	// 																			to.Ptr("dewPoint"),
	// 																			to.Ptr("growingDegreeDay"),
	// 																			to.Ptr("precipitation"),
	// 																			to.Ptr("pressure"),
	// 																			to.Ptr("relativeHumidity"),
	// 																			to.Ptr("temperature"),
	// 																			to.Ptr("wetBulbTemperature"),
	// 																			to.Ptr("dateTime"),
	// 																			to.Ptr("windChill"),
	// 																			to.Ptr("windSpeed"),
	// 																			to.Ptr("windDirection")},
	// 																			UnitsSupported: &armagrifood.UnitSystemsInfo{
	// 																				Key: to.Ptr("units"),
	// 																				Values: []*string{
	// 																					to.Ptr("us"),
	// 																					to.Ptr("si")},
	// 																				},
	// 																		}},
	// 																		ExtensionAPIDocsLink: to.Ptr("https://cs-docs.dtn.com/api/weather-observations-and-forecasts-rest-api/"),
	// 																		ExtensionAuthLink: to.Ptr("https://www.dtn.com/dtn-content-integration/"),
	// 																		ExtensionCategory: to.Ptr("Weather"),
	// 																		FarmBeatsExtensionID: to.Ptr("DTN.ContentServices"),
	// 																		FarmBeatsExtensionName: to.Ptr("DTN"),
	// 																		FarmBeatsExtensionVersion: to.Ptr("1.0"),
	// 																		PublisherID: to.Ptr("dtn"),
	// 																		TargetResourceType: to.Ptr("FarmBeats"),
	// 																	},
	// 																}
}
