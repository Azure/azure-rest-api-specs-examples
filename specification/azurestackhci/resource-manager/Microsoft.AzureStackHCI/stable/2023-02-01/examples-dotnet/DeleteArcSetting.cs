using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Hci;
using Azure.ResourceManager.Hci.Models;

// Generated from example definition: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/stable/2023-02-01/examples/DeleteArcSetting.json
// this example is just showing the usage of "ArcSettings_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ArcSettingResource created on azure
// for more information of creating ArcSettingResource, please refer to the document of ArcSettingResource
string subscriptionId = "fd3c3665-1729-4b7b-9a38-238e83b0f98b";
string resourceGroupName = "test-rg";
string clusterName = "myCluster";
string arcSettingName = "default";
ResourceIdentifier arcSettingResourceId = ArcSettingResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, clusterName, arcSettingName);
ArcSettingResource arcSetting = client.GetArcSettingResource(arcSettingResourceId);

// invoke the operation
await arcSetting.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
