/** Samples for Policy Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesdatareplication/resource-manager/Microsoft.DataReplication/preview/2021-02-16-preview/examples/Policy_Get.json
     */
    /**
     * Sample code: Policy_Get.
     *
     * @param manager Entry point to RecoveryServicesDataReplicationManager.
     */
    public static void policyGet(
        com.azure.resourcemanager.recoveryservicesdatareplication.RecoveryServicesDataReplicationManager manager) {
        manager
            .policies()
            .getWithResponse("rgrecoveryservicesdatareplication", "4", "wdqsacasc", com.azure.core.util.Context.NONE);
    }
}
