
import com.azure.resourcemanager.computeschedule.models.ExecuteStartRequest;
import com.azure.resourcemanager.computeschedule.models.ExecutionParameters;
import com.azure.resourcemanager.computeschedule.models.Resources;
import com.azure.resourcemanager.computeschedule.models.RetryPolicy;
import java.util.Arrays;

/**
 * Samples for ScheduledActions VirtualMachinesExecuteStart.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-10-01/ScheduledActions_VirtualMachinesExecuteStart.json
     */
    /**
     * Sample code: ScheduledActions_VirtualMachinesExecuteStart.
     * 
     * @param manager Entry point to ComputeScheduleManager.
     */
    public static void scheduledActionsVirtualMachinesExecuteStart(
        com.azure.resourcemanager.computeschedule.ComputeScheduleManager manager) {
        manager.scheduledActions().virtualMachinesExecuteStartWithResponse("eastus2euap", new ExecuteStartRequest()
            .withExecutionParameters(new ExecutionParameters()
                .withRetryPolicy(new RetryPolicy().withRetryCount(2).withRetryWindowInMinutes(27)))
            .withResources(new Resources().withIds(Arrays.asList(
                "/subscriptions/YourSubscriptionId/resourceGroups/YourResourceGroupName/providers/Microsoft.Compute/virtualMachines/testResource3")))
            .withCorrelationid("23480d2f-1dca-4610-afb4-dd25eec1f34r"), com.azure.core.util.Context.NONE);
    }
}
