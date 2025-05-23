using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Synapse;

// Generated from example definition: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/DeleteWorkspaceAadAdmin.json
// this example is just showing the usage of "WorkspaceSqlAadAdmins_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SynapseWorkspaceSqlAdministratorResource created on azure
// for more information of creating SynapseWorkspaceSqlAdministratorResource, please refer to the document of SynapseWorkspaceSqlAdministratorResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "resourceGroup1";
string workspaceName = "workspace1";
ResourceIdentifier synapseWorkspaceSqlAdministratorResourceId = SynapseWorkspaceSqlAdministratorResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName);
SynapseWorkspaceSqlAdministratorResource synapseWorkspaceSqlAdministratorResource = client.GetSynapseWorkspaceSqlAdministratorResource(synapseWorkspaceSqlAdministratorResourceId);

// invoke the operation
await synapseWorkspaceSqlAdministratorResource.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
