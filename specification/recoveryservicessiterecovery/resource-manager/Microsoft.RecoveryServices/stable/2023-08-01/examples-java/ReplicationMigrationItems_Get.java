
/**
 * Samples for ReplicationMigrationItems Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-08-01/examples
     * /ReplicationMigrationItems_Get.json
     */
    /**
     * Sample code: Gets the details of a migration item.
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void getsTheDetailsOfAMigrationItem(
        com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager.replicationMigrationItems().getWithResponse("migrationvault", "resourcegroup1", "vmwarefabric1",
            "vmwareContainer1", "virtualmachine1", com.azure.core.util.Context.NONE);
    }
}
