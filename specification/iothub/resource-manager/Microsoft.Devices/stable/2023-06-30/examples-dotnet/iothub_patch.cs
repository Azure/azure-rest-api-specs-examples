using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.IotHub.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.IotHub;

// Generated from example definition: specification/iothub/resource-manager/Microsoft.Devices/stable/2023-06-30/examples/iothub_patch.json
// this example is just showing the usage of "IotHubResource_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this IotHubDescriptionResource created on azure
// for more information of creating IotHubDescriptionResource, please refer to the document of IotHubDescriptionResource
string subscriptionId = "91d12660-3dec-467a-be2a-213b5544ddc0";
string resourceGroupName = "myResourceGroup";
string resourceName = "myHub";
ResourceIdentifier iotHubDescriptionResourceId = IotHubDescriptionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, resourceName);
IotHubDescriptionResource iotHubDescription = client.GetIotHubDescriptionResource(iotHubDescriptionResourceId);

// invoke the operation
IotHubDescriptionPatch patch = new IotHubDescriptionPatch()
{
    Tags =
    {
    ["foo"] = "bar",
    },
};
ArmOperation<IotHubDescriptionResource> lro = await iotHubDescription.UpdateAsync(WaitUntil.Completed, patch);
IotHubDescriptionResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
IotHubDescriptionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
