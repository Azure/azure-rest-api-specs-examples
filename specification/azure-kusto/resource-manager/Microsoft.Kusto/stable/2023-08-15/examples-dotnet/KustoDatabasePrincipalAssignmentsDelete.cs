using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Kusto;
using Azure.ResourceManager.Kusto.Models;

// Generated from example definition: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2023-08-15/examples/KustoDatabasePrincipalAssignmentsDelete.json
// this example is just showing the usage of "DatabasePrincipalAssignments_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this KustoDatabasePrincipalAssignmentResource created on azure
// for more information of creating KustoDatabasePrincipalAssignmentResource, please refer to the document of KustoDatabasePrincipalAssignmentResource
string subscriptionId = "12345678-1234-1234-1234-123456789098";
string resourceGroupName = "kustorptest";
string clusterName = "kustoCluster";
string databaseName = "Kustodatabase8";
string principalAssignmentName = "kustoprincipal1";
ResourceIdentifier kustoDatabasePrincipalAssignmentResourceId = KustoDatabasePrincipalAssignmentResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, clusterName, databaseName, principalAssignmentName);
KustoDatabasePrincipalAssignmentResource kustoDatabasePrincipalAssignment = client.GetKustoDatabasePrincipalAssignmentResource(kustoDatabasePrincipalAssignmentResourceId);

// invoke the operation
await kustoDatabasePrincipalAssignment.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
