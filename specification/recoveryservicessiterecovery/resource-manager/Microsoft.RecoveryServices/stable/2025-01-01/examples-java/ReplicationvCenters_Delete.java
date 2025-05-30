
/**
 * Samples for ReplicationvCenters Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2025-01-01/examples
     * /ReplicationvCenters_Delete.json
     */
    /**
     * Sample code: Remove vCenter operation.
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void
        removeVCenterOperation(com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager.replicationvCenters().delete("MadhaviVRG", "MadhaviVault", "MadhaviFabric", "esx-78",
            com.azure.core.util.Context.NONE);
    }
}
