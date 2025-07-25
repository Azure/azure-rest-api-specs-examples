using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.NetworkCloud;

// Generated from example definition: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2025-02-01/examples/RackSkus_Get.json
// this example is just showing the usage of "RackSkus_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetworkCloudRackSkuResource created on azure
// for more information of creating NetworkCloudRackSkuResource, please refer to the document of NetworkCloudRackSkuResource
string subscriptionId = "123e4567-e89b-12d3-a456-426655440000";
string rackSkuName = "rackSkuName";
ResourceIdentifier networkCloudRackSkuResourceId = NetworkCloudRackSkuResource.CreateResourceIdentifier(subscriptionId, rackSkuName);
NetworkCloudRackSkuResource networkCloudRackSku = client.GetNetworkCloudRackSkuResource(networkCloudRackSkuResourceId);

// invoke the operation
NetworkCloudRackSkuResource result = await networkCloudRackSku.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
NetworkCloudRackSkuData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
