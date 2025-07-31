
/**
 * Samples for ScheduledActions Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-04-15-preview/ScheduledActions_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: ScheduledActions_Delete_MaximumSet.
     * 
     * @param manager Entry point to ComputeScheduleManager.
     */
    public static void
        scheduledActionsDeleteMaximumSet(com.azure.resourcemanager.computeschedule.ComputeScheduleManager manager) {
        manager.scheduledActions().delete("rgcomputeschedule", "myScheduledAction", com.azure.core.util.Context.NONE);
    }
}
