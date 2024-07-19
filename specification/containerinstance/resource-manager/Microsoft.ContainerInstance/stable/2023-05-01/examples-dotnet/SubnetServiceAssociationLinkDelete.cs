using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.ContainerInstance;

// Generated from example definition: specification/containerinstance/resource-manager/Microsoft.ContainerInstance/stable/2023-05-01/examples/SubnetServiceAssociationLinkDelete.json
// this example is just showing the usage of "SubnetServiceAssociationLink_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "subid";
string resourceGroupName = "demo";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// invoke the operation
string virtualNetworkName = "demo2";
string subnetName = "demo3";
await resourceGroupResource.DeleteSubnetServiceAssociationLinkAsync(WaitUntil.Completed, virtualNetworkName, subnetName);

Console.WriteLine($"Succeeded");
