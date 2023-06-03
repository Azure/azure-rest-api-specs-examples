using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Compute;
using Azure.ResourceManager.Compute.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/CloudserviceRP/stable/2022-09-04/examples/CloudService_Get_WithMultiRoleAndRDP.json
// this example is just showing the usage of "CloudServices_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "ConstosoRG";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this CloudServiceResource
CloudServiceCollection collection = resourceGroupResource.GetCloudServices();

// invoke the operation
string cloudServiceName = "{cs-name}";
bool result = await collection.ExistsAsync(cloudServiceName);

Console.WriteLine($"Succeeded: {result}");
