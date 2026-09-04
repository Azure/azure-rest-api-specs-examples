
/**
 * Samples for NodePools ListByParent.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-09-01-preview/NodePools_ListByParent_MaximumSet_Gen.json
     */
    /**
     * Sample code: NodePools_ListByParent.
     * 
     * @param manager Entry point to RedHatOpenShiftHostedControlPlanesManager.
     */
    public static void nodePoolsListByParent(
        com.azure.resourcemanager.redhatopenshifthcp.RedHatOpenShiftHostedControlPlanesManager manager) {
        manager.nodePools().listByParent("rgopenapi", "hcpCluster-name", com.azure.core.util.Context.NONE);
    }
}
