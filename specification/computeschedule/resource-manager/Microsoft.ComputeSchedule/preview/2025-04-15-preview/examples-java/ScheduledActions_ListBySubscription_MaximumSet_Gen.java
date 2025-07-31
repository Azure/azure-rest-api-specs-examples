
/**
 * Samples for ScheduledActions List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-04-15-preview/ScheduledActions_ListBySubscription_MaximumSet_Gen.json
     */
    /**
     * Sample code: ScheduledActions_ListBySubscription_MaximumSet.
     * 
     * @param manager Entry point to ComputeScheduleManager.
     */
    public static void scheduledActionsListBySubscriptionMaximumSet(
        com.azure.resourcemanager.computeschedule.ComputeScheduleManager manager) {
        manager.scheduledActions().list(com.azure.core.util.Context.NONE);
    }
}
