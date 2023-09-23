import com.azure.resourcemanager.recoveryservicessiterecovery.models.ResumeReplicationInput;
import com.azure.resourcemanager.recoveryservicessiterecovery.models.ResumeReplicationInputProperties;
import com.azure.resourcemanager.recoveryservicessiterecovery.models.VMwareCbtResumeReplicationInput;

/** Samples for ReplicationMigrationItems ResumeReplication. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-06-01/examples/ReplicationMigrationItems_ResumeReplication.json
     */
    /**
     * Sample code: Resume replication.
     *
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void resumeReplication(
        com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager
            .replicationMigrationItems()
            .resumeReplication(
                "migrationvault",
                "resourcegroup1",
                "vmwarefabric1",
                "vmwareContainer1",
                "virtualmachine1",
                new ResumeReplicationInput()
                    .withProperties(
                        new ResumeReplicationInputProperties()
                            .withProviderSpecificDetails(
                                new VMwareCbtResumeReplicationInput().withDeleteMigrationResources("false"))),
                com.azure.core.util.Context.NONE);
    }
}
