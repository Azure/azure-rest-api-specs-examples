using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Kusto.Models;
using Azure.ResourceManager.Kusto;

// Generated from example definition: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2024-04-13/examples/KustoAttachedDatabaseConfigurationsCreateOrUpdate.json
// this example is just showing the usage of "AttachedDatabaseConfigurations_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

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
KustoAttachedDatabaseConfigurationData data = new KustoAttachedDatabaseConfigurationData
{
    Location = new AzureLocation("westus"),
    DatabaseName = "kustodatabase",
    ClusterResourceId = new ResourceIdentifier("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Kusto/Clusters/kustoCluster2"),
    DefaultPrincipalsModificationKind = KustoDatabaseDefaultPrincipalsModificationKind.Union,
    TableLevelSharingProperties = new KustoDatabaseTableLevelSharingProperties
    {
        TablesToInclude = { "Table1" },
        TablesToExclude = { "Table2" },
        ExternalTablesToInclude = { "ExternalTable1" },
        ExternalTablesToExclude = { "ExternalTable2" },
        MaterializedViewsToInclude = { "MaterializedViewTable1" },
        MaterializedViewsToExclude = { "MaterializedViewTable2" },
    },
    DatabaseNameOverride = "overridekustodatabase",
};
ArmOperation<KustoAttachedDatabaseConfigurationResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, attachedDatabaseConfigurationName, data);
KustoAttachedDatabaseConfigurationResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
KustoAttachedDatabaseConfigurationData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
