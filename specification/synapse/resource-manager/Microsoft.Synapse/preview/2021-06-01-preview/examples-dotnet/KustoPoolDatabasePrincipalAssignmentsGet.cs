using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Synapse.Models;
using Azure.ResourceManager.Synapse;

// Generated from example definition: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolDatabasePrincipalAssignmentsGet.json
// this example is just showing the usage of "KustoPoolDatabasePrincipalAssignments_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SynapseDatabaseResource created on azure
// for more information of creating SynapseDatabaseResource, please refer to the document of SynapseDatabaseResource
string subscriptionId = "12345678-1234-1234-1234-123456789098";
string resourceGroupName = "kustorptest";
string workspaceName = "synapseWorkspaceName";
string kustoPoolName = "kustoclusterrptest4";
string databaseName = "Kustodatabase8";
ResourceIdentifier synapseDatabaseResourceId = SynapseDatabaseResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName, kustoPoolName, databaseName);
SynapseDatabaseResource synapseDatabase = client.GetSynapseDatabaseResource(synapseDatabaseResourceId);

// get the collection of this SynapseDatabasePrincipalAssignmentResource
SynapseDatabasePrincipalAssignmentCollection collection = synapseDatabase.GetSynapseDatabasePrincipalAssignments();

// invoke the operation
string principalAssignmentName = "kustoprincipal1";
NullableResponse<SynapseDatabasePrincipalAssignmentResource> response = await collection.GetIfExistsAsync(principalAssignmentName);
SynapseDatabasePrincipalAssignmentResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine("Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    SynapseDatabasePrincipalAssignmentData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
