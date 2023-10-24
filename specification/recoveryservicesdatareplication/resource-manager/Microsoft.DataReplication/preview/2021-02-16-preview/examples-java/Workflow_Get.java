/** Samples for Workflow Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesdatareplication/resource-manager/Microsoft.DataReplication/preview/2021-02-16-preview/examples/Workflow_Get.json
     */
    /**
     * Sample code: Workflow_Get.
     *
     * @param manager Entry point to RecoveryServicesDataReplicationManager.
     */
    public static void workflowGet(
        com.azure.resourcemanager.recoveryservicesdatareplication.RecoveryServicesDataReplicationManager manager) {
        manager
            .workflows()
            .getWithResponse("rgrecoveryservicesdatareplication", "4", "ZGH4y", com.azure.core.util.Context.NONE);
    }
}
