/** Samples for ReplicationExtension List. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesdatareplication/resource-manager/Microsoft.DataReplication/preview/2021-02-16-preview/examples/ReplicationExtension_List.json
     */
    /**
     * Sample code: ReplicationExtension_List.
     *
     * @param manager Entry point to RecoveryServicesDataReplicationManager.
     */
    public static void replicationExtensionList(
        com.azure.resourcemanager.recoveryservicesdatareplication.RecoveryServicesDataReplicationManager manager) {
        manager
            .replicationExtensions()
            .list("rgrecoveryservicesdatareplication", "4", com.azure.core.util.Context.NONE);
    }
}
