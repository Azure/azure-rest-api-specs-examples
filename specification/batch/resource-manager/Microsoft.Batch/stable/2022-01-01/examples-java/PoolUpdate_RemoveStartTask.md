```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.batch.models.Pool;
import com.azure.resourcemanager.batch.models.StartTask;

/** Samples for Pool Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2022-01-01/examples/PoolUpdate_RemoveStartTask.json
     */
    /**
     * Sample code: UpdatePool - Remove Start Task.
     *
     * @param manager Entry point to BatchManager.
     */
    public static void updatePoolRemoveStartTask(com.azure.resourcemanager.batch.BatchManager manager) {
        Pool resource =
            manager
                .pools()
                .getWithResponse("default-azurebatch-japaneast", "sampleacct", "testpool", Context.NONE)
                .getValue();
        resource.update().withStartTask(new StartTask()).apply();
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-batch_1.0.0/sdk/batch/azure-resourcemanager-batch/README.md) on how to add the SDK to your project and authenticate.
