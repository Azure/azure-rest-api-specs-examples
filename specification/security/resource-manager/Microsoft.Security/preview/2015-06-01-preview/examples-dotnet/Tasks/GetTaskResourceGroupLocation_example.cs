using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.SecurityCenter;

// Generated from example definition: specification/security/resource-manager/Microsoft.Security/preview/2015-06-01-preview/examples/Tasks/GetTaskResourceGroupLocation_example.json
// this example is just showing the usage of "Tasks_GetResourceGroupLevelTask" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
string resourceGroupName = "myRg";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this ResourceGroupSecurityTaskResource
AzureLocation ascLocation = new AzureLocation("westeurope");
ResourceGroupSecurityTaskCollection collection = resourceGroupResource.GetResourceGroupSecurityTasks(ascLocation);

// invoke the operation
string taskName = "d55b4dc0-779c-c66c-33e5-d7bce24c4222";
bool result = await collection.ExistsAsync(taskName);

Console.WriteLine($"Succeeded: {result}");
