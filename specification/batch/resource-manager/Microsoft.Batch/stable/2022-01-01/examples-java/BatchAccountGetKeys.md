```java
import com.azure.core.util.Context;

/** Samples for BatchAccount GetKeys. */
public final class Main {
    /*
     * x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2022-01-01/examples/BatchAccountGetKeys.json
     */
    /**
     * Sample code: BatchAccountGetKeys.
     *
     * @param manager Entry point to BatchManager.
     */
    public static void batchAccountGetKeys(com.azure.resourcemanager.batch.BatchManager manager) {
        manager.batchAccounts().getKeysWithResponse("default-azurebatch-japaneast", "sampleacct", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-batch_1.0.0/sdk/batch/azure-resourcemanager-batch/README.md) on how to add the SDK to your project and authenticate.
