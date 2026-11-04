
import com.azure.resourcemanager.recoveryservicessiterecovery.models.PauseReplicationInput;
import com.azure.resourcemanager.recoveryservicessiterecovery.models.PauseReplicationInputProperties;

/**
 * Samples for ReplicationMigrationItems PauseReplication.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/ReplicationMigrationItems_PauseReplication.json
     */
    /**
     * Sample code: Pause replication.
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void
        pauseReplication(com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager.replicationMigrationItems()
            .pauseReplication("resourcegroup1", "migrationvault", "vmwarefabric1", "vmwareContainer1",
                "virtualmachine1",
                new PauseReplicationInput()
                    .withProperties(new PauseReplicationInputProperties().withInstanceType("VMwareCbt")),
                com.azure.core.util.Context.NONE);
    }
}
