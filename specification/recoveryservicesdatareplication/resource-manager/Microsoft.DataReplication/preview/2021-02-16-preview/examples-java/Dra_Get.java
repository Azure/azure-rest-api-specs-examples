/** Samples for Dra Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesdatareplication/resource-manager/Microsoft.DataReplication/preview/2021-02-16-preview/examples/Dra_Get.json
     */
    /**
     * Sample code: Dra_Get.
     *
     * @param manager Entry point to RecoveryServicesDataReplicationManager.
     */
    public static void draGet(
        com.azure.resourcemanager.recoveryservicesdatareplication.RecoveryServicesDataReplicationManager manager) {
        manager
            .dras()
            .getWithResponse("rgrecoveryservicesdatareplication", "wPR", "M", com.azure.core.util.Context.NONE);
    }
}
