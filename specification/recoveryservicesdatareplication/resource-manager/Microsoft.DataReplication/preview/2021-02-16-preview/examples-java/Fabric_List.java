/** Samples for Fabric ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesdatareplication/resource-manager/Microsoft.DataReplication/preview/2021-02-16-preview/examples/Fabric_List.json
     */
    /**
     * Sample code: Fabric_List.
     *
     * @param manager Entry point to RecoveryServicesDataReplicationManager.
     */
    public static void fabricList(
        com.azure.resourcemanager.recoveryservicesdatareplication.RecoveryServicesDataReplicationManager manager) {
        manager
            .fabrics()
            .listByResourceGroup("rgrecoveryservicesdatareplication", "mjzsxwwmtvd", com.azure.core.util.Context.NONE);
    }
}
