Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-kusto_1.0.0-beta.3/sdk/kusto/azure-resourcemanager-kusto/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for ManagedPrivateEndpoints Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2021-08-27/examples/KustoManagedPrivateEndpointsDelete.json
     */
    /**
     * Sample code: ManagedPrivateEndpointsDelete.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void managedPrivateEndpointsDelete(com.azure.resourcemanager.kusto.KustoManager manager) {
        manager
            .managedPrivateEndpoints()
            .delete("kustorptest", "kustoclusterrptest4", "kustoManagedPrivateEndpoint1", Context.NONE);
    }
}
```
