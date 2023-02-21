using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Synapse;
using Azure.ResourceManager.Synapse.Models;

// Generated from example definition: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolDatabasePrincipalAssignmentsDelete.json
// this example is just showing the usage of "KustoPoolDatabasePrincipalAssignments_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SynapseDatabasePrincipalAssignmentResource created on azure
// for more information of creating SynapseDatabasePrincipalAssignmentResource, please refer to the document of SynapseDatabasePrincipalAssignmentResource
string subscriptionId = "12345678-1234-1234-1234-123456789098";
string resourceGroupName = "kustorptest";
string workspaceName = "synapseWorkspaceName";
string kustoPoolName = "kustoclusterrptest4";
string databaseName = "Kustodatabase8";
string principalAssignmentName = "kustoprincipal1";
ResourceIdentifier synapseDatabasePrincipalAssignmentResourceId = SynapseDatabasePrincipalAssignmentResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName, kustoPoolName, databaseName, principalAssignmentName);
SynapseDatabasePrincipalAssignmentResource synapseDatabasePrincipalAssignment = client.GetSynapseDatabasePrincipalAssignmentResource(synapseDatabasePrincipalAssignmentResourceId);

// invoke the operation
await synapseDatabasePrincipalAssignment.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
