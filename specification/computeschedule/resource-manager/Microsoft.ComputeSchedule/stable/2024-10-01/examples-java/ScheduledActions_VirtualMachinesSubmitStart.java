
import com.azure.resourcemanager.computeschedule.models.DeadlineType;
import com.azure.resourcemanager.computeschedule.models.ExecutionParameters;
import com.azure.resourcemanager.computeschedule.models.Resources;
import com.azure.resourcemanager.computeschedule.models.RetryPolicy;
import com.azure.resourcemanager.computeschedule.models.Schedule;
import com.azure.resourcemanager.computeschedule.models.SubmitStartRequest;
import java.time.OffsetDateTime;
import java.util.Arrays;

/**
 * Samples for ScheduledActions VirtualMachinesSubmitStart.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-10-01/ScheduledActions_VirtualMachinesSubmitStart.json
     */
    /**
     * Sample code: ScheduledActions_VirtualMachinesSubmitStart.
     * 
     * @param manager Entry point to ComputeScheduleManager.
     */
    public static void scheduledActionsVirtualMachinesSubmitStart(
        com.azure.resourcemanager.computeschedule.ComputeScheduleManager manager) {
        manager.scheduledActions().virtualMachinesSubmitStartWithResponse("eastus2euap", new SubmitStartRequest()
            .withSchedule(new Schedule().withDeadline(OffsetDateTime.parse("2024-11-01T17:52:54.215Z"))
                .withTimezone("UTC").withDeadlineType(DeadlineType.INITIATE_AT))
            .withExecutionParameters(new ExecutionParameters()
                .withRetryPolicy(new RetryPolicy().withRetryCount(5).withRetryWindowInMinutes(27)))
            .withResources(new Resources().withIds(Arrays.asList(
                "/subscriptions/YourSubscriptionId/resourceGroups/YourResourceGroupName/providers/Microsoft.Compute/virtualMachines/testResource3")))
            .withCorrelationid("23480d2f-1dca-4610-afb4-dd25eec1f34r"), com.azure.core.util.Context.NONE);
    }
}
