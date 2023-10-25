using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Kusto;
using Azure.ResourceManager.Kusto.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2023-08-15/examples/KustoClusterDetachFollowerDatabases.json
// this example is just showing the usage of "Clusters_DetachFollowerDatabases" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this KustoClusterResource created on azure
// for more information of creating KustoClusterResource, please refer to the document of KustoClusterResource
string subscriptionId = "12345678-1234-1234-1234-123456789098";
string resourceGroupName = "kustorptest";
string clusterName = "kustoCluster";
ResourceIdentifier kustoClusterResourceId = KustoClusterResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, clusterName);
KustoClusterResource kustoCluster = client.GetKustoClusterResource(kustoClusterResourceId);

// invoke the operation
KustoFollowerDatabaseDefinition followerDatabaseToRemove = new KustoFollowerDatabaseDefinition(new ResourceIdentifier("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Kusto/clusters/kustoCluster2"), "attachedDatabaseConfigurationsTest");
await kustoCluster.DetachFollowerDatabasesAsync(WaitUntil.Completed, followerDatabaseToRemove);

Console.WriteLine($"Succeeded");
