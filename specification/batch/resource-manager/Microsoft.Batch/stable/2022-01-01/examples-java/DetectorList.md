```java
import com.azure.core.util.Context;

/** Samples for BatchAccount ListDetectors. */
public final class Main {
    /*
     * x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2022-01-01/examples/DetectorList.json
     */
    /**
     * Sample code: ListDetectors.
     *
     * @param manager Entry point to BatchManager.
     */
    public static void listDetectors(com.azure.resourcemanager.batch.BatchManager manager) {
        manager.batchAccounts().listDetectors("default-azurebatch-japaneast", "sampleacct", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-batch_1.0.0/sdk/batch/azure-resourcemanager-batch/README.md) on how to add the SDK to your project and authenticate.
