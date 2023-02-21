using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Synapse;

// Generated from example definition: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/GetWorkspaceAadAdmin.json
// this example is just showing the usage of "WorkspaceAadAdmins_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SynapseWorkspaceAdministratorResource created on azure
// for more information of creating SynapseWorkspaceAdministratorResource, please refer to the document of SynapseWorkspaceAdministratorResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "resourceGroup1";
string workspaceName = "workspace1";
ResourceIdentifier synapseWorkspaceAdministratorResourceId = SynapseWorkspaceAdministratorResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName);
SynapseWorkspaceAdministratorResource synapseWorkspaceAdministratorResource = client.GetSynapseWorkspaceAdministratorResource(synapseWorkspaceAdministratorResourceId);

// invoke the operation
SynapseWorkspaceAdministratorResource result = await synapseWorkspaceAdministratorResource.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SynapseWorkspaceAadAdminInfoData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
