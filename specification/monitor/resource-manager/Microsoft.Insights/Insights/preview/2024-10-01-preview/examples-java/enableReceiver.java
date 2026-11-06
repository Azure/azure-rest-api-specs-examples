
import com.azure.resourcemanager.monitor.models.EnableRequest;

/**
 * Samples for ActionGroups EnableReceiver.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-10-01-preview/enableReceiver.json
     */
    /**
     * Sample code: Enable the receiver.
     * 
     * @param manager Entry point to MonitorManager.
     */
    public static void enableTheReceiver(com.azure.resourcemanager.monitor.MonitorManager manager) {
        manager.serviceClient().getActionGroups().enableReceiverWithResponse("Default-NotificationRules",
            "SampleActionGroup", new EnableRequest().withReceiverName("John Doe's mobile"),
            com.azure.core.util.Context.NONE);
    }
}
