
/**
 * Samples for HcpOpenShiftClusters RevokeCredentials.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-09-01-preview/HcpOpenShiftClusters_RevokeCredentials_MaximumSet_Gen.json
     */
    /**
     * Sample code: HcpOpenShiftClusters_RevokeCredentials_MaximumSet.
     * 
     * @param manager Entry point to RedHatOpenShiftHostedControlPlanesManager.
     */
    public static void hcpOpenShiftClustersRevokeCredentialsMaximumSet(
        com.azure.resourcemanager.redhatopenshifthcp.RedHatOpenShiftHostedControlPlanesManager manager) {
        manager.hcpOpenShiftClusters().revokeCredentials("rgopenapi", "hcpCluster-name",
            com.azure.core.util.Context.NONE);
    }
}
