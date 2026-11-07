
/**
 * Samples for ActionGroups GetTestNotificationsAtActionGroupResourceLevel.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-10-01-preview/getTestNotificationsAtActionGroupResourceLevel.json
     */
    /**
     * Sample code: Get notification details at resource group level.
     * 
     * @param manager Entry point to MonitorManager.
     */
    public static void
        getNotificationDetailsAtResourceGroupLevel(com.azure.resourcemanager.monitor.MonitorManager manager) {
        manager.serviceClient().getActionGroups().getTestNotificationsAtActionGroupResourceLevelWithResponse(
            "TestRgName", "TestAgName", "11000222191287", com.azure.core.util.Context.NONE);
    }
}
