import com.azure.resourcemanager.recoveryservicessiterecovery.models.ResyncInput;
import com.azure.resourcemanager.recoveryservicessiterecovery.models.ResyncInputProperties;
import com.azure.resourcemanager.recoveryservicessiterecovery.models.VMwareCbtResyncInput;

/** Samples for ReplicationMigrationItems Resync. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-06-01/examples/ReplicationMigrationItems_Resync.json
     */
    /**
     * Sample code: Resynchronizes replication.
     *
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void resynchronizesReplication(
        com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager
            .replicationMigrationItems()
            .resync(
                "migrationvault",
                "resourcegroup1",
                "vmwarefabric1",
                "vmwareContainer1",
                "virtualmachine1",
                new ResyncInput()
                    .withProperties(
                        new ResyncInputProperties()
                            .withProviderSpecificDetails(new VMwareCbtResyncInput().withSkipCbtReset("true"))),
                com.azure.core.util.Context.NONE);
    }
}
