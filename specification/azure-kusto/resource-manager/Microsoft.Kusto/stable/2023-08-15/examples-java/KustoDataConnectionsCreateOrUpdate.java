import com.azure.resourcemanager.kusto.models.EventHubDataConnection;

/** Samples for DataConnections CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2023-08-15/examples/KustoDataConnectionsCreateOrUpdate.json
     */
    /**
     * Sample code: KustoDataConnectionsCreateOrUpdate.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void kustoDataConnectionsCreateOrUpdate(com.azure.resourcemanager.kusto.KustoManager manager) {
        manager
            .dataConnections()
            .createOrUpdate(
                "kustorptest",
                "kustoCluster",
                "KustoDatabase8",
                "dataConnectionTest",
                new EventHubDataConnection()
                    .withLocation("westus")
                    .withEventHubResourceId(
                        "/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.EventHub/namespaces/eventhubTestns1/eventhubs/eventhubTest1")
                    .withConsumerGroup("testConsumerGroup1")
                    .withManagedIdentityResourceId(
                        "/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.ManagedIdentity/userAssignedIdentities/managedidentityTest1"),
                com.azure.core.util.Context.NONE);
    }
}
