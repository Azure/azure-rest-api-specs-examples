
/**
 * Samples for Application Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/ApplicationCreate.json
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
