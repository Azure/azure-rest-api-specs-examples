using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Kusto;
using Azure.ResourceManager.Kusto.Models;

// Generated from example definition: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2022-11-11/examples/KustoAttachedDatabaseConfigurationsGet.json
// this example is just showing the usage of "AttachedDatabaseConfigurations_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this KustoClusterResource created on azure
// for more information of creating KustoClusterResource, please refer to the document of KustoClusterResource
string subscriptionId = "12345678-1234-1234-1234-123456789098";
string resourceGroupName = "kustorptest";
string clusterName = "kustoCluster2";
ResourceIdentifier kustoClusterResourceId = KustoClusterResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, clusterName);
KustoClusterResource kustoCluster = client.GetKustoClusterResource(kustoClusterResourceId);

// get the collection of this KustoAttachedDatabaseConfigurationResource
KustoAttachedDatabaseConfigurationCollection collection = kustoCluster.GetKustoAttachedDatabaseConfigurations();

// invoke the operation
string attachedDatabaseConfigurationName = "attachedDatabaseConfigurationsTest";
bool result = await collection.ExistsAsync(attachedDatabaseConfigurationName);

Console.WriteLine($"Succeeded: {result}");
