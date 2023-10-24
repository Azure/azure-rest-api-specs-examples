/** Samples for ReplicationExtension Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesdatareplication/resource-manager/Microsoft.DataReplication/preview/2021-02-16-preview/examples/ReplicationExtension_Get.json
     */
    /**
     * Sample code: ReplicationExtension_Get.
     *
     * @param manager Entry point to RecoveryServicesDataReplicationManager.
     */
    public static void replicationExtensionGet(
        com.azure.resourcemanager.recoveryservicesdatareplication.RecoveryServicesDataReplicationManager manager) {
        manager
            .replicationExtensions()
            .getWithResponse("rgrecoveryservicesdatareplication", "4", "g16yjJ", com.azure.core.util.Context.NONE);
    }
}
