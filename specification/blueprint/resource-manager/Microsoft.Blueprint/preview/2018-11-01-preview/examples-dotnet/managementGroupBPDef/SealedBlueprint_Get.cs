using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Blueprint;

// Generated from example definition: specification/blueprint/resource-manager/Microsoft.Blueprint/preview/2018-11-01-preview/examples/managementGroupBPDef/SealedBlueprint_Get.json
// this example is just showing the usage of "PublishedBlueprints_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this BlueprintResource created on azure
// for more information of creating BlueprintResource, please refer to the document of BlueprintResource
string resourceScope = "providers/Microsoft.Management/managementGroups/ContosoOnlineGroup";
string blueprintName = "simpleBlueprint";
ResourceIdentifier blueprintResourceId = BlueprintResource.CreateResourceIdentifier(resourceScope, blueprintName);
BlueprintResource blueprint = client.GetBlueprintResource(blueprintResourceId);

// get the collection of this PublishedBlueprintResource
PublishedBlueprintCollection collection = blueprint.GetPublishedBlueprints();

// invoke the operation
string versionId = "v2";
bool result = await collection.ExistsAsync(versionId);

Console.WriteLine($"Succeeded: {result}");
