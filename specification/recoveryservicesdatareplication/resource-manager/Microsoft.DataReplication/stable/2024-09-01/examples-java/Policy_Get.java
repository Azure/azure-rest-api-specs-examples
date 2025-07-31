
/**
 * Samples for Policy Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/Policy_Get.json
     */
    /**
     * Sample code: Gets the policy.
     * 
     * @param manager Entry point to RecoveryServicesDataReplicationManager.
     */
    public static void getsThePolicy(
        com.azure.resourcemanager.recoveryservicesdatareplication.RecoveryServicesDataReplicationManager manager) {
        manager.policies().getWithResponse("rgrecoveryservicesdatareplication", "4", "wdqsacasc",
            com.azure.core.util.Context.NONE);
    }
}
