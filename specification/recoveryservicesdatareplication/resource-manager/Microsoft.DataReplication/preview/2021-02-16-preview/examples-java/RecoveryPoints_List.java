/** Samples for RecoveryPoints List. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesdatareplication/resource-manager/Microsoft.DataReplication/preview/2021-02-16-preview/examples/RecoveryPoints_List.json
     */
    /**
     * Sample code: RecoveryPoints_List.
     *
     * @param manager Entry point to RecoveryServicesDataReplicationManager.
     */
    public static void recoveryPointsList(
        com.azure.resourcemanager.recoveryservicesdatareplication.RecoveryServicesDataReplicationManager manager) {
        manager.recoveryPoints().list("rgrecoveryservicesdatareplication", "4", "d", com.azure.core.util.Context.NONE);
    }
}
