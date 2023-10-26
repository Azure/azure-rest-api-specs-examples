/** Samples for ReplicationExtension Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesdatareplication/resource-manager/Microsoft.DataReplication/preview/2021-02-16-preview/examples/ReplicationExtension_Delete.json
     */
    /**
     * Sample code: ReplicationExtension_Delete.
     *
     * @param manager Entry point to RecoveryServicesDataReplicationManager.
     */
    public static void replicationExtensionDelete(
        com.azure.resourcemanager.recoveryservicesdatareplication.RecoveryServicesDataReplicationManager manager) {
        manager
            .replicationExtensions()
            .delete("rgrecoveryservicesdatareplication", "4", "g16yjJ", com.azure.core.util.Context.NONE);
    }
}
