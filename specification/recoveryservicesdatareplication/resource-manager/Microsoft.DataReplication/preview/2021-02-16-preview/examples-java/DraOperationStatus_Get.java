/** Samples for DraOperationStatus Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesdatareplication/resource-manager/Microsoft.DataReplication/preview/2021-02-16-preview/examples/DraOperationStatus_Get.json
     */
    /**
     * Sample code: DraOperationStatus_Get.
     *
     * @param manager Entry point to RecoveryServicesDataReplicationManager.
     */
    public static void draOperationStatusGet(
        com.azure.resourcemanager.recoveryservicesdatareplication.RecoveryServicesDataReplicationManager manager) {
        manager
            .draOperationStatus()
            .getWithResponse(
                "rgrecoveryservicesdatareplication", "wPR", "M", "dadsqwcq", com.azure.core.util.Context.NONE);
    }
}
