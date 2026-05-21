package armdevhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devhub/armdevhub"
)

// Generated from example definition: 2025-03-01-preview/VersionedTemplate_Generate.json
func ExampleVersionedTemplateClient_Generate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevhub.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVersionedTemplateClient().Generate(ctx, "example-template", "1.0.0", map[string]*string{
		"appName":                   to.Ptr("my-app"),
		"dockerfileGenerationMode":  to.Ptr("enabled"),
		"dockerfileOutputDirectory": to.Ptr("./"),
		"generationLanguage":        to.Ptr("javascript"),
		"imageName":                 to.Ptr("myimage"),
		"imageTag":                  to.Ptr("latest"),
		"languageVersion":           to.Ptr("14"),
		"manifestGenerationMode":    to.Ptr("enabled"),
		"manifestOutputDirectory":   to.Ptr("./"),
		"manifestType":              to.Ptr("kube"),
		"namespace":                 to.Ptr("my-namespace"),
		"port":                      to.Ptr("80"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdevhub.VersionedTemplateClientGenerateResponse{
	// 	GenerateVersionedTemplateResponse: armdevhub.GenerateVersionedTemplateResponse{
	// 		GeneratedFiles: map[string]*string{
	// 			"dockerfiles/example-dockerfile": to.Ptr("dockerfile-content \n including newlines"),
	// 			"manifests/example-manifest-file-1": to.Ptr("manifest file 1 content \n including newlines"),
	// 			"manifests/example-manifest-file-2": to.Ptr("manifest file 2 content \n including newlines"),
	// 		},
	// 	},
	// }
}
