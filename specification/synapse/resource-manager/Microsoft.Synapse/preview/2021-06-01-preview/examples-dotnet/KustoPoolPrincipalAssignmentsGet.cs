using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Synapse;
using Azure.ResourceManager.Synapse.Models;

// Generated from example definition: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolPrincipalAssignmentsGet.json
// this example is just showing the usage of "KustoPoolPrincipalAssignments_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SynapseClusterPrincipalAssignmentResource created on azure
// for more information of creating SynapseClusterPrincipalAssignmentResource, please refer to the document of SynapseClusterPrincipalAssignmentResource
string subscriptionId = "12345678-1234-1234-1234-123456789098";
string resourceGroupName = "kustorptest";
string workspaceName = "synapseWorkspaceName";
string kustoPoolName = "kustoclusterrptest4";
string principalAssignmentName = "kustoprincipal1";
ResourceIdentifier synapseClusterPrincipalAssignmentResourceId = SynapseClusterPrincipalAssignmentResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName, kustoPoolName, principalAssignmentName);
SynapseClusterPrincipalAssignmentResource synapseClusterPrincipalAssignment = client.GetSynapseClusterPrincipalAssignmentResource(synapseClusterPrincipalAssignmentResourceId);

// invoke the operation
SynapseClusterPrincipalAssignmentResource result = await synapseClusterPrincipalAssignment.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SynapseClusterPrincipalAssignmentData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
