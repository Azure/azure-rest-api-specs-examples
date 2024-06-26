
/**
 * Samples for ReplicationMigrationItems List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-08-01/examples
     * /ReplicationMigrationItems_List.json
     */
    /**
     * Sample code: Gets the list of migration items in the vault.
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void getsTheListOfMigrationItemsInTheVault(
        com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager.replicationMigrationItems().list("migrationvault", "resourcegroup1", null, null, null,
            com.azure.core.util.Context.NONE);
    }
}
