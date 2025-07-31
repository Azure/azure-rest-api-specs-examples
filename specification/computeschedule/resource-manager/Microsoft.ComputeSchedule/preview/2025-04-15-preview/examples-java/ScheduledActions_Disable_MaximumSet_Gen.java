
/**
 * Samples for ScheduledActions Disable.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-04-15-preview/ScheduledActions_Disable_MaximumSet_Gen.json
     */
    /**
     * Sample code: ScheduledActions_Disable_MaximumSet.
     * 
     * @param manager Entry point to ComputeScheduleManager.
     */
    public static void
        scheduledActionsDisableMaximumSet(com.azure.resourcemanager.computeschedule.ComputeScheduleManager manager) {
        manager.scheduledActions().disableWithResponse("rgcomputeschedule", "myScheduledAction",
            com.azure.core.util.Context.NONE);
    }
}
