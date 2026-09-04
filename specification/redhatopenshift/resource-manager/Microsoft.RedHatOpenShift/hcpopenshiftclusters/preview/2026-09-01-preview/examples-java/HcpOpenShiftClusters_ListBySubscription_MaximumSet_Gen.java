
/**
 * Samples for HcpOpenShiftClusters List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-09-01-preview/HcpOpenShiftClusters_ListBySubscription_MaximumSet_Gen.json
     */
    /**
     * Sample code: HcpOpenShiftClusters_ListBySubscription.
     * 
     * @param manager Entry point to RedHatOpenShiftHostedControlPlanesManager.
     */
    public static void hcpOpenShiftClustersListBySubscription(
        com.azure.resourcemanager.redhatopenshifthcp.RedHatOpenShiftHostedControlPlanesManager manager) {
        manager.hcpOpenShiftClusters().list(com.azure.core.util.Context.NONE);
    }
}
