import com.azure.resourcemanager.kusto.fluent.models.DataConnectionValidationInner;
import com.azure.resourcemanager.kusto.models.Compression;
import com.azure.resourcemanager.kusto.models.EventHubDataConnection;
import com.azure.resourcemanager.kusto.models.EventHubDataFormat;

/** Samples for DataConnections DataConnectionValidation. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2023-08-15/examples/KustoDataConnectionValidationAsync.json
     */
    /**
     * Sample code: KustoDataConnectionValidation.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void kustoDataConnectionValidation(com.azure.resourcemanager.kusto.KustoManager manager) {
        manager
            .dataConnections()
            .dataConnectionValidation(
                "kustorptest",
                "kustoCluster",
                "KustoDatabase8",
                new DataConnectionValidationInner()
                    .withDataConnectionName("dataConnectionTest")
                    .withProperties(
                        new EventHubDataConnection()
                            .withEventHubResourceId(
                                "/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.EventHub/namespaces/eventhubTestns1/eventhubs/eventhubTest1")
                            .withConsumerGroup("testConsumerGroup1")
                            .withTableName("TestTable")
                            .withMappingRuleName("TestMapping")
                            .withDataFormat(EventHubDataFormat.JSON)
                            .withCompression(Compression.NONE)
                            .withManagedIdentityResourceId(
                                "/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.ManagedIdentity/userAssignedIdentities/managedidentityTest1")),
                com.azure.core.util.Context.NONE);
    }
}
