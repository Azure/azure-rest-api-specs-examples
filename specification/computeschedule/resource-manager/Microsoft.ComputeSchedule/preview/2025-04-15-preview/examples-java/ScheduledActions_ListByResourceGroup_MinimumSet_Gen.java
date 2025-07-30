
/**
 * Samples for ScheduledActions ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-04-15-preview/ScheduledActions_ListByResourceGroup_MinimumSet_Gen.json
     */
    /**
     * Sample code: ScheduledActions_ListByResourceGroup_MinimumSet.
     * 
     * @param manager Entry point to ComputeScheduleManager.
     */
    public static void scheduledActionsListByResourceGroupMinimumSet(
        com.azure.resourcemanager.computeschedule.ComputeScheduleManager manager) {
        manager.scheduledActions().listByResourceGroup("rgcomputeschedule", com.azure.core.util.Context.NONE);
    }
}
