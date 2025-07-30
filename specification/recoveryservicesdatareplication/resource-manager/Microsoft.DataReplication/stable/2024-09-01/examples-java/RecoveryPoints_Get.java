
/**
 * Samples for RecoveryPoint Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/RecoveryPoints_Get.json
     */
    /**
     * Sample code: Gets the recovery point.
     * 
     * @param manager Entry point to RecoveryServicesDataReplicationManager.
     */
    public static void getsTheRecoveryPoint(
        com.azure.resourcemanager.recoveryservicesdatareplication.RecoveryServicesDataReplicationManager manager) {
        manager.recoveryPoints().getWithResponse("rgrecoveryservicesdatareplication", "4", "d", "1X",
            com.azure.core.util.Context.NONE);
    }
}
