
/**
 * Samples for Datastores Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/Datastores_Delete.json
     */
    /**
     * Sample code: Datastores_Delete.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void datastoresDelete(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.datastores().delete("group1", "cloud1", "cluster1", "datastore1", com.azure.core.util.Context.NONE);
    }
}
