
/**
 * Samples for ReplicationMigrationItems Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/ReplicationMigrationItems_Get.json
     */
    /**
     * Sample code: Gets the details of a migration item.
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void getsTheDetailsOfAMigrationItem(
        com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager.replicationMigrationItems().getWithResponse("resourcegroup1", "migrationvault", "vmwarefabric1",
            "vmwareContainer1", "virtualmachine1", com.azure.core.util.Context.NONE);
    }
}
