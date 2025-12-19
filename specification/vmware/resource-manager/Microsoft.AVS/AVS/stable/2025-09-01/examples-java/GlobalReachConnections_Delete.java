
/**
 * Samples for GlobalReachConnections Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/GlobalReachConnections_Delete.json
     */
    /**
     * Sample code: GlobalReachConnections_Delete.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void globalReachConnectionsDelete(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.globalReachConnections().delete("group1", "cloud1", "connection1", com.azure.core.util.Context.NONE);
    }
}
