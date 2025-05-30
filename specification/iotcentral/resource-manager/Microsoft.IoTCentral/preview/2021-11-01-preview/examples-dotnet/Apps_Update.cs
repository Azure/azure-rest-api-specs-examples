using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.IotCentral.Models;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.IotCentral;

// Generated from example definition: specification/iotcentral/resource-manager/Microsoft.IoTCentral/preview/2021-11-01-preview/examples/Apps_Update.json
// this example is just showing the usage of "Apps_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this IotCentralAppResource created on azure
// for more information of creating IotCentralAppResource, please refer to the document of IotCentralAppResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "resRg";
string resourceName = "myIoTCentralApp";
ResourceIdentifier iotCentralAppResourceId = IotCentralAppResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, resourceName);
IotCentralAppResource iotCentralApp = client.GetIotCentralAppResource(iotCentralAppResourceId);

// invoke the operation
IotCentralAppPatch patch = new IotCentralAppPatch
{
    Identity = new ManagedServiceIdentity("SystemAssigned"),
    DisplayName = "My IoT Central App 2",
};
await iotCentralApp.UpdateAsync(WaitUntil.Completed, patch);

Console.WriteLine("Succeeded");
