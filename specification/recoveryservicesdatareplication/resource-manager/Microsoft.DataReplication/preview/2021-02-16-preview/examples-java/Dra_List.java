/** Samples for Dra List. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesdatareplication/resource-manager/Microsoft.DataReplication/preview/2021-02-16-preview/examples/Dra_List.json
     */
    /**
     * Sample code: Dra_List.
     *
     * @param manager Entry point to RecoveryServicesDataReplicationManager.
     */
    public static void draList(
        com.azure.resourcemanager.recoveryservicesdatareplication.RecoveryServicesDataReplicationManager manager) {
        manager.dras().list("rgrecoveryservicesdatareplication", "wPR", com.azure.core.util.Context.NONE);
    }
}
