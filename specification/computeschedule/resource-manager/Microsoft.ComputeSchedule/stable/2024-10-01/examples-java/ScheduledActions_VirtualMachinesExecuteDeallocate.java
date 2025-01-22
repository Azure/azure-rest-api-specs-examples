
import com.azure.resourcemanager.computeschedule.models.ExecuteDeallocateRequest;
import com.azure.resourcemanager.computeschedule.models.ExecutionParameters;
import com.azure.resourcemanager.computeschedule.models.Resources;
import com.azure.resourcemanager.computeschedule.models.RetryPolicy;
import java.util.Arrays;

/**
 * Samples for ScheduledActions VirtualMachinesExecuteDeallocate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-10-01/ScheduledActions_VirtualMachinesExecuteDeallocate.json
     */
    /**
     * Sample code: ScheduledActions_VirtualMachinesExecuteDeallocate.
     * 
     * @param manager Entry point to ComputeScheduleManager.
     */
    public static void scheduledActionsVirtualMachinesExecuteDeallocate(
        com.azure.resourcemanager.computeschedule.ComputeScheduleManager manager) {
        manager.scheduledActions().virtualMachinesExecuteDeallocateWithResponse("eastus2euap",
            new ExecuteDeallocateRequest()
                .withExecutionParameters(new ExecutionParameters()
                    .withRetryPolicy(new RetryPolicy().withRetryCount(4).withRetryWindowInMinutes(27)))
                .withResources(new Resources().withIds(Arrays.asList(
                    "/subscriptions/YourSubscriptionId/resourceGroups/YourResourceGroupName/providers/Microsoft.Compute/virtualMachines/testResource3")))
                .withCorrelationid("23480d2f-1dca-4610-afb4-dd25eec1f34r"),
            com.azure.core.util.Context.NONE);
    }
}
