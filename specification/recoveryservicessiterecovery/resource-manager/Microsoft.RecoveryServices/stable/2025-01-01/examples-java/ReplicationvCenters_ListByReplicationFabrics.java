
/**
 * Samples for ReplicationvCenters ListByReplicationFabrics.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2025-01-01/examples
     * /ReplicationvCenters_ListByReplicationFabrics.json
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
