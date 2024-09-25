
import com.azure.resourcemanager.computeschedule.models.GetOperationErrorsRequest;
import java.util.Arrays;

/**
 * Samples for ScheduledActions VirtualMachinesGetOperationErrors.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-08-15-preview/ScheduledActions_VirtualMachinesGetOperationErrors_MinimumSet_Gen.json
     */
    /**
     * Sample code: CS_ScheduledActions_VirtualMachinesGetOperationErrors_Min.
     * 
     * @param manager Entry point to ComputeScheduleManager.
     */
    public static void csScheduledActionsVirtualMachinesGetOperationErrorsMin(
        com.azure.resourcemanager.computeschedule.ComputeScheduleManager manager) {
        manager.scheduledActions().virtualMachinesGetOperationErrorsWithResponse("ggxoaxzxtdbi",
            new GetOperationErrorsRequest().withOperationIds(Arrays.asList("qeicik")),
            com.azure.core.util.Context.NONE);
    }
}
