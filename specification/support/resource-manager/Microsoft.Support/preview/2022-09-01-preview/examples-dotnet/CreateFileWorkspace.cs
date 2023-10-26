using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Support;

// Generated from example definition: specification/support/resource-manager/Microsoft.Support/preview/2022-09-01-preview/examples/CreateFileWorkspace.json
// this example is just showing the usage of "FileWorkspacesNoSubscription_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this TenantFileWorkspaceResource created on azure
// for more information of creating TenantFileWorkspaceResource, please refer to the document of TenantFileWorkspaceResource
string fileWorkspaceName = "testworkspace";
ResourceIdentifier tenantFileWorkspaceResourceId = TenantFileWorkspaceResource.CreateResourceIdentifier(fileWorkspaceName);
TenantFileWorkspaceResource tenantFileWorkspace = client.GetTenantFileWorkspaceResource(tenantFileWorkspaceResourceId);

// invoke the operation
ArmOperation<TenantFileWorkspaceResource> lro = await tenantFileWorkspace.UpdateAsync(WaitUntil.Completed);
TenantFileWorkspaceResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
FileWorkspaceDetailData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
