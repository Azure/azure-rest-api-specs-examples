
/**
 * Samples for AppServiceEnvironments TestUpgradeAvailableNotification.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/AppServiceEnvironments_TestUpgradeAvailableNotification.json
     */
    /**
     * Sample code: Send a test notification that an upgrade is available for this App Service Environment.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void sendATestNotificationThatAnUpgradeIsAvailableForThisAppServiceEnvironment(
        com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getAppServiceEnvironments().testUpgradeAvailableNotificationWithResponse("rg",
            "SampleHostingEnvironment", com.azure.core.util.Context.NONE);
    }
}
