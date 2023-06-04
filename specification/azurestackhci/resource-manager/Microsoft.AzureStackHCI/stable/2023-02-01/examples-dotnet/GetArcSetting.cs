using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Hci;

// Generated from example definition: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/stable/2023-02-01/examples/GetArcSetting.json
// this example is just showing the usage of "ArcSettings_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this HciClusterResource created on azure
// for more information of creating HciClusterResource, please refer to the document of HciClusterResource
string subscriptionId = "fd3c3665-1729-4b7b-9a38-238e83b0f98b";
string resourceGroupName = "test-rg";
string clusterName = "myCluster";
ResourceIdentifier hciClusterResourceId = HciClusterResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, clusterName);
HciClusterResource hciCluster = client.GetHciClusterResource(hciClusterResourceId);

// get the collection of this ArcSettingResource
ArcSettingCollection collection = hciCluster.GetArcSettings();

// invoke the operation
string arcSettingName = "default";
bool result = await collection.ExistsAsync(arcSettingName);

Console.WriteLine($"Succeeded: {result}");
