
/**
 * Samples for ScheduledActions ListResources.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-04-15-preview/ScheduledActions_ListResources_MaximumSet_Gen.json
     */
    /**
     * Sample code: ScheduledActions_ListResources_MaximumSet.
     * 
     * @param manager Entry point to ComputeScheduleManager.
     */
    public static void scheduledActionsListResourcesMaximumSet(
        com.azure.resourcemanager.computeschedule.ComputeScheduleManager manager) {
        manager.scheduledActions().listResources("rgcomputeschedule", "myScheduledAction",
            com.azure.core.util.Context.NONE);
    }
}
