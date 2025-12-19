
/**
 * Samples for Datastores List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/Datastores_List.json
     */
    /**
     * Sample code: Datastores_List.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void datastoresList(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.datastores().list("group1", "cloud1", "cluster1", com.azure.core.util.Context.NONE);
    }
}
