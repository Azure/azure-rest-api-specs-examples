
import com.azure.resourcemanager.computeschedule.models.GetOperationStatusRequest;
import java.util.Arrays;

/**
 * Samples for ScheduledActions VirtualMachinesGetOperationStatus.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-04-15-preview/ScheduledActions_VirtualMachinesGetOperationStatus_MinimumSet_Gen.json
     */
    /**
     * Sample code: ScheduledActions_VirtualMachinesGetOperationStatus_MinimumSet.
     * 
     * @param manager Entry point to ComputeScheduleManager.
     */
    public static void scheduledActionsVirtualMachinesGetOperationStatusMinimumSet(
        com.azure.resourcemanager.computeschedule.ComputeScheduleManager manager) {
        manager.scheduledActions().virtualMachinesGetOperationStatusWithResponse("ykvvjfoopmkwznctgaiblzvea",
            new GetOperationStatusRequest().withOperationIds(Arrays.asList("duhqnwosjzexcfwfhryvy"))
                .withCorrelationid("b211f086-4b91-4686-a453-2f5c012e4d80"),
            com.azure.core.util.Context.NONE);
    }
}
