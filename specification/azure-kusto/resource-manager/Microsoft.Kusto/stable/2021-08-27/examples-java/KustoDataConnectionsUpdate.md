Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-kusto_1.0.0-beta.3/sdk/kusto/azure-resourcemanager-kusto/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.kusto.models.EventHubDataConnection;

/** Samples for DataConnections Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2021-08-27/examples/KustoDataConnectionsUpdate.json
     */
    /**
     * Sample code: KustoDataConnectionsUpdate.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void kustoDataConnectionsUpdate(com.azure.resourcemanager.kusto.KustoManager manager) {
        manager
            .dataConnections()
            .update(
                "kustorptest",
                "kustoclusterrptest4",
                "KustoDatabase8",
                "DataConnections8",
                new EventHubDataConnection()
                    .withLocation("westus")
                    .withEventHubResourceId(
                        "/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.EventHub/namespaces/eventhubTestns1/eventhubs/eventhubTest1")
                    .withConsumerGroup("testConsumerGroup1")
                    .withManagedIdentityResourceId(
                        "/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.ManagedIdentity/userAssignedIdentities/managedidentityTest1"),
                Context.NONE);
    }
}
```
