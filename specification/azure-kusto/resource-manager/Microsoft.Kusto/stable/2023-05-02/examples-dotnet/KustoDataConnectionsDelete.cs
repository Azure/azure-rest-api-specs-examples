using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Kusto;
using Azure.ResourceManager.Kusto.Models;

// Generated from example definition: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2023-05-02/examples/KustoDataConnectionsDelete.json
// this example is just showing the usage of "DataConnections_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this KustoDataConnectionResource created on azure
// for more information of creating KustoDataConnectionResource, please refer to the document of KustoDataConnectionResource
string subscriptionId = "12345678-1234-1234-1234-123456789098";
string resourceGroupName = "kustorptest";
string clusterName = "kustoCluster";
string databaseName = "KustoDatabase8";
string dataConnectionName = "dataConnectionTest";
ResourceIdentifier kustoDataConnectionResourceId = KustoDataConnectionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, clusterName, databaseName, dataConnectionName);
KustoDataConnectionResource kustoDataConnection = client.GetKustoDataConnectionResource(kustoDataConnectionResourceId);

// invoke the operation
await kustoDataConnection.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
