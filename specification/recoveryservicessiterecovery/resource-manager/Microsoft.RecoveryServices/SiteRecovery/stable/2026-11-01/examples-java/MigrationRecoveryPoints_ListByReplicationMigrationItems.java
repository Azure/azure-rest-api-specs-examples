
/**
 * Samples for MigrationRecoveryPoints ListByReplicationMigrationItems.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/MigrationRecoveryPoints_ListByReplicationMigrationItems.json
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
