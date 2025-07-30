
/**
 * Samples for ScheduledActions GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-04-15-preview/ScheduledActions_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: ScheduledActions_Get_MaximumSet.
     * 
     * @param manager Entry point to ComputeScheduleManager.
     */
    public static void
        scheduledActionsGetMaximumSet(com.azure.resourcemanager.computeschedule.ComputeScheduleManager manager) {
        manager.scheduledActions().getByResourceGroupWithResponse("rgcomputeschedule", "myScheduledAction",
            com.azure.core.util.Context.NONE);
    }
}
