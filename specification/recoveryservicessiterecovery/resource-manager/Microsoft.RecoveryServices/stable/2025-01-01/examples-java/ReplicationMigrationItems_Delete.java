
/**
 * Samples for ReplicationMigrationItems Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2025-01-01/examples
     * /ReplicationMigrationItems_Delete.json
     */
    /**
     * Sample code: Delete the migration item.
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void
        deleteTheMigrationItem(com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager.replicationMigrationItems().delete("resourcegroup1", "migrationvault", "vmwarefabric1",
            "vmwareContainer1", "virtualmachine1", null, com.azure.core.util.Context.NONE);
    }
}
