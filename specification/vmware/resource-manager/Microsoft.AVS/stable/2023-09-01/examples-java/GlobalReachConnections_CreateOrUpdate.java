
/**
 * Samples for GlobalReachConnections CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2023-09-01/examples/
     * GlobalReachConnections_CreateOrUpdate.json
     */
    /**
     * Sample code: GlobalReachConnections_CreateOrUpdate.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void globalReachConnectionsCreateOrUpdate(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.globalReachConnections().define("connection1").withExistingPrivateCloud("group1", "cloud1")
            .withAuthorizationKey("01010101-0101-0101-0101-010101010101")
            .withPeerExpressRouteCircuit(
                "/subscriptions/12341234-1234-1234-1234-123412341234/resourceGroups/mygroup/providers/Microsoft.Network/expressRouteCircuits/mypeer")
            .create();
    }
}
