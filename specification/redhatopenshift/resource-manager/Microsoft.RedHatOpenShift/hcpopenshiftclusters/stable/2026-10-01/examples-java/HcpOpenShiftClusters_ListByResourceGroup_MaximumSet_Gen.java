
/**
 * Samples for HcpOpenShiftClusters ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/HcpOpenShiftClusters_ListByResourceGroup_MaximumSet_Gen.json
     */
    /**
     * Sample code: HcpOpenShiftClusters_ListByResourceGroup.
     * 
     * @param manager Entry point to RedHatOpenShiftHostedControlPlanesManager.
     */
    public static void hcpOpenShiftClustersListByResourceGroup(
        com.azure.resourcemanager.redhatopenshifthcp.RedHatOpenShiftHostedControlPlanesManager manager) {
        manager.hcpOpenShiftClusters().listByResourceGroup("rgopenapi", com.azure.core.util.Context.NONE);
    }
}
