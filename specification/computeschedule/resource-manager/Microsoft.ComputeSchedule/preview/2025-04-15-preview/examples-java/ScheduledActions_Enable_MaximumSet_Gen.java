
/**
 * Samples for ScheduledActions Enable.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-04-15-preview/ScheduledActions_Enable_MaximumSet_Gen.json
     */
    /**
     * Sample code: ScheduledActions_Enable_MaximumSet.
     * 
     * @param manager Entry point to ComputeScheduleManager.
     */
    public static void
        scheduledActionsEnableMaximumSet(com.azure.resourcemanager.computeschedule.ComputeScheduleManager manager) {
        manager.scheduledActions().enableWithResponse("rgcomputeschedule", "myScheduledAction",
            com.azure.core.util.Context.NONE);
    }
}
