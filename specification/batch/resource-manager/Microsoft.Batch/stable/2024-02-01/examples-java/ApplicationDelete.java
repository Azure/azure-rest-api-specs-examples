
/**
 * Samples for Application Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/batch/resource-manager/Microsoft.Batch/stable/2024-02-01/examples/ApplicationDelete.json
     */
    /**
     * Sample code: ApplicationDelete.
     * 
     * @param manager Entry point to BatchManager.
     */
    public static void applicationDelete(com.azure.resourcemanager.batch.BatchManager manager) {
        manager.applications().deleteWithResponse("default-azurebatch-japaneast", "sampleacct", "app1",
            com.azure.core.util.Context.NONE);
    }
}
