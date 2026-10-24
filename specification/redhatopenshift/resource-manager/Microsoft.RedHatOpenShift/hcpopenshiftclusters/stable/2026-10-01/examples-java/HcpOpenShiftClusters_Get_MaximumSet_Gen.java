
/**
 * Samples for HcpOpenShiftClusters GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/HcpOpenShiftClusters_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: HcpOpenShiftClusters_Get.
     * 
     * @param manager Entry point to RedHatOpenShiftHostedControlPlanesManager.
     */
    public static void hcpOpenShiftClustersGet(
        com.azure.resourcemanager.redhatopenshifthcp.RedHatOpenShiftHostedControlPlanesManager manager) {
        manager.hcpOpenShiftClusters().getByResourceGroupWithResponse("rgopenapi", "my-cool-cluster",
            com.azure.core.util.Context.NONE);
    }
}
