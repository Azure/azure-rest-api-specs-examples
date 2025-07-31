
/**
 * Samples for ReplicationExtension Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/ReplicationExtension_Delete.json
     */
    /**
     * Sample code: Deletes the replication extension.
     * 
     * @param manager Entry point to RecoveryServicesDataReplicationManager.
     */
    public static void deletesTheReplicationExtension(
        com.azure.resourcemanager.recoveryservicesdatareplication.RecoveryServicesDataReplicationManager manager) {
        manager.replicationExtensions().delete("rgrecoveryservicesdatareplication", "4", "g16yjJ",
            com.azure.core.util.Context.NONE);
    }
}
