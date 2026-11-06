
/**
 * Samples for ActionGroups GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-10-01-preview/getActionGroup.json
     */
    /**
     * Sample code: Get an action group.
     * 
     * @param manager Entry point to MonitorManager.
     */
    public static void getAnActionGroup(com.azure.resourcemanager.monitor.MonitorManager manager) {
        manager.serviceClient().getActionGroups().getByResourceGroupWithResponse("Default-NotificationRules",
            "SampleActionGroup", com.azure.core.util.Context.NONE);
    }
}
