
/**
 * Samples for Application Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/batch/resource-manager/Microsoft.Batch/stable/2024-07-01/examples/ApplicationGet.json
     */
    /**
     * Sample code: ApplicationGet.
     * 
     * @param manager Entry point to BatchManager.
     */
    public static void applicationGet(com.azure.resourcemanager.batch.BatchManager manager) {
        manager.applications().getWithResponse("default-azurebatch-japaneast", "sampleacct", "app1",
            com.azure.core.util.Context.NONE);
    }
}
