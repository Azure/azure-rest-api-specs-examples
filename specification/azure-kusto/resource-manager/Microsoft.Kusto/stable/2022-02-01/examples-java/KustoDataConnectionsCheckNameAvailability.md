```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.kusto.models.DataConnectionCheckNameRequest;

/** Samples for DataConnections CheckNameAvailability. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2022-02-01/examples/KustoDataConnectionsCheckNameAvailability.json
     */
    /**
     * Sample code: KustoDataConnectionsCheckNameAvailability.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void kustoDataConnectionsCheckNameAvailability(com.azure.resourcemanager.kusto.KustoManager manager) {
        manager
            .dataConnections()
            .checkNameAvailabilityWithResponse(
                "kustorptest",
                "kustoCluster",
                "KustoDatabase8",
                new DataConnectionCheckNameRequest().withName("DataConnections8"),
                Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-kusto_1.0.0-beta.4/sdk/kusto/azure-resourcemanager-kusto/README.md) on how to add the SDK to your project and authenticate.
