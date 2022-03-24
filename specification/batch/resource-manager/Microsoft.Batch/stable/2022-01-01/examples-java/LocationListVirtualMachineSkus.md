Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-batch_1.0.0/sdk/batch/azure-resourcemanager-batch/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for Location ListSupportedVirtualMachineSkus. */
public final class Main {
    /*
     * x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2022-01-01/examples/LocationListVirtualMachineSkus.json
     */
    /**
     * Sample code: LocationListVirtualMachineSkus.
     *
     * @param manager Entry point to BatchManager.
     */
    public static void locationListVirtualMachineSkus(com.azure.resourcemanager.batch.BatchManager manager) {
        manager.locations().listSupportedVirtualMachineSkus("japaneast", null, null, Context.NONE);
    }
}
```
