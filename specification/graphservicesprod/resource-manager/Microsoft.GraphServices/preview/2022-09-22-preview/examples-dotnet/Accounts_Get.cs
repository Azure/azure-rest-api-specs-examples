using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.GraphServices;
using Azure.ResourceManager.GraphServices.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/graphservicesprod/resource-manager/Microsoft.GraphServices/preview/2022-09-22-preview/examples/Accounts_Get.json
// this example is just showing the usage of "Account_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "testResourceGroupGRAM";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this GraphServicesAccountResource
GraphServicesAccountResourceCollection collection = resourceGroupResource.GetGraphServicesAccountResources();

// invoke the operation
string resourceName = "11111111-aaaa-1111-bbbb-111111111111";
bool result = await collection.ExistsAsync(resourceName);

Console.WriteLine($"Succeeded: {result}");
