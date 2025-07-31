
/**
 * Samples for Event Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/Event_Get.json
     */
    /**
     * Sample code: Gets the event.
     * 
     * @param manager Entry point to RecoveryServicesDataReplicationManager.
     */
    public static void getsTheEvent(
        com.azure.resourcemanager.recoveryservicesdatareplication.RecoveryServicesDataReplicationManager manager) {
        manager.events().getWithResponse("rgswagger_2024-09-01", "4", "231CIG", com.azure.core.util.Context.NONE);
    }
}
