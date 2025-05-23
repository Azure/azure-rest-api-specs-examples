using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ManagedNetworkFabric;

// Generated from example definition: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/NetworkDeviceSkus_Get_MaximumSet_Gen.json
// this example is just showing the usage of "NetworkDeviceSkus_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetworkDeviceSkuResource created on azure
// for more information of creating NetworkDeviceSkuResource, please refer to the document of NetworkDeviceSkuResource
string subscriptionId = "1234ABCD-0A1B-1234-5678-123456ABCDEF";
string networkDeviceSkuName = "example-deviceSku";
ResourceIdentifier networkDeviceSkuResourceId = NetworkDeviceSkuResource.CreateResourceIdentifier(subscriptionId, networkDeviceSkuName);
NetworkDeviceSkuResource networkDeviceSku = client.GetNetworkDeviceSkuResource(networkDeviceSkuResourceId);

// invoke the operation
NetworkDeviceSkuResource result = await networkDeviceSku.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
NetworkDeviceSkuData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
