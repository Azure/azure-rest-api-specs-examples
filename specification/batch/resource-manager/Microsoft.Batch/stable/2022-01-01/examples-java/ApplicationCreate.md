Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-batch_1.0.0/sdk/batch/azure-resourcemanager-batch/README.md) on how to add the SDK to your project and authenticate.

```java
/** Samples for Application Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2022-01-01/examples/ApplicationCreate.json
     */
    /**
     * Sample code: ApplicationCreate.
     *
     * @param manager Entry point to BatchManager.
     */
    public static void applicationCreate(com.azure.resourcemanager.batch.BatchManager manager) {
        manager
            .applications()
            .define("app1")
            .withExistingBatchAccount("default-azurebatch-japaneast", "sampleacct")
            .withDisplayName("myAppName")
            .withAllowUpdates(false)
            .create();
    }
}
```
