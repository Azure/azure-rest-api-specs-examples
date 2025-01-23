
import com.azure.resourcemanager.computeschedule.models.GetOperationErrorsRequest;
import java.util.Arrays;

/**
 * Samples for ScheduledActions VirtualMachinesGetOperationErrors.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-10-01/ScheduledActions_VirtualMachinesGetOperationErrors.json
     */
    /**
     * Sample code: ScheduledActions_VirtualMachinesGetOperationErrors.
     * 
     * @param manager Entry point to ComputeScheduleManager.
     */
    public static void scheduledActionsVirtualMachinesGetOperationErrors(
        com.azure.resourcemanager.computeschedule.ComputeScheduleManager manager) {
        manager.scheduledActions().virtualMachinesGetOperationErrorsWithResponse("eastus2euap",
            new GetOperationErrorsRequest().withOperationIds(Arrays.asList("23480d2f-1dca-4610-afb4-dd25eec1f34r")),
            com.azure.core.util.Context.NONE);
    }
}
