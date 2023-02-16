using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Blueprint;

// Generated from example definition: specification/blueprint/resource-manager/Microsoft.Blueprint/preview/2018-11-01-preview/examples/subscriptionBPDef/SealedARMTemplateArtifact_Get.json
// this example is just showing the usage of "PublishedArtifacts_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this PublishedBlueprintResource created on azure
// for more information of creating PublishedBlueprintResource, please refer to the document of PublishedBlueprintResource
string resourceScope = "subscriptions/00000000-0000-0000-0000-000000000000";
string blueprintName = "simpleBlueprint";
string versionId = "V2";
ResourceIdentifier publishedBlueprintResourceId = PublishedBlueprintResource.CreateResourceIdentifier(resourceScope, blueprintName, versionId);
PublishedBlueprintResource publishedBlueprint = client.GetPublishedBlueprintResource(publishedBlueprintResourceId);

// get the collection of this BlueprintVersionArtifactResource
BlueprintVersionArtifactCollection collection = publishedBlueprint.GetBlueprintVersionArtifacts();

// invoke the operation
string artifactName = "storageTemplate";
bool result = await collection.ExistsAsync(artifactName);

Console.WriteLine($"Succeeded: {result}");
