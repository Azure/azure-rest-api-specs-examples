using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Synapse.Models;
using Azure.ResourceManager.Synapse;

// Generated from example definition: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/GetWorkspaceManagedSqlExtendedServerBlobAuditingSettings.json
// this example is just showing the usage of "WorkspaceManagedSqlServerExtendedBlobAuditingPolicies_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SynapseExtendedServerBlobAuditingPolicyResource created on azure
// for more information of creating SynapseExtendedServerBlobAuditingPolicyResource, please refer to the document of SynapseExtendedServerBlobAuditingPolicyResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "wsg-7398";
string workspaceName = "testWorkspace";
SynapseBlobAuditingPolicyName blobAuditingPolicyName = SynapseBlobAuditingPolicyName.Default;
ResourceIdentifier synapseExtendedServerBlobAuditingPolicyResourceId = SynapseExtendedServerBlobAuditingPolicyResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName, blobAuditingPolicyName);
SynapseExtendedServerBlobAuditingPolicyResource synapseExtendedServerBlobAuditingPolicy = client.GetSynapseExtendedServerBlobAuditingPolicyResource(synapseExtendedServerBlobAuditingPolicyResourceId);

// invoke the operation
SynapseExtendedServerBlobAuditingPolicyResource result = await synapseExtendedServerBlobAuditingPolicy.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SynapseExtendedServerBlobAuditingPolicyData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
