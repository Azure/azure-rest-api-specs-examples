
import com.azure.resourcemanager.computeschedule.fluent.models.ScheduledActionResourceInner;
import com.azure.resourcemanager.computeschedule.models.Language;
import com.azure.resourcemanager.computeschedule.models.NotificationProperties;
import com.azure.resourcemanager.computeschedule.models.NotificationType;
import com.azure.resourcemanager.computeschedule.models.ResourceAttachRequest;
import java.util.Arrays;

/**
 * Samples for ScheduledActions AttachResources.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-04-15-preview/ScheduledActions_AttachResources_MaximumSet_Gen.json
     */
    /**
     * Sample code: ScheduledActions_AttachResources_MaximumSet.
     * 
     * @param manager Entry point to ComputeScheduleManager.
     */
    public static void scheduledActionsAttachResourcesMaximumSet(
        com.azure.resourcemanager.computeschedule.ComputeScheduleManager manager) {
        manager.scheduledActions().attachResourcesWithResponse("rgcomputeschedule", "myScheduledAction",
            new ResourceAttachRequest().withResources(Arrays.asList(new ScheduledActionResourceInner().withResourceId(
                "/subscriptions/1d04e8f1-ee04-4056-b0b2-718f5bb45b04/resourceGroups/myRg/providers/Microsoft.Compute/virtualMachines/myVm")
                .withNotificationSettings(
                    Arrays.asList(new NotificationProperties().withDestination("wbhryycyolvnypjxzlawwvb")
                        .withType(NotificationType.EMAIL).withLanguage(Language.EN_US).withDisabled(true))))),
            com.azure.core.util.Context.NONE);
    }
}
