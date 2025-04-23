
/**
 * Samples for ReplicationNetworks List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2025-01-01/examples
     * /ReplicationNetworks_List.json
     */
    /**
     * Sample code: Gets the list of networks. View-only API.
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void getsTheListOfNetworksViewOnlyAPI(
        com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager.replicationNetworks().list("srcBvte2a14C27", "srce2avaultbvtaC27", com.azure.core.util.Context.NONE);
    }
}
