/** Samples for FabricOperationsStatus Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesdatareplication/resource-manager/Microsoft.DataReplication/preview/2021-02-16-preview/examples/FabricOperationsStatus_Get.json
     */
    /**
     * Sample code: FabricOperationsStatus_Get.
     *
     * @param manager Entry point to RecoveryServicesDataReplicationManager.
     */
    public static void fabricOperationsStatusGet(
        com.azure.resourcemanager.recoveryservicesdatareplication.RecoveryServicesDataReplicationManager manager) {
        manager
            .fabricOperationsStatus()
            .getWithResponse("rgrecoveryservicesdatareplication", "wPR", "dvfwerv", com.azure.core.util.Context.NONE);
    }
}
