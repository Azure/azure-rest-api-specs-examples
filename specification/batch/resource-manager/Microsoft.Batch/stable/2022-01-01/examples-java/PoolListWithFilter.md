Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-batch_1.0.0/sdk/batch/azure-resourcemanager-batch/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for Pool ListByBatchAccount. */
public final class Main {
    /*
     * x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2022-01-01/examples/PoolListWithFilter.json
     */
    /**
     * Sample code: ListPoolWithFilter.
     *
     * @param manager Entry point to BatchManager.
     */
    public static void listPoolWithFilter(com.azure.resourcemanager.batch.BatchManager manager) {
        manager
            .pools()
            .listByBatchAccount(
                "default-azurebatch-japaneast",
                "sampleacct",
                50,
                "properties/allocationState,properties/provisioningStateTransitionTime,properties/currentDedicatedNodes,properties/currentLowPriorityNodes",
                "startswith(name, 'po') or (properties/allocationState eq 'Steady' and"
                    + " properties/provisioningStateTransitionTime lt datetime'2017-02-02')",
                Context.NONE);
    }
}
```
