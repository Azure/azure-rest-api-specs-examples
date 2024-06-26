
/**
 * Samples for GlobalReachConnections Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/vmware/resource-manager/Microsoft.AVS/stable/2023-09-01/examples/GlobalReachConnections_Get.json
     */
    /**
     * Sample code: GlobalReachConnections_Get.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void globalReachConnectionsGet(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.globalReachConnections().getWithResponse("group1", "cloud1", "connection1",
            com.azure.core.util.Context.NONE);
    }
}
