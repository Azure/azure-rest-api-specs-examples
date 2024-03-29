
/**
 * Samples for Application Create.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/batch/resource-manager/Microsoft.Batch/stable/2024-02-01/examples/ApplicationCreate.json
     */
    /**
     * Sample code: ApplicationCreate.
     * 
     * @param manager Entry point to BatchManager.
     */
    public static void applicationCreate(com.azure.resourcemanager.batch.BatchManager manager) {
        manager.applications().define("app1").withExistingBatchAccount("default-azurebatch-japaneast", "sampleacct")
            .withDisplayName("myAppName").withAllowUpdates(false).create();
    }
}
