
import com.azure.resourcemanager.computeschedule.models.ResourceDetachRequest;
import java.util.Arrays;

/**
 * Samples for ScheduledActions DetachResources.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-04-15-preview/ScheduledActions_DetachResources_MaximumSet_Gen.json
     */
    /**
     * Sample code: ScheduledActions_DetachResources_MaximumSet.
     * 
     * @param manager Entry point to ComputeScheduleManager.
     */
    public static void scheduledActionsDetachResourcesMaximumSet(
        com.azure.resourcemanager.computeschedule.ComputeScheduleManager manager) {
        manager.scheduledActions().detachResourcesWithResponse("rgcomputeschedule", "myScheduledAction",
            new ResourceDetachRequest().withResources(Arrays.asList(
                "/subscriptions/1d04e8f1-ee04-4056-b0b2-718f5bb45b04/resourceGroups/myRg/providers/Microsoft.Compute/virtualMachines/myVm")),
            com.azure.core.util.Context.NONE);
    }
}
