/** Samples for ProtectedItemOperationStatus Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesdatareplication/resource-manager/Microsoft.DataReplication/preview/2021-02-16-preview/examples/ProtectedItemOperationStatus_Get.json
     */
    /**
     * Sample code: ProtectedItemOperationStatus_Get.
     *
     * @param manager Entry point to RecoveryServicesDataReplicationManager.
     */
    public static void protectedItemOperationStatusGet(
        com.azure.resourcemanager.recoveryservicesdatareplication.RecoveryServicesDataReplicationManager manager) {
        manager
            .protectedItemOperationStatus()
            .getWithResponse(
                "rgrecoveryservicesdatareplication", "4", "d", "wdqacsc", com.azure.core.util.Context.NONE);
    }
}
