package armiotcentral_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iotcentral/armiotcentral"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/iotcentral/resource-manager/Microsoft.IoTCentral/stable/2021-06-01/examples/Apps_Templates.json
func ExampleAppsClient_NewListTemplatesPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armiotcentral.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAppsClient().NewListTemplatesPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.AppTemplatesResult = armiotcentral.AppTemplatesResult{
		// 	Value: []*armiotcentral.AppTemplate{
		// 		{
		// 			Name: to.Ptr("Store Analytics – Condition Monitoring"),
		// 			Description: to.Ptr("Digitally connect and monitor your store environment to reduce operating costs and create experiences that customers love."),
		// 			Industry: to.Ptr("Retail"),
		// 			Locations: []*armiotcentral.AppTemplateLocations{
		// 				{
		// 					DisplayName: to.Ptr("United States"),
		// 					ID: to.Ptr("unitedstates"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Europe"),
		// 					ID: to.Ptr("europe"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Asia Pacific"),
		// 					ID: to.Ptr("asiapacific"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Australia"),
		// 					ID: to.Ptr("australia"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("United Kingdom"),
		// 					ID: to.Ptr("uk"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Japan"),
		// 					ID: to.Ptr("japan"),
		// 			}},
		// 			ManifestID: to.Ptr("iotc-condition"),
		// 			ManifestVersion: to.Ptr("1.0.0"),
		// 			Order: to.Ptr[float32](99),
		// 			Title: to.Ptr("In-store Analytics – Condition Monitoring"),
		// 		},
		// 		{
		// 			Name: to.Ptr("IoT Central Water Consumption application template"),
		// 			Description: to.Ptr("Enable remote tracking of water consumption to reduce field operations, detect leaks in time, while empowering cities to conserve water."),
		// 			Industry: to.Ptr("Government"),
		// 			Locations: []*armiotcentral.AppTemplateLocations{
		// 				{
		// 					DisplayName: to.Ptr("United States"),
		// 					ID: to.Ptr("unitedstates"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Europe"),
		// 					ID: to.Ptr("europe"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Asia Pacific"),
		// 					ID: to.Ptr("asiapacific"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Australia"),
		// 					ID: to.Ptr("australia"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("United Kingdom"),
		// 					ID: to.Ptr("uk"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Japan"),
		// 					ID: to.Ptr("japan"),
		// 			}},
		// 			ManifestID: to.Ptr("iotc-consumption"),
		// 			ManifestVersion: to.Ptr("1.0.0"),
		// 			Order: to.Ptr[float32](99),
		// 			Title: to.Ptr("Water Consumption Monitoring"),
		// 		},
		// 		{
		// 			Name: to.Ptr("IoT Central Digital Distribution Center application template"),
		// 			Description: to.Ptr("Digitally manage warehouse conveyor belt system efficiency using object detection and tracking."),
		// 			Industry: to.Ptr("Retail"),
		// 			Locations: []*armiotcentral.AppTemplateLocations{
		// 				{
		// 					DisplayName: to.Ptr("United States"),
		// 					ID: to.Ptr("unitedstates"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Europe"),
		// 					ID: to.Ptr("europe"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Asia Pacific"),
		// 					ID: to.Ptr("asiapacific"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Australia"),
		// 					ID: to.Ptr("australia"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("United Kingdom"),
		// 					ID: to.Ptr("uk"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Japan"),
		// 					ID: to.Ptr("japan"),
		// 			}},
		// 			ManifestID: to.Ptr("iotc-distribution"),
		// 			ManifestVersion: to.Ptr("1.0.0"),
		// 			Order: to.Ptr[float32](99),
		// 			Title: to.Ptr("Digital Distribution Center"),
		// 		},
		// 		{
		// 			Name: to.Ptr("IoT Central Smart Inventory Management application template"),
		// 			Description: to.Ptr("Enable accurate inventory tracking and ensure shelves are always stocked."),
		// 			Industry: to.Ptr("Retail"),
		// 			Locations: []*armiotcentral.AppTemplateLocations{
		// 				{
		// 					DisplayName: to.Ptr("United States"),
		// 					ID: to.Ptr("unitedstates"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Europe"),
		// 					ID: to.Ptr("europe"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Asia Pacific"),
		// 					ID: to.Ptr("asiapacific"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Australia"),
		// 					ID: to.Ptr("australia"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("United Kingdom"),
		// 					ID: to.Ptr("uk"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Japan"),
		// 					ID: to.Ptr("japan"),
		// 			}},
		// 			ManifestID: to.Ptr("iotc-inventory"),
		// 			ManifestVersion: to.Ptr("1.0.0"),
		// 			Order: to.Ptr[float32](99),
		// 			Title: to.Ptr("Smart Inventory Management"),
		// 		},
		// 		{
		// 			Name: to.Ptr("IoT Central Connected Logistics application template"),
		// 			Description: to.Ptr("Track your shipment in real-time across air, water and land with location and condition monitoring."),
		// 			Industry: to.Ptr("Retail"),
		// 			Locations: []*armiotcentral.AppTemplateLocations{
		// 				{
		// 					DisplayName: to.Ptr("United States"),
		// 					ID: to.Ptr("unitedstates"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Europe"),
		// 					ID: to.Ptr("europe"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Asia Pacific"),
		// 					ID: to.Ptr("asiapacific"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Australia"),
		// 					ID: to.Ptr("australia"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("United Kingdom"),
		// 					ID: to.Ptr("uk"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Japan"),
		// 					ID: to.Ptr("japan"),
		// 			}},
		// 			ManifestID: to.Ptr("iotc-logistics"),
		// 			ManifestVersion: to.Ptr("1.0.0"),
		// 			Order: to.Ptr[float32](99),
		// 			Title: to.Ptr("Connected Logistics"),
		// 		},
		// 		{
		// 			Name: to.Ptr("IoT Central Smart Meter Analytics application template"),
		// 			Description: to.Ptr("Connect utility meters to gain insights into billing, forecast consumption, and proactively detect outages."),
		// 			Industry: to.Ptr("Energy"),
		// 			Locations: []*armiotcentral.AppTemplateLocations{
		// 				{
		// 					DisplayName: to.Ptr("United States"),
		// 					ID: to.Ptr("unitedstates"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Europe"),
		// 					ID: to.Ptr("europe"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Asia Pacific"),
		// 					ID: to.Ptr("asiapacific"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Australia"),
		// 					ID: to.Ptr("australia"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("United Kingdom"),
		// 					ID: to.Ptr("uk"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Japan"),
		// 					ID: to.Ptr("japan"),
		// 			}},
		// 			ManifestID: to.Ptr("iotc-meter"),
		// 			ManifestVersion: to.Ptr("1.0.0"),
		// 			Order: to.Ptr[float32](99),
		// 			Title: to.Ptr("Smart Meter Analytics"),
		// 		},
		// 		{
		// 			Name: to.Ptr("IoT Central Micro-fulfillment Center"),
		// 			Description: to.Ptr("Digitally connect, monitor and manage all aspects of a fully automated fulfillment center to reduce costs by eliminating downtime while increasing security and overall efficiency."),
		// 			Industry: to.Ptr("Retail"),
		// 			Locations: []*armiotcentral.AppTemplateLocations{
		// 				{
		// 					DisplayName: to.Ptr("United States"),
		// 					ID: to.Ptr("unitedstates"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Europe"),
		// 					ID: to.Ptr("europe"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Asia Pacific"),
		// 					ID: to.Ptr("asiapacific"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Australia"),
		// 					ID: to.Ptr("australia"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("United Kingdom"),
		// 					ID: to.Ptr("uk"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Japan"),
		// 					ID: to.Ptr("japan"),
		// 			}},
		// 			ManifestID: to.Ptr("iotc-mfc"),
		// 			ManifestVersion: to.Ptr("1.0.0"),
		// 			Order: to.Ptr[float32](99),
		// 			Title: to.Ptr("Micro-fulfillment Center"),
		// 		},
		// 		{
		// 			Name: to.Ptr("IoT Central Phone-as-a-device application template"),
		// 			Description: to.Ptr("Create application with [\"paad\"] capabilities."),
		// 			Industry: to.Ptr("Utility"),
		// 			Locations: []*armiotcentral.AppTemplateLocations{
		// 				{
		// 					DisplayName: to.Ptr("United States"),
		// 					ID: to.Ptr("unitedstates"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Europe"),
		// 					ID: to.Ptr("europe"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Asia Pacific"),
		// 					ID: to.Ptr("asiapacific"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Australia"),
		// 					ID: to.Ptr("australia"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("United Kingdom"),
		// 					ID: to.Ptr("uk"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Japan"),
		// 					ID: to.Ptr("japan"),
		// 			}},
		// 			ManifestID: to.Ptr("iotc-paad"),
		// 			ManifestVersion: to.Ptr("1.0.0"),
		// 			Order: to.Ptr[float32](99),
		// 			Title: to.Ptr("Paad"),
		// 		},
		// 		{
		// 			Name: to.Ptr("IoT Central Continuous Patient Monitoring application template"),
		// 			Description: to.Ptr("Connect and manage devices for in-patient and remote monitoring to improve patient outcomes, reduce re-admissions, and manage chronic diseases."),
		// 			Industry: to.Ptr("Health"),
		// 			Locations: []*armiotcentral.AppTemplateLocations{
		// 				{
		// 					DisplayName: to.Ptr("United States"),
		// 					ID: to.Ptr("unitedstates"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Europe"),
		// 					ID: to.Ptr("europe"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Asia Pacific"),
		// 					ID: to.Ptr("asiapacific"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Australia"),
		// 					ID: to.Ptr("australia"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("United Kingdom"),
		// 					ID: to.Ptr("uk"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Japan"),
		// 					ID: to.Ptr("japan"),
		// 			}},
		// 			ManifestID: to.Ptr("iotc-patient"),
		// 			ManifestVersion: to.Ptr("1.0.0"),
		// 			Order: to.Ptr[float32](99),
		// 			Title: to.Ptr("Continuous Patient Monitoring"),
		// 		},
		// 		{
		// 			Name: to.Ptr("IoT Central PnP template (preview)"),
		// 			Description: to.Ptr("Create an application with Azure IoT Plug and Play."),
		// 			Industry: to.Ptr(""),
		// 			Locations: []*armiotcentral.AppTemplateLocations{
		// 				{
		// 					DisplayName: to.Ptr("United States"),
		// 					ID: to.Ptr("unitedstates"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Europe"),
		// 					ID: to.Ptr("europe"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Asia Pacific"),
		// 					ID: to.Ptr("asiapacific"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Australia"),
		// 					ID: to.Ptr("australia"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("United Kingdom"),
		// 					ID: to.Ptr("uk"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Japan"),
		// 					ID: to.Ptr("japan"),
		// 			}},
		// 			ManifestID: to.Ptr("iotc-pnp-preview"),
		// 			ManifestVersion: to.Ptr("1.0.0"),
		// 			Order: to.Ptr[float32](1),
		// 			Title: to.Ptr("Custom application"),
		// 		},
		// 		{
		// 			Name: to.Ptr("IoT Central Solar Power Monitoring application template"),
		// 			Description: to.Ptr("Connect, monitor, and manage your solar panels and energy generation."),
		// 			Industry: to.Ptr("Energy"),
		// 			Locations: []*armiotcentral.AppTemplateLocations{
		// 				{
		// 					DisplayName: to.Ptr("United States"),
		// 					ID: to.Ptr("unitedstates"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Europe"),
		// 					ID: to.Ptr("europe"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Asia Pacific"),
		// 					ID: to.Ptr("asiapacific"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Australia"),
		// 					ID: to.Ptr("australia"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("United Kingdom"),
		// 					ID: to.Ptr("uk"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Japan"),
		// 					ID: to.Ptr("japan"),
		// 			}},
		// 			ManifestID: to.Ptr("iotc-power"),
		// 			ManifestVersion: to.Ptr("1.0.0"),
		// 			Order: to.Ptr[float32](99),
		// 			Title: to.Ptr("Solar Power Monitoring"),
		// 		},
		// 		{
		// 			Name: to.Ptr("IoT Central Water Quality Monitoring application template"),
		// 			Description: to.Ptr("Improve water quality and detect issues earlier by analyzing real-time measurements across your environment."),
		// 			Industry: to.Ptr("Government"),
		// 			Locations: []*armiotcentral.AppTemplateLocations{
		// 				{
		// 					DisplayName: to.Ptr("United States"),
		// 					ID: to.Ptr("unitedstates"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Europe"),
		// 					ID: to.Ptr("europe"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Asia Pacific"),
		// 					ID: to.Ptr("asiapacific"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Australia"),
		// 					ID: to.Ptr("australia"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("United Kingdom"),
		// 					ID: to.Ptr("uk"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Japan"),
		// 					ID: to.Ptr("japan"),
		// 			}},
		// 			ManifestID: to.Ptr("iotc-quality"),
		// 			ManifestVersion: to.Ptr("1.0.0"),
		// 			Order: to.Ptr[float32](99),
		// 			Title: to.Ptr("Water Quality Monitoring"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Store Analytics – Checkout"),
		// 			Description: to.Ptr("Monitor and manage the checkout flow inside your store to improve efficiency and reduce wait times."),
		// 			Industry: to.Ptr("Retail"),
		// 			Locations: []*armiotcentral.AppTemplateLocations{
		// 				{
		// 					DisplayName: to.Ptr("United States"),
		// 					ID: to.Ptr("unitedstates"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Europe"),
		// 					ID: to.Ptr("europe"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Asia Pacific"),
		// 					ID: to.Ptr("asiapacific"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Australia"),
		// 					ID: to.Ptr("australia"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("United Kingdom"),
		// 					ID: to.Ptr("uk"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Japan"),
		// 					ID: to.Ptr("japan"),
		// 			}},
		// 			ManifestID: to.Ptr("iotc-store"),
		// 			ManifestVersion: to.Ptr("1.0.0"),
		// 			Order: to.Ptr[float32](99),
		// 			Title: to.Ptr("In-store Analytics – Checkout"),
		// 		},
		// 		{
		// 			Name: to.Ptr("IoT Central Video analytics - object and motion detection application template"),
		// 			Description: to.Ptr("Use cameras as a sensor in intelligent video analytics solutions powered by Azure IoT Edge, AI, and Azure Media Services."),
		// 			Industry: to.Ptr("Retail"),
		// 			Locations: []*armiotcentral.AppTemplateLocations{
		// 				{
		// 					DisplayName: to.Ptr("United States"),
		// 					ID: to.Ptr("unitedstates"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Europe"),
		// 					ID: to.Ptr("europe"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Asia Pacific"),
		// 					ID: to.Ptr("asiapacific"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Australia"),
		// 					ID: to.Ptr("australia"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("United Kingdom"),
		// 					ID: to.Ptr("uk"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Japan"),
		// 					ID: to.Ptr("japan"),
		// 			}},
		// 			ManifestID: to.Ptr("iotc-video-analytics-om"),
		// 			ManifestVersion: to.Ptr("1.0.0"),
		// 			Order: to.Ptr[float32](99),
		// 			Title: to.Ptr("Video analytics - object and motion detection"),
		// 		},
		// 		{
		// 			Name: to.Ptr("IoT Central Connected Waste Management application template"),
		// 			Description: to.Ptr("Maximize efficiency in the collection of solid wastes by dispatching field operators at the right time along an optimized collection route."),
		// 			Industry: to.Ptr("Government"),
		// 			Locations: []*armiotcentral.AppTemplateLocations{
		// 				{
		// 					DisplayName: to.Ptr("United States"),
		// 					ID: to.Ptr("unitedstates"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Europe"),
		// 					ID: to.Ptr("europe"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Asia Pacific"),
		// 					ID: to.Ptr("asiapacific"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Australia"),
		// 					ID: to.Ptr("australia"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("United Kingdom"),
		// 					ID: to.Ptr("uk"),
		// 				},
		// 				{
		// 					DisplayName: to.Ptr("Japan"),
		// 					ID: to.Ptr("japan"),
		// 			}},
		// 			ManifestID: to.Ptr("iotc-waste"),
		// 			ManifestVersion: to.Ptr("1.0.0"),
		// 			Order: to.Ptr[float32](99),
		// 			Title: to.Ptr("Connected Waste Management"),
		// 	}},
		// }
	}
}
