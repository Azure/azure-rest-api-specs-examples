
import com.azure.resourcemanager.recoveryservicessiterecovery.models.TestMigrateCleanupInput;
import com.azure.resourcemanager.recoveryservicessiterecovery.models.TestMigrateCleanupInputProperties;

/**
 * Samples for ReplicationMigrationItems TestMigrateCleanup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/ReplicationMigrationItems_TestMigrateCleanup.json
     */
    /**
     * Sample code: Test migrate cleanup.
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void
        testMigrateCleanup(com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager.replicationMigrationItems().testMigrateCleanup("resourcegroup1", "migrationvault", "vmwarefabric1",
            "vmwareContainer1", "virtualmachine1",
            new TestMigrateCleanupInput()
                .withProperties(new TestMigrateCleanupInputProperties().withComments("Test Failover Cleanup")),
            com.azure.core.util.Context.NONE);
    }
}
