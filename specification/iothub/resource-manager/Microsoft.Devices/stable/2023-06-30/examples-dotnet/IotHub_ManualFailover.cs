using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.IotHub;
using Azure.ResourceManager.IotHub.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/iothub/resource-manager/Microsoft.Devices/stable/2023-06-30/examples/IotHub_ManualFailover.json
// this example is just showing the usage of "IotHub_ManualFailover" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this IotHubDescriptionResource created on azure
// for more information of creating IotHubDescriptionResource, please refer to the document of IotHubDescriptionResource
string subscriptionId = "91d12660-3dec-467a-be2a-213b5544ddc0";
string resourceGroupName = "myResourceGroup";
string iotHubName = "testHub";
ResourceIdentifier iotHubDescriptionResourceId = IotHubDescriptionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, iotHubName);
IotHubDescriptionResource iotHubDescription = client.GetIotHubDescriptionResource(iotHubDescriptionResourceId);

// invoke the operation
IotHubFailoverContent content = new IotHubFailoverContent("testHub");
await iotHubDescription.ManualFailoverIotHubAsync(WaitUntil.Completed, content);

Console.WriteLine($"Succeeded");
