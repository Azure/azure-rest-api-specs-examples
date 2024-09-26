
import com.azure.resourcemanager.computeschedule.models.GetOperationErrorsRequest;
import java.util.Arrays;

/**
 * Samples for ScheduledActions VirtualMachinesGetOperationErrors.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-08-15-preview/ScheduledActions_VirtualMachinesGetOperationErrors_MaximumSet_Gen.json
     */
    /**
     * Sample code: CS_ScheduledActions_VirtualMachinesGetOperationErrors_Max.
     * 
     * @param manager Entry point to ComputeScheduleManager.
     */
    public static void csScheduledActionsVirtualMachinesGetOperationErrorsMax(
        com.azure.resourcemanager.computeschedule.ComputeScheduleManager manager) {
        manager.scheduledActions().virtualMachinesGetOperationErrorsWithResponse("hfsa",
            new GetOperationErrorsRequest().withOperationIds(Arrays.asList("DE84A209-5715-43E7-BC76-3E208A9A323")),
            com.azure.core.util.Context.NONE);
    }
}
