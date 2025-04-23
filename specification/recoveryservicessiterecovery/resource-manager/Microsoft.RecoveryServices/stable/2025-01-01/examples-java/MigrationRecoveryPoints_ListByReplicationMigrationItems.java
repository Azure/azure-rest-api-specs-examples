
/**
 * Samples for MigrationRecoveryPoints ListByReplicationMigrationItems.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2025-01-01/examples
     * /MigrationRecoveryPoints_ListByReplicationMigrationItems.json
     */
    /**
     * Sample code: Gets the recovery points for a migration item.
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void getsTheRecoveryPointsForAMigrationItem(
        com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager.migrationRecoveryPoints().listByReplicationMigrationItems("resourcegroup1", "migrationvault",
            "vmwarefabric1", "vmwareContainer1", "virtualmachine1", com.azure.core.util.Context.NONE);
    }
}
