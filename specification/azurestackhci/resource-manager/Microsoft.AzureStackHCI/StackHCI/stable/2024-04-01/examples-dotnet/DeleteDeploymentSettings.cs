using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Hci.Models;
using Azure.ResourceManager.Hci;

// Generated from example definition: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/StackHCI/stable/2024-04-01/examples/DeleteDeploymentSettings.json
// this example is just showing the usage of "DeploymentSettings_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this HciClusterDeploymentSettingResource created on azure
// for more information of creating HciClusterDeploymentSettingResource, please refer to the document of HciClusterDeploymentSettingResource
string subscriptionId = "fd3c3665-1729-4b7b-9a38-238e83b0f98b";
string resourceGroupName = "test-rg";
string clusterName = "myCluster";
string deploymentSettingsName = "default";
ResourceIdentifier hciClusterDeploymentSettingResourceId = HciClusterDeploymentSettingResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, clusterName, deploymentSettingsName);
HciClusterDeploymentSettingResource hciClusterDeploymentSetting = client.GetHciClusterDeploymentSettingResource(hciClusterDeploymentSettingResourceId);

// invoke the operation
await hciClusterDeploymentSetting.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
