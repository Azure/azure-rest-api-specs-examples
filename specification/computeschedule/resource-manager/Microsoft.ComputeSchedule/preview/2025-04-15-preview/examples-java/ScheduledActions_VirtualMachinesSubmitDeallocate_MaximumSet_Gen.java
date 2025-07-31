
import com.azure.resourcemanager.computeschedule.models.DeadlineType;
import com.azure.resourcemanager.computeschedule.models.ExecutionParameters;
import com.azure.resourcemanager.computeschedule.models.OptimizationPreference;
import com.azure.resourcemanager.computeschedule.models.Resources;
import com.azure.resourcemanager.computeschedule.models.RetryPolicy;
import com.azure.resourcemanager.computeschedule.models.Schedule;
import com.azure.resourcemanager.computeschedule.models.SubmitDeallocateRequest;
import java.time.OffsetDateTime;
import java.util.Arrays;

/**
 * Samples for ScheduledActions VirtualMachinesSubmitDeallocate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-04-15-preview/ScheduledActions_VirtualMachinesSubmitDeallocate_MaximumSet_Gen.json
     */
    /**
     * Sample code: ScheduledActions_VirtualMachinesSubmitDeallocate_MaximumSet.
     * 
     * @param manager Entry point to ComputeScheduleManager.
     */
    public static void scheduledActionsVirtualMachinesSubmitDeallocateMaximumSet(
        com.azure.resourcemanager.computeschedule.ComputeScheduleManager manager) {
        manager.scheduledActions().virtualMachinesSubmitDeallocateWithResponse("ycipx", new SubmitDeallocateRequest()
            .withSchedule(new Schedule().withDeadLine(OffsetDateTime.parse("2025-04-17T00:23:56.803Z"))
                .withTimeZone("aigbjdnldtzkteqi").withDeadlineType(DeadlineType.UNKNOWN))
            .withExecutionParameters(new ExecutionParameters().withOptimizationPreference(OptimizationPreference.COST)
                .withRetryPolicy(new RetryPolicy().withRetryCount(17).withRetryWindowInMinutes(29)))
            .withResources(new Resources().withIds(Arrays.asList(
                "/subscriptions/YourSubscriptionId/resourceGroups/YourResourceGroupName/providers/Microsoft.Compute/virtualMachines/testResource4")))
            .withCorrelationid("b211f086-4b91-4686-a453-2f5c012e4d80"), com.azure.core.util.Context.NONE);
    }
}
