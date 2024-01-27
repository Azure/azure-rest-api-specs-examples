
import com.azure.core.util.Context;
import com.azure.resourcemanager.monitor.models.EnableRequest;

/** Samples for ActionGroups EnableReceiver. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/monitor/resource-manager/Microsoft.Insights/stable/2022-06-01/examples/enableReceiver.json
     */
    /**
     * Sample code: Enable the receiver.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void enableTheReceiver(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.diagnosticSettings().manager().serviceClient().getActionGroups().enableReceiverWithResponse(
            "Default-NotificationRules", "SampleActionGroup", new EnableRequest().withReceiverName("John Doe's mobile"),
            Context.NONE);
    }
}
