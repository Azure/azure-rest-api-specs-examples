
/**
 * Samples for HcpOpenShiftClusters Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/HcpOpenShiftClusters_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: HcpOpenShiftClusters_Delete.
     * 
     * @param manager Entry point to RedHatOpenShiftHostedControlPlanesManager.
     */
    public static void hcpOpenShiftClustersDelete(
        com.azure.resourcemanager.redhatopenshifthcp.RedHatOpenShiftHostedControlPlanesManager manager) {
        manager.hcpOpenShiftClusters().delete("rgopenapi", "hcpCluster-name", com.azure.core.util.Context.NONE);
    }
}
