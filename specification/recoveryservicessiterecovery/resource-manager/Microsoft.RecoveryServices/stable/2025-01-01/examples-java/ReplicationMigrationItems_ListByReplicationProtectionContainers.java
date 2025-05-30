
/**
 * Samples for ReplicationMigrationItems ListByReplicationProtectionContainers.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2025-01-01/examples
     * /ReplicationMigrationItems_ListByReplicationProtectionContainers.json
     */
    /**
     * Sample code: Gets the list of migration items in the protection container.
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void getsTheListOfMigrationItemsInTheProtectionContainer(
        com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager.replicationMigrationItems().listByReplicationProtectionContainers("resourcegroup1", "migrationvault",
            "vmwarefabric1", "vmwareContainer1", null, null, null, com.azure.core.util.Context.NONE);
    }
}
