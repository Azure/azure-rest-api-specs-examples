
/**
 * Samples for NodePools Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-09-01-preview/NodePools_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: NodePools_Get.
     * 
     * @param manager Entry point to RedHatOpenShiftHostedControlPlanesManager.
     */
    public static void
        nodePoolsGet(com.azure.resourcemanager.redhatopenshifthcp.RedHatOpenShiftHostedControlPlanesManager manager) {
        manager.nodePools().getWithResponse("rgopenapi", "hcpCluster-name", "nodepool-name",
            com.azure.core.util.Context.NONE);
    }
}
