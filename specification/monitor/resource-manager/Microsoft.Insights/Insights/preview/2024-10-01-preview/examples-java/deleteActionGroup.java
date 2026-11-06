
/**
 * Samples for ActionGroups Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-10-01-preview/deleteActionGroup.json
     */
    /**
     * Sample code: Delete an action group.
     * 
     * @param manager Entry point to MonitorManager.
     */
    public static void deleteAnActionGroup(com.azure.resourcemanager.monitor.MonitorManager manager) {
        manager.serviceClient().getActionGroups().deleteWithResponse("Default-NotificationRules", "SampleActionGroup",
            com.azure.core.util.Context.NONE);
    }
}
