package armdeploymentmanager_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deploymentmanager/armdeploymentmanager"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/deploymentmanager/resource-manager/Microsoft.DeploymentManager/preview/2019-11-01-preview/examples/step_health_check_createorupdate.json
func ExampleStepsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdeploymentmanager.NewStepsClient("caac1590-e859-444f-a9e0-62091c0f5929", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.CreateOrUpdate(ctx,
		"myResourceGroup",
		"healthCheckStep",
		&armdeploymentmanager.StepsClientCreateOrUpdateOptions{StepInfo: &armdeploymentmanager.StepResource{
			Location: to.Ptr("centralus"),
			Tags:     map[string]*string{},
			Properties: &armdeploymentmanager.HealthCheckStepProperties{
				StepType: to.Ptr(armdeploymentmanager.StepTypeHealthCheck),
				Attributes: &armdeploymentmanager.RestHealthCheckStepAttributes{
					Type:                 to.Ptr("REST"),
					HealthyStateDuration: to.Ptr("PT2H"),
					MaxElasticDuration:   to.Ptr("PT30M"),
					WaitDuration:         to.Ptr("PT15M"),
					Properties: &armdeploymentmanager.RestParameters{
						HealthChecks: []*armdeploymentmanager.RestHealthCheck{
							{
								Name: to.Ptr("appHealth"),
								Response: &armdeploymentmanager.RestResponse{
									Regex: &armdeploymentmanager.RestResponseRegex{
										MatchQuantifier: to.Ptr(armdeploymentmanager.RestMatchQuantifierAll),
										Matches: []*string{
											to.Ptr("(?i)Contoso-App"),
											to.Ptr("(?i)\"health_status\":((.|\n)*)\"(green|yellow)\""),
											to.Ptr("(?mi)^(\"application_host\": 94781052)$")},
									},
									SuccessStatusCodes: []*string{
										to.Ptr("OK")},
								},
								Request: &armdeploymentmanager.RestRequest{
									Method: to.Ptr(armdeploymentmanager.RestRequestMethodGET),
									Authentication: &armdeploymentmanager.APIKeyAuthentication{
										Type:  to.Ptr(armdeploymentmanager.RestAuthTypeAPIKey),
										Name:  to.Ptr("Code"),
										In:    to.Ptr(armdeploymentmanager.RestAuthLocationQuery),
										Value: to.Ptr("NBCapiMOBQyAAbCkeytoPadnvO0eGHmidwFz5rXpappznKp3Jt7LLg=="),
									},
									URI: to.Ptr("https://resthealth.healthservice.com/api/applications/contosoApp/healthStatus"),
								},
							},
							{
								Name: to.Ptr("serviceHealth"),
								Response: &armdeploymentmanager.RestResponse{
									Regex: &armdeploymentmanager.RestResponseRegex{
										MatchQuantifier: to.Ptr(armdeploymentmanager.RestMatchQuantifierAll),
										Matches: []*string{
											to.Ptr("(?i)Contoso-Service-EndToEnd"),
											to.Ptr("(?i)\"health_status\":((.|\n)*)\"(green)\"")},
									},
									SuccessStatusCodes: []*string{
										to.Ptr("OK")},
								},
								Request: &armdeploymentmanager.RestRequest{
									Method: to.Ptr(armdeploymentmanager.RestRequestMethodGET),
									Authentication: &armdeploymentmanager.APIKeyAuthentication{
										Type:  to.Ptr(armdeploymentmanager.RestAuthTypeAPIKey),
										Name:  to.Ptr("code"),
										In:    to.Ptr(armdeploymentmanager.RestAuthLocationHeader),
										Value: to.Ptr("NBCapiMOBQyAAbCkeytoPadnvO0eGHmidwFz5rXpappznKp3Jt7LLg=="),
									},
									URI: to.Ptr("https://resthealth.healthservice.com/api/services/contosoService/healthStatus"),
								},
							}},
					},
				},
			},
		},
		})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
