
/**
 * Samples for Event List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/Event_List.json
     */
    /**
     * Sample code: Lists the events.
     * 
     * @param manager Entry point to RecoveryServicesDataReplicationManager.
     */
    public static void listsTheEvents(
        com.azure.resourcemanager.recoveryservicesdatareplication.RecoveryServicesDataReplicationManager manager) {
        manager.events().list("rgswagger_2024-09-01", "4", null, "gabpzsxrifposvleqqcjnvofz", null,
            com.azure.core.util.Context.NONE);
    }
}
