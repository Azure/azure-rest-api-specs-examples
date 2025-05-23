
/**
 * Samples for Jobs List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/databox/resource-manager/Microsoft.DataBox/stable/2025-02-01/examples/JobsList.json
     */
    /**
     * Sample code: JobsList.
     * 
     * @param manager Entry point to DataBoxManager.
     */
    public static void jobsList(com.azure.resourcemanager.databox.DataBoxManager manager) {
        manager.jobs().list(null, com.azure.core.util.Context.NONE);
    }
}
