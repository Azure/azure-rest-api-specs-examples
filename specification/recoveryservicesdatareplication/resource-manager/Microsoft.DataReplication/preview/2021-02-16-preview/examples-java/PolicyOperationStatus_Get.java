/** Samples for PolicyOperationStatus Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesdatareplication/resource-manager/Microsoft.DataReplication/preview/2021-02-16-preview/examples/PolicyOperationStatus_Get.json
     */
    /**
     * Sample code: PolicyOperationStatus_Get.
     *
     * @param manager Entry point to RecoveryServicesDataReplicationManager.
     */
    public static void policyOperationStatusGet(
        com.azure.resourcemanager.recoveryservicesdatareplication.RecoveryServicesDataReplicationManager manager) {
        manager
            .policyOperationStatus()
            .getWithResponse(
                "rgrecoveryservicesdatareplication", "4", "xczxcwec", "wdqfsdxv", com.azure.core.util.Context.NONE);
    }
}
