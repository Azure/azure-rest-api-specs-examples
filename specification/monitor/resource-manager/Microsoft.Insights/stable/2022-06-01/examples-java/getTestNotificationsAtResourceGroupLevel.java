
import com.azure.core.util.Context;

/** Samples for ActionGroups GetTestNotificationsAtResourceGroupLevel. */
public final class Main {
    /*
     * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2022-06-01/examples/
     * getTestNotificationsAtResourceGroupLevel.json
     */
    /**
     * Sample code: Get notification details at resource group level.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        getNotificationDetailsAtResourceGroupLevel(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.diagnosticSettings().manager().serviceClient().getActionGroups()
            .getTestNotificationsAtResourceGroupLevelWithResponse("Default-TestNotifications", "11000222191287",
                Context.NONE);
    }
}
