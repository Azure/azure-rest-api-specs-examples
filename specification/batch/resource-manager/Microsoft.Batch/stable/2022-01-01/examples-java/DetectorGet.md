```java
import com.azure.core.util.Context;

/** Samples for BatchAccount GetDetector. */
public final class Main {
    /*
     * x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2022-01-01/examples/DetectorGet.json
     */
    /**
     * Sample code: GetDetector.
     *
     * @param manager Entry point to BatchManager.
     */
    public static void getDetector(com.azure.resourcemanager.batch.BatchManager manager) {
        manager
            .batchAccounts()
            .getDetectorWithResponse("default-azurebatch-japaneast", "sampleacct", "poolsAndNodes", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-batch_1.0.0/sdk/batch/azure-resourcemanager-batch/README.md) on how to add the SDK to your project and authenticate.
