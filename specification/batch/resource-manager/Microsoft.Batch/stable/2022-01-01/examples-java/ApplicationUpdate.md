```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.batch.models.Application;

/** Samples for Application Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2022-01-01/examples/ApplicationUpdate.json
     */
    /**
     * Sample code: ApplicationUpdate.
     *
     * @param manager Entry point to BatchManager.
     */
    public static void applicationUpdate(com.azure.resourcemanager.batch.BatchManager manager) {
        Application resource =
            manager
                .applications()
                .getWithResponse("default-azurebatch-japaneast", "sampleacct", "app1", Context.NONE)
                .getValue();
        resource.update().withDisplayName("myAppName").withAllowUpdates(true).withDefaultVersion("2").apply();
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-batch_1.0.0/sdk/batch/azure-resourcemanager-batch/README.md) on how to add the SDK to your project and authenticate.
