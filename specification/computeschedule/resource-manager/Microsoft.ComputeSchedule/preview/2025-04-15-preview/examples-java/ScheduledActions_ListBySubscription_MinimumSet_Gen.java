
/**
 * Samples for ScheduledActions List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-04-15-preview/ScheduledActions_ListBySubscription_MinimumSet_Gen.json
     */
    /**
     * Sample code: ScheduledActions_ListBySubscription_MinimumSet.
     * 
     * @param manager Entry point to ComputeScheduleManager.
     */
    public static void scheduledActionsListBySubscriptionMinimumSet(
        com.azure.resourcemanager.computeschedule.ComputeScheduleManager manager) {
        manager.scheduledActions().list(com.azure.core.util.Context.NONE);
    }
}
