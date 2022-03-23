Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-batch_1.0.0/sdk/batch/azure-resourcemanager-batch/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for ApplicationPackage Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2022-01-01/examples/ApplicationPackageGet.json
     */
    /**
     * Sample code: ApplicationPackageGet.
     *
     * @param manager Entry point to BatchManager.
     */
    public static void applicationPackageGet(com.azure.resourcemanager.batch.BatchManager manager) {
        manager
            .applicationPackages()
            .getWithResponse("default-azurebatch-japaneast", "sampleacct", "app1", "1", Context.NONE);
    }
}
```
