```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.kusto.models.BlobStorageEventType;
import com.azure.resourcemanager.kusto.models.DatabaseRouting;
import com.azure.resourcemanager.kusto.models.EventGridDataConnection;
import com.azure.resourcemanager.kusto.models.EventGridDataFormat;

/** Samples for DataConnections CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2022-02-01/examples/KustoDataConnectionsEventGridCreateOrUpdate.json
     */
    /**
     * Sample code: KustoDataConnectionsEventGridCreateOrUpdate.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void kustoDataConnectionsEventGridCreateOrUpdate(
        com.azure.resourcemanager.kusto.KustoManager manager) {
        manager
            .dataConnections()
            .createOrUpdate(
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
                Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-kusto_1.0.0-beta.4/sdk/kusto/azure-resourcemanager-kusto/README.md) on how to add the SDK to your project and authenticate.
