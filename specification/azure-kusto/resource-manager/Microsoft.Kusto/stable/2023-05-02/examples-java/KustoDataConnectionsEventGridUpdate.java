import com.azure.resourcemanager.kusto.models.BlobStorageEventType;
import com.azure.resourcemanager.kusto.models.DatabaseRouting;
import com.azure.resourcemanager.kusto.models.EventGridDataConnection;
import com.azure.resourcemanager.kusto.models.EventGridDataFormat;

/** Samples for DataConnections Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2023-05-02/examples/KustoDataConnectionsEventGridUpdate.json
     */
    /**
     * Sample code: KustoDataConnectionsEventGridUpdate.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void kustoDataConnectionsEventGridUpdate(com.azure.resourcemanager.kusto.KustoManager manager) {
        manager
            .dataConnections()
            .update(
                "kustorptest",
                "kustoCluster",
                "KustoDatabase8",
                "dataConnectionTest",
                new EventGridDataConnection()
                    .withLocation("westus")
                    .withStorageAccountResourceId(
                        "/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Storage/storageAccounts/teststorageaccount")
                    .withEventGridResourceId(
                        "/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Storage/storageAccounts/teststorageaccount/providers/Microsoft.EventGrid/eventSubscriptions/eventSubscriptionTest")
                    .withEventHubResourceId(
                        "/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.EventHub/namespaces/eventhubTestns1/eventhubs/eventhubTest2")
                    .withConsumerGroup("$Default")
                    .withTableName("TestTable")
                    .withMappingRuleName("TestMapping")
                    .withDataFormat(EventGridDataFormat.JSON)
                    .withIgnoreFirstRecord(false)
                    .withBlobStorageEventType(BlobStorageEventType.MICROSOFT_STORAGE_BLOB_CREATED)
                    .withManagedIdentityResourceId(
                        "/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.ManagedIdentity/userAssignedIdentities/managedidentityTest1")
                    .withDatabaseRouting(DatabaseRouting.SINGLE),
                com.azure.core.util.Context.NONE);
    }
}
