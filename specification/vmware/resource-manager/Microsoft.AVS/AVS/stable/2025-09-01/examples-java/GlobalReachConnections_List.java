
/**
 * Samples for GlobalReachConnections List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/GlobalReachConnections_List.json
     */
    /**
     * Sample code: GlobalReachConnections_List.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void globalReachConnectionsList(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.globalReachConnections().list("group1", "cloud1", com.azure.core.util.Context.NONE);
    }
}
