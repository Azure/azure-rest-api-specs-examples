
import com.azure.resourcemanager.computeschedule.models.CancelOccurrenceRequest;
import java.util.Arrays;

/**
 * Samples for Occurrences Cancel.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-04-15-preview/Occurrences_Cancel_MaximumSet_Gen.json
     */
    /**
     * Sample code: Occurrences_Cancel_MaximumSet.
     * 
     * @param manager Entry point to ComputeScheduleManager.
     */
    public static void
        occurrencesCancelMaximumSet(com.azure.resourcemanager.computeschedule.ComputeScheduleManager manager) {
        manager.occurrences().cancelWithResponse("rgcomputeschedule", "myScheduledAction",
            "CB26D7CB-3E27-465F-99C8-EAF7A4118245",
            new CancelOccurrenceRequest().withResourceIds(Arrays.asList(
                "/subscriptions/CB26D7CB-3E27-465F-99C8-EAF7A4118245/resourceGroups/myRg/providers/Microsoft.Compute/virtualMachines/myVm")),
            com.azure.core.util.Context.NONE);
    }
}
