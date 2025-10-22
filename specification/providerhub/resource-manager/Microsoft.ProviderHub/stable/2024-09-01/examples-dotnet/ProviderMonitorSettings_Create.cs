using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.ProviderHub;

// Generated from example definition: specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2024-09-01/examples/ProviderMonitorSettings_Create.json
// this example is just showing the usage of "ProviderMonitorSettings_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "ab7a8701-f7ef-471a-a2f4-d0ebbf494f77";
string resourceGroupName = "default";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this ProviderMonitorSettingResource
ProviderMonitorSettingCollection collection = resourceGroupResource.GetProviderMonitorSettings();

// invoke the operation
string providerMonitorSettingName = "ContosoMonitorSetting";
ProviderMonitorSettingData data = new ProviderMonitorSettingData(new AzureLocation("eastus"));
ArmOperation<ProviderMonitorSettingResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, providerMonitorSettingName, data);
ProviderMonitorSettingResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ProviderMonitorSettingData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
