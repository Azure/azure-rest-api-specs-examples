```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.batch.models.AccountKeyType;
import com.azure.resourcemanager.batch.models.BatchAccountRegenerateKeyParameters;

/** Samples for BatchAccount RegenerateKey. */
public final class Main {
    /*
     * x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2022-01-01/examples/BatchAccountRegenerateKey.json
     */
    /**
     * Sample code: BatchAccountRegenerateKey.
     *
     * @param manager Entry point to BatchManager.
     */
    public static void batchAccountRegenerateKey(com.azure.resourcemanager.batch.BatchManager manager) {
        manager
            .batchAccounts()
            .regenerateKeyWithResponse(
                "default-azurebatch-japaneast",
                "sampleacct",
                new BatchAccountRegenerateKeyParameters().withKeyName(AccountKeyType.PRIMARY),
                Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-batch_1.0.0/sdk/batch/azure-resourcemanager-batch/README.md) on how to add the SDK to your project and authenticate.
