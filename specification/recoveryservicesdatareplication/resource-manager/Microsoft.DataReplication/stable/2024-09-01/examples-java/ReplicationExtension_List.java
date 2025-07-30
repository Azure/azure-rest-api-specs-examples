
/**
 * Samples for ReplicationExtension List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/ReplicationExtension_List.json
     */
    /**
     * Sample code: Lists the replication extensions.
     * 
     * @param manager Entry point to RecoveryServicesDataReplicationManager.
     */
    public static void listsTheReplicationExtensions(
        com.azure.resourcemanager.recoveryservicesdatareplication.RecoveryServicesDataReplicationManager manager) {
        manager.replicationExtensions().list("rgrecoveryservicesdatareplication", "4",
            com.azure.core.util.Context.NONE);
    }
}
