
import com.azure.resourcemanager.computeschedule.models.CancelOperationsRequest;
import java.util.Arrays;

/**
 * Samples for ScheduledActions VirtualMachinesCancelOperations.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-04-15-preview/ScheduledActions_VirtualMachinesCancelOperations_MinimumSet_Gen.json
     */
    /**
     * Sample code: ScheduledActions_VirtualMachinesCancelOperations_MinimumSet.
     * 
     * @param manager Entry point to ComputeScheduleManager.
     */
    public static void scheduledActionsVirtualMachinesCancelOperationsMinimumSet(
        com.azure.resourcemanager.computeschedule.ComputeScheduleManager manager) {
        manager.scheduledActions().virtualMachinesCancelOperationsWithResponse("fhdunfivmjiltaiakwhhwdgemfcld",
            new CancelOperationsRequest().withOperationIds(Arrays.asList("b211f086-4b91-4686-a453-2f5c012e4d80"))
                .withCorrelationid("b211f086-4b91-4686-a453-2f5c012e4d80"),
            com.azure.core.util.Context.NONE);
    }
}
