using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.SecurityCenter;

// Generated from example definition: specification/security/resource-manager/Microsoft.Security/preview/2017-08-01-preview/examples/WorkspaceSettings/UpdateWorkspaceSetting_example.json
// this example is just showing the usage of "WorkspaceSettings_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SecurityWorkspaceSettingResource created on azure
// for more information of creating SecurityWorkspaceSettingResource, please refer to the document of SecurityWorkspaceSettingResource
string subscriptionId = "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
string workspaceSettingName = "default";
ResourceIdentifier securityWorkspaceSettingResourceId = SecurityWorkspaceSettingResource.CreateResourceIdentifier(subscriptionId, workspaceSettingName);
SecurityWorkspaceSettingResource securityWorkspaceSetting = client.GetSecurityWorkspaceSettingResource(securityWorkspaceSettingResourceId);

// invoke the operation
SecurityWorkspaceSettingData data = new SecurityWorkspaceSettingData()
{
    WorkspaceId = new ResourceIdentifier("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace"),
};
SecurityWorkspaceSettingResource result = await securityWorkspaceSetting.UpdateAsync(data);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SecurityWorkspaceSettingData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
