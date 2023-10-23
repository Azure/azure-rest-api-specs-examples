import com.azure.resourcemanager.kusto.fluent.models.DataConnectionValidationInner;
import com.azure.resourcemanager.kusto.models.BlobStorageEventType;
import com.azure.resourcemanager.kusto.models.DatabaseRouting;
import com.azure.resourcemanager.kusto.models.EventGridDataConnection;
import com.azure.resourcemanager.kusto.models.EventGridDataFormat;

/** Samples for DataConnections DataConnectionValidation. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2023-08-15/examples/KustoDataConnectionEventGridValidationAsync.json
     */
    /**
     * Sample code: KustoDataConnectionEventGridValidation.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void kustoDataConnectionEventGridValidation(com.azure.resourcemanager.kusto.KustoManager manager) {
        manager
            .dataConnections()
            .dataConnectionValidation(
                "kustorptest",
                "kustoCluster",
                "KustoDatabase8",
                new DataConnectionValidationInner()
                    .withDataConnectionName("dataConnectionTest")
                    .withProperties(
                        new EventGridDataConnection()
                            .withStorageAccountResourceId(
                                "/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Storage/storageAccounts/teststorageaccount")
                            .withEventGridResourceId(
                                "/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Storage/storageAccounts/teststorageaccount/providers/Microsoft.EventGrid/eventSubscriptions/eventSubscriptionTest")
                            .withEventHubResourceId(
                                "/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.EventHub/namespaces/eventhubTestns1/eventhubs/eventhubTest1")
                            .withConsumerGroup("$Default")
                            .withTableName("TestTable")
                            .withMappingRuleName("TestMapping")
                            .withDataFormat(EventGridDataFormat.JSON)
                            .withIgnoreFirstRecord(false)
                            .withBlobStorageEventType(BlobStorageEventType.MICROSOFT_STORAGE_BLOB_CREATED)
                            .withManagedIdentityResourceId(
                                "/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.ManagedIdentity/userAssignedIdentities/managedidentityTest1")
                            .withDatabaseRouting(DatabaseRouting.SINGLE)),
                com.azure.core.util.Context.NONE);
    }
}
