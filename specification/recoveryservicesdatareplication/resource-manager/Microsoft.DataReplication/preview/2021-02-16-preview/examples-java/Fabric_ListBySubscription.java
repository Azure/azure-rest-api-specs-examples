/** Samples for Fabric List. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesdatareplication/resource-manager/Microsoft.DataReplication/preview/2021-02-16-preview/examples/Fabric_ListBySubscription.json
     */
    /**
     * Sample code: Fabric_ListBySubscription.
     *
     * @param manager Entry point to RecoveryServicesDataReplicationManager.
     */
    public static void fabricListBySubscription(
        com.azure.resourcemanager.recoveryservicesdatareplication.RecoveryServicesDataReplicationManager manager) {
        manager.fabrics().list("rmgqrpzucsizbyjscxzockbiyg", com.azure.core.util.Context.NONE);
    }
}
