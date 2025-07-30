
/**
 * Samples for ReplicationExtension Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/ReplicationExtension_Create.json
     */
    /**
     * Sample code: Puts the replication extension.
     * 
     * @param manager Entry point to RecoveryServicesDataReplicationManager.
     */
    public static void putsTheReplicationExtension(
        com.azure.resourcemanager.recoveryservicesdatareplication.RecoveryServicesDataReplicationManager manager) {
        manager.replicationExtensions().define("g16yjJ")
            .withExistingReplicationVault("rgrecoveryservicesdatareplication", "4").create();
    }
}
