
import com.azure.resourcemanager.computeschedule.models.ExecuteHibernateRequest;
import com.azure.resourcemanager.computeschedule.models.ExecutionParameters;
import com.azure.resourcemanager.computeschedule.models.OptimizationPreference;
import com.azure.resourcemanager.computeschedule.models.Resources;
import com.azure.resourcemanager.computeschedule.models.RetryPolicy;
import java.util.Arrays;

/**
 * Samples for ScheduledActions VirtualMachinesExecuteHibernate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-04-15-preview/ScheduledActions_VirtualMachinesExecuteHibernate_MaximumSet_Gen.json
     */
    /**
     * Sample code: ScheduledActions_VirtualMachinesExecuteHibernate_MaximumSet.
     * 
     * @param manager Entry point to ComputeScheduleManager.
     */
    public static void scheduledActionsVirtualMachinesExecuteHibernateMaximumSet(
        com.azure.resourcemanager.computeschedule.ComputeScheduleManager manager) {
        manager.scheduledActions().virtualMachinesExecuteHibernateWithResponse("gztd", new ExecuteHibernateRequest()
            .withExecutionParameters(new ExecutionParameters().withOptimizationPreference(OptimizationPreference.COST)
                .withRetryPolicy(new RetryPolicy().withRetryCount(17).withRetryWindowInMinutes(29)))
            .withResources(new Resources().withIds(Arrays.asList(
                "/subscriptions/YourSubscriptionId/resourceGroups/YourResourceGroupName/providers/Microsoft.Compute/virtualMachines/testResource4")))
            .withCorrelationid("b211f086-4b91-4686-a453-2f5c012e4d80"), com.azure.core.util.Context.NONE);
    }
}
