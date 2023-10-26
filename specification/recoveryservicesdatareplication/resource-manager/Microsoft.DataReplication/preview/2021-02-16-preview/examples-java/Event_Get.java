/** Samples for Event Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesdatareplication/resource-manager/Microsoft.DataReplication/preview/2021-02-16-preview/examples/Event_Get.json
     */
    /**
     * Sample code: Event_Get.
     *
     * @param manager Entry point to RecoveryServicesDataReplicationManager.
     */
    public static void eventGet(
        com.azure.resourcemanager.recoveryservicesdatareplication.RecoveryServicesDataReplicationManager manager) {
        manager
            .events()
            .getWithResponse("rgrecoveryservicesdatareplication", "4", "231CIG", com.azure.core.util.Context.NONE);
    }
}
