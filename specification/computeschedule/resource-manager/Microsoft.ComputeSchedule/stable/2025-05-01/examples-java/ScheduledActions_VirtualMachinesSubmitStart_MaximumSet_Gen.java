
import com.azure.resourcemanager.computeschedule.models.DeadlineType;
import com.azure.resourcemanager.computeschedule.models.ExecutionParameters;
import com.azure.resourcemanager.computeschedule.models.OptimizationPreference;
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
     * x-ms-original-file: 2025-05-01/ScheduledActions_VirtualMachinesSubmitStart_MaximumSet_Gen.json
     */
    /**
     * Sample code: ScheduledActions_VirtualMachinesSubmitStart_MaximumSet_Gen - generated by [MaximumSet] rule.
     * 
     * @param manager Entry point to ComputeScheduleManager.
     */
    public static void scheduledActionsVirtualMachinesSubmitStartMaximumSetGenGeneratedByMaximumSetRule(
        com.azure.resourcemanager.computeschedule.ComputeScheduleManager manager) {
        manager.scheduledActions().virtualMachinesSubmitStartWithResponse("bgyvpodcjmcmbxohvil",
            new SubmitStartRequest()
                .withSchedule(new Schedule().withDeadline(OffsetDateTime.parse("2025-04-15T19:47:04.403Z"))
                    .withDeadLine(OffsetDateTime.parse("2025-04-15T19:47:04.403Z"))
                    .withTimezone("qacufsmctpgjozovlsihrzoctatcsj").withTimeZone("upnmayfebiadztdktxzq")
                    .withDeadlineType(DeadlineType.UNKNOWN))
                .withExecutionParameters(
                    new ExecutionParameters().withOptimizationPreference(OptimizationPreference.COST)
                        .withRetryPolicy(new RetryPolicy().withRetryCount(25).withRetryWindowInMinutes(4)))
                .withResources(new Resources().withIds(Arrays.asList(
                    "/subscriptions/YourSubscriptionId/resourceGroups/YourResourceGroupName/providers/Microsoft.Compute/virtualMachines/testResource3")))
                .withCorrelationid("bvmpxvbd"),
            com.azure.core.util.Context.NONE);
    }
}
