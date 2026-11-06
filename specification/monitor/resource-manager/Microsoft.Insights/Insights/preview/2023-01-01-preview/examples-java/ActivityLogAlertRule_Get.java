
/**
 * Samples for ActivityLogAlerts GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2023-01-01-preview/ActivityLogAlertRule_Get.json
     */
    /**
     * Sample code: Get an Activity Log Alert rule.
     * 
     * @param manager Entry point to MonitorManager.
     */
    public static void getAnActivityLogAlertRule(com.azure.resourcemanager.monitor.MonitorManager manager) {
        manager.serviceClient().getActivityLogAlerts().getByResourceGroupWithResponse("MyResourceGroup",
            "SampleActivityLogAlertRule", com.azure.core.util.Context.NONE);
    }
}
