
import com.azure.resourcemanager.computeschedule.models.GetOperationStatusRequest;
import java.util.Arrays;

/**
 * Samples for ScheduledActions VirtualMachinesGetOperationStatus.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-10-01/ScheduledActions_VirtualMachinesGetOperationStatus.json
     */
    /**
     * Sample code: ScheduledActions_VirtualMachinesGetOperationStatus.
     * 
     * @param manager Entry point to ComputeScheduleManager.
     */
    public static void scheduledActionsVirtualMachinesGetOperationStatus(
        com.azure.resourcemanager.computeschedule.ComputeScheduleManager manager) {
        manager.scheduledActions().virtualMachinesGetOperationStatusWithResponse("eastus2euap",
            new GetOperationStatusRequest().withOperationIds(Arrays.asList("23480d2f-1dca-4610-afb4-dd25eec1f34r"))
                .withCorrelationid("35780d2f-1dca-4610-afb4-dd25eec1f34r"),
            com.azure.core.util.Context.NONE);
    }
}
