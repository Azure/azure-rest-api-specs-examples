
import com.azure.resourcemanager.computeschedule.models.DeadlineType;
import com.azure.resourcemanager.computeschedule.models.ExecutionParameters;
import com.azure.resourcemanager.computeschedule.models.Resources;
import com.azure.resourcemanager.computeschedule.models.Schedule;
import com.azure.resourcemanager.computeschedule.models.SubmitHibernateRequest;
import java.time.OffsetDateTime;
import java.util.Arrays;

/**
 * Samples for ScheduledActions VirtualMachinesSubmitHibernate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-04-15-preview/ScheduledActions_VirtualMachinesSubmitHibernate_MinimumSet_Gen.json
     */
    /**
     * Sample code: ScheduledActions_VirtualMachinesSubmitHibernate_MinimumSet.
     * 
     * @param manager Entry point to ComputeScheduleManager.
     */
    public static void scheduledActionsVirtualMachinesSubmitHibernateMinimumSet(
        com.azure.resourcemanager.computeschedule.ComputeScheduleManager manager) {
        manager.scheduledActions().virtualMachinesSubmitHibernateWithResponse("zuevcqpgdohzbjodhachtr",
            new SubmitHibernateRequest()
                .withSchedule(new Schedule().withDeadLine(OffsetDateTime.parse("2025-04-17T00:23:56.803Z"))
                    .withTimeZone("aigbjdnldtzkteqi").withDeadlineType(DeadlineType.UNKNOWN))
                .withExecutionParameters(new ExecutionParameters())
                .withResources(new Resources().withIds(Arrays.asList(
                    "/subscriptions/YourSubscriptionId/resourceGroups/YourResourceGroupName/providers/Microsoft.Compute/virtualMachines/testResource4")))
                .withCorrelationid("b211f086-4b91-4686-a453-2f5c012e4d80"),
            com.azure.core.util.Context.NONE);
    }
}
