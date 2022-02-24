Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-kusto_1.0.0-beta.4/sdk/kusto/azure-resourcemanager-kusto/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for DataConnections Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2022-02-01/examples/KustoDataConnectionsEventGridGet.json
     */
    /**
     * Sample code: KustoDataConnectionsEventGridGet.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void kustoDataConnectionsEventGridGet(com.azure.resourcemanager.kusto.KustoManager manager) {
        manager
            .dataConnections()
            .getWithResponse("kustorptest", "kustoCluster", "KustoDatabase8", "dataConnectionTest", Context.NONE);
    }
}
```
