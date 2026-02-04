using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/CloudServiceSwapPut.json
// this example is just showing the usage of "VipSwap_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this CloudServiceSwapResource created on azure
// for more information of creating CloudServiceSwapResource, please refer to the document of CloudServiceSwapResource
string subscriptionId = "subid";
string groupName = "rg1";
string resourceName = "testCloudService";
ResourceIdentifier cloudServiceSwapResourceId = CloudServiceSwapResource.CreateResourceIdentifier(subscriptionId, groupName, resourceName);
CloudServiceSwapResource cloudServiceSwap = client.GetCloudServiceSwapResource(cloudServiceSwapResourceId);

// invoke the operation
CloudServiceSwapData data = new CloudServiceSwapData
{
    CloudServiceSwapSlotType = SwapSlotType.Production,
};
await cloudServiceSwap.UpdateAsync(WaitUntil.Completed, data);

Console.WriteLine("Succeeded");
