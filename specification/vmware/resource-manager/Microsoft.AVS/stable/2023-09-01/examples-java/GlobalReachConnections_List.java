
/**
 * Samples for GlobalReachConnections List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/vmware/resource-manager/Microsoft.AVS/stable/2023-09-01/examples/GlobalReachConnections_List.json
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
