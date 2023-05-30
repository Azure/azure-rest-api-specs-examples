using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Blueprint;
using Azure.ResourceManager.Blueprint.Models;

// Generated from example definition: specification/blueprint/resource-manager/Microsoft.Blueprint/preview/2018-11-01-preview/examples/managementGroupBPDef/Blueprint_Get.json
// this example is just showing the usage of "Blueprints_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ArmResource created on azure
// for more information of creating ArmResource, please refer to the document of ArmResource

// get the collection of this BlueprintResource
string resourceScope = "providers/Microsoft.Management/managementGroups/ContosoOnlineGroup";
ResourceIdentifier scopeId = new ResourceIdentifier(string.Format("/{0}", resourceScope));
BlueprintCollection collection = client.GetBlueprints(scopeId);

// invoke the operation
string blueprintName = "simpleBlueprint";
bool result = await collection.ExistsAsync(blueprintName);

Console.WriteLine($"Succeeded: {result}");
