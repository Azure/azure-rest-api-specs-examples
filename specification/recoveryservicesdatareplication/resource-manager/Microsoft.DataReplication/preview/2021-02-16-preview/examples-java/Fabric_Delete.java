/** Samples for Fabric Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesdatareplication/resource-manager/Microsoft.DataReplication/preview/2021-02-16-preview/examples/Fabric_Delete.json
     */
    /**
     * Sample code: Fabric_Delete.
     *
     * @param manager Entry point to RecoveryServicesDataReplicationManager.
     */
    public static void fabricDelete(
        com.azure.resourcemanager.recoveryservicesdatareplication.RecoveryServicesDataReplicationManager manager) {
        manager.fabrics().delete("rgrecoveryservicesdatareplication", "wPR", com.azure.core.util.Context.NONE);
    }
}
