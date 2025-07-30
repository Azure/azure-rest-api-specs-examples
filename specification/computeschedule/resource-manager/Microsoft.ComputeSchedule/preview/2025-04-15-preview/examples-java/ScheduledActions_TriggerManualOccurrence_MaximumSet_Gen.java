
/**
 * Samples for ScheduledActions TriggerManualOccurrence.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-04-15-preview/ScheduledActions_TriggerManualOccurrence_MaximumSet_Gen.json
     */
    /**
     * Sample code: ScheduledActions_TriggerManualOccurrence_MaximumSet.
     * 
     * @param manager Entry point to ComputeScheduleManager.
     */
    public static void scheduledActionsTriggerManualOccurrenceMaximumSet(
        com.azure.resourcemanager.computeschedule.ComputeScheduleManager manager) {
        manager.scheduledActions().triggerManualOccurrenceWithResponse("rgcomputeschedule", "myScheduledAction",
            com.azure.core.util.Context.NONE);
    }
}
