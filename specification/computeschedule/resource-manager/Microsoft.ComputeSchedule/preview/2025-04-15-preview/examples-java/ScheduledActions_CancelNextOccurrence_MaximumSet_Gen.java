
import com.azure.resourcemanager.computeschedule.models.CancelOccurrenceRequest;
import java.util.Arrays;

/**
 * Samples for ScheduledActions CancelNextOccurrence.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-04-15-preview/ScheduledActions_CancelNextOccurrence_MaximumSet_Gen.json
     */
    /**
     * Sample code: ScheduledActions_CancelNextOccurrence_MaximumSet.
     * 
     * @param manager Entry point to ComputeScheduleManager.
     */
    public static void scheduledActionsCancelNextOccurrenceMaximumSet(
        com.azure.resourcemanager.computeschedule.ComputeScheduleManager manager) {
        manager.scheduledActions().cancelNextOccurrenceWithResponse("rgcomputeschedule", "myScheduledAction",
            new CancelOccurrenceRequest().withResourceIds(Arrays.asList(
                "/subscriptions/1d04e8f1-ee04-4056-b0b2-718f5bb45b04/resourceGroups/myRg/providers/Microsoft.Compute/virtualMachines/myVm")),
            com.azure.core.util.Context.NONE);
    }
}
