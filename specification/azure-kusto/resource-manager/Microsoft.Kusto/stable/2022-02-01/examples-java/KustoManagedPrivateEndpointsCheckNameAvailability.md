Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-kusto_1.0.0-beta.4/sdk/kusto/azure-resourcemanager-kusto/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.kusto.models.ManagedPrivateEndpointsCheckNameRequest;

/** Samples for ManagedPrivateEndpoints CheckNameAvailability. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2022-02-01/examples/KustoManagedPrivateEndpointsCheckNameAvailability.json
     */
    /**
     * Sample code: KustoManagedPrivateEndpointsCheckNameAvailability.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void kustoManagedPrivateEndpointsCheckNameAvailability(
        com.azure.resourcemanager.kusto.KustoManager manager) {
        manager
            .managedPrivateEndpoints()
            .checkNameAvailabilityWithResponse(
                "kustorptest",
                "kustoCluster",
                new ManagedPrivateEndpointsCheckNameRequest().withName("pme1"),
                Context.NONE);
    }
}
```
