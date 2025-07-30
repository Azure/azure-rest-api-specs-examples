
/**
 * Samples for RecoveryPoint List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/RecoveryPoints_List.json
     */
    /**
     * Sample code: Lists the recovery points.
     * 
     * @param manager Entry point to RecoveryServicesDataReplicationManager.
     */
    public static void listsTheRecoveryPoints(
        com.azure.resourcemanager.recoveryservicesdatareplication.RecoveryServicesDataReplicationManager manager) {
        manager.recoveryPoints().list("rgrecoveryservicesdatareplication", "4", "d", com.azure.core.util.Context.NONE);
    }
}
