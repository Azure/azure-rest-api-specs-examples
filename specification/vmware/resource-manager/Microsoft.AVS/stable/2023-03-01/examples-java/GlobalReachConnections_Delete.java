/** Samples for GlobalReachConnections Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2023-03-01/examples/GlobalReachConnections_Delete.json
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
