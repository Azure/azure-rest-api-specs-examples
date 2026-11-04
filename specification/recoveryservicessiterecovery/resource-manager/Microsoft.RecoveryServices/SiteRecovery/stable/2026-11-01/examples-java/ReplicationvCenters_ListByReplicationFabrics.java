
/**
 * Samples for ReplicationvCenters ListByReplicationFabrics.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/ReplicationvCenters_ListByReplicationFabrics.json
     */
    /**
     * Sample code: Gets the list of vCenter registered under a fabric.
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void getsTheListOfVCenterRegisteredUnderAFabric(
        com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager.replicationvCenters().listByReplicationFabrics("MadhaviVRG", "MadhaviVault", "MadhaviFabric",
            com.azure.core.util.Context.NONE);
    }
}
