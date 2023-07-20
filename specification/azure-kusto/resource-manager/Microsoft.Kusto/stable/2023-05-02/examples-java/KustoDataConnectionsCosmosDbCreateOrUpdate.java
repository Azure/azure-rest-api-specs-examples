import com.azure.resourcemanager.kusto.models.CosmosDbDataConnection;
import java.time.OffsetDateTime;

/** Samples for DataConnections CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2023-05-02/examples/KustoDataConnectionsCosmosDbCreateOrUpdate.json
     */
    /**
     * Sample code: KustoDataConnectionsCosmosDbCreateOrUpdate.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void kustoDataConnectionsCosmosDbCreateOrUpdate(
        com.azure.resourcemanager.kusto.KustoManager manager) {
        manager
            .dataConnections()
            .createOrUpdate(
                "kustorptest",
                "kustoCluster",
                "KustoDatabase1",
                "dataConnectionTest",
                new CosmosDbDataConnection()
                    .withLocation("westus")
                    .withTableName("TestTable")
                    .withMappingRuleName("TestMapping")
                    .withManagedIdentityResourceId(
                        "/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.ManagedIdentity/userAssignedIdentities/managedidentityTest1")
                    .withCosmosDbAccountResourceId(
                        "/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.DocumentDb/databaseAccounts/cosmosDbAccountTest1")
                    .withCosmosDbDatabase("cosmosDbDatabaseTest")
                    .withCosmosDbContainer("cosmosDbContainerTest")
                    .withRetrievalStartDate(OffsetDateTime.parse("2022-07-29T12:00:00.6554616Z")),
                com.azure.core.util.Context.NONE);
    }
}
