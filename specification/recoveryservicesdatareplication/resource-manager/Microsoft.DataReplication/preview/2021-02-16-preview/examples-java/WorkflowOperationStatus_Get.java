/** Samples for WorkflowOperationStatus Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesdatareplication/resource-manager/Microsoft.DataReplication/preview/2021-02-16-preview/examples/WorkflowOperationStatus_Get.json
     */
    /**
     * Sample code: WorkflowOperationStatus_Get.
     *
     * @param manager Entry point to RecoveryServicesDataReplicationManager.
     */
    public static void workflowOperationStatusGet(
        com.azure.resourcemanager.recoveryservicesdatareplication.RecoveryServicesDataReplicationManager manager) {
        manager
            .workflowOperationStatus()
            .getWithResponse(
                "rgrecoveryservicesdatareplication", "4", "ZGH4y", "wdqcxc", com.azure.core.util.Context.NONE);
    }
}
