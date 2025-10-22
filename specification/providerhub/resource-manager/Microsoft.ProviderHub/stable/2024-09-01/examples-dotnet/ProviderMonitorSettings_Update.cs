using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ProviderHub;

// Generated from example definition: specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2024-09-01/examples/ProviderMonitorSettings_Update.json
// this example is just showing the usage of "ProviderMonitorSettings_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ProviderMonitorSettingResource created on azure
// for more information of creating ProviderMonitorSettingResource, please refer to the document of ProviderMonitorSettingResource
string subscriptionId = "ab7a8701-f7ef-471a-a2f4-d0ebbf494f77";
string resourceGroupName = "default";
string providerMonitorSettingName = "ContosoMonitorSetting";
ResourceIdentifier providerMonitorSettingResourceId = ProviderMonitorSettingResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, providerMonitorSettingName);
ProviderMonitorSettingResource providerMonitorSetting = client.GetProviderMonitorSettingResource(providerMonitorSettingResourceId);

// invoke the operation
ProviderMonitorSettingResource result = await providerMonitorSetting.UpdateAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ProviderMonitorSettingData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
