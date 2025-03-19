using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.IotCentral.Models;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.IotCentral;

// Generated from example definition: specification/iotcentral/resource-manager/Microsoft.IoTCentral/preview/2021-11-01-preview/examples/Apps_CreateOrUpdate.json
// this example is just showing the usage of "Apps_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "resRg";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this IotCentralAppResource
IotCentralAppCollection collection = resourceGroupResource.GetIotCentralApps();

// invoke the operation
string resourceName = "myIoTCentralApp";
IotCentralAppData data = new IotCentralAppData(new AzureLocation("westus"), new IotCentralAppSkuInfo(IotCentralAppSku.ST2))
{
    Identity = new ManagedServiceIdentity("SystemAssigned"),
    DisplayName = "My IoT Central App",
    Subdomain = "my-iot-central-app",
    Template = "iotc-pnp-preview@1.0.0",
};
ArmOperation<IotCentralAppResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, resourceName, data);
IotCentralAppResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
IotCentralAppData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
