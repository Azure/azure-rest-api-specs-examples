
/**
 * Samples for Application List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/ApplicationList.json
     */
    /**
     * Sample code: ApplicationList.
     * 
     * @param manager Entry point to BatchManager.
     */
    public static void applicationList(com.azure.resourcemanager.batch.BatchManager manager) {
        manager.applications().list("default-azurebatch-japaneast", "sampleacct", null,
            com.azure.core.util.Context.NONE);
    }
}
