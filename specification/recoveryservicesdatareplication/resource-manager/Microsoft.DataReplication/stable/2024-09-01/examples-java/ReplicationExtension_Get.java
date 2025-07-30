
/**
 * Samples for ReplicationExtension Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/ReplicationExtension_Get.json
     */
    /**
     * Sample code: Gets the replication extension.
     * 
     * @param manager Entry point to RecoveryServicesDataReplicationManager.
     */
    public static void getsTheReplicationExtension(
        com.azure.resourcemanager.recoveryservicesdatareplication.RecoveryServicesDataReplicationManager manager) {
        manager.replicationExtensions().getWithResponse("rgrecoveryservicesdatareplication", "4", "g16yjJ",
            com.azure.core.util.Context.NONE);
    }
}
