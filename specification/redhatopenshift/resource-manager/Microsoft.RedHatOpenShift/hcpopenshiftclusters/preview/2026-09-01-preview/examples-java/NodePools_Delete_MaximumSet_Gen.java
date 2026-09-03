
/**
 * Samples for NodePools Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-09-01-preview/NodePools_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: NodePools_Delete.
     * 
     * @param manager Entry point to RedHatOpenShiftHostedControlPlanesManager.
     */
    public static void nodePoolsDelete(
        com.azure.resourcemanager.redhatopenshifthcp.RedHatOpenShiftHostedControlPlanesManager manager) {
        manager.nodePools().delete("rgopenapi", "hcpCluster-name", "nodePool-name", com.azure.core.util.Context.NONE);
    }
}
