
/**
 * Samples for ActionGroups GetTestNotificationsAtActionGroupResourceLevel.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2022-06-01/examples/
     * getTestNotificationsAtActionGroupResourceLevel.json
     */
    /**
     * Sample code: Get notification details at resource group level.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        getNotificationDetailsAtResourceGroupLevel(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.diagnosticSettings().manager().serviceClient().getActionGroups()
            .getTestNotificationsAtActionGroupResourceLevelWithResponse("TestRgName", "TestAgName", "11000222191287",
                com.azure.core.util.Context.NONE);
    }
}
