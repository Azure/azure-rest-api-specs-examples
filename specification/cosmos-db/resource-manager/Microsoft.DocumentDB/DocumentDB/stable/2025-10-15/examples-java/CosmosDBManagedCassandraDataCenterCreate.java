
import com.azure.resourcemanager.cosmos.fluent.models.DataCenterResourceInner;
import com.azure.resourcemanager.cosmos.models.DataCenterResourceProperties;

/**
 * Samples for CassandraDataCenters CreateUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cosmos-db/resource-manager/Microsoft.DocumentDB/DocumentDB/stable/2025-10-15/examples/
     * CosmosDBManagedCassandraDataCenterCreate.json
     */
    /**
     * Sample code: CosmosDBManagedCassandraDataCenterCreate.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBManagedCassandraDataCenterCreate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cosmosDBAccounts().manager().serviceClient().getCassandraDataCenters().createUpdate("cassandra-prod-rg",
            "cassandra-prod", "dc1",
            new DataCenterResourceInner().withProperties(new DataCenterResourceProperties()
                .withDataCenterLocation("West US 2")
                .withDelegatedSubnetId(
                    "/subscriptions/536e130b-d7d6-4ac7-98a5-de20d69588d2/resourceGroups/customer-vnet-rg/providers/Microsoft.Network/virtualNetworks/customer-vnet/subnets/dc1-subnet")
                .withNodeCount(9).withBase64EncodedCassandraYamlFragment("fakeTokenPlaceholder")),
            com.azure.core.util.Context.NONE);
    }
}
