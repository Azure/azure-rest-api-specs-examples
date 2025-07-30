
/**
 * Samples for Occurrences ListByScheduledAction.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-04-15-preview/Occurrences_ListByScheduledAction_MaximumSet_Gen.json
     */
    /**
     * Sample code: Occurrences_ListByScheduledAction_MaximumSet.
     * 
     * @param manager Entry point to ComputeScheduleManager.
     */
    public static void occurrencesListByScheduledActionMaximumSet(
        com.azure.resourcemanager.computeschedule.ComputeScheduleManager manager) {
        manager.occurrences().listByScheduledAction("rgcomputeschedule", "myScheduledAction",
            com.azure.core.util.Context.NONE);
    }
}
