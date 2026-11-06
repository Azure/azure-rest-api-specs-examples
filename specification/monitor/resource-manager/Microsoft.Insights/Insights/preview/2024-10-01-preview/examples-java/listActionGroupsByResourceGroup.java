
/**
 * Samples for ActionGroups ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-10-01-preview/listActionGroupsByResourceGroup.json
     */
    /**
     * Sample code: List action groups at resource group level.
     * 
     * @param manager Entry point to MonitorManager.
     */
    public static void listActionGroupsAtResourceGroupLevel(com.azure.resourcemanager.monitor.MonitorManager manager) {
        manager.serviceClient().getActionGroups().listByResourceGroup("Default-NotificationRules",
            com.azure.core.util.Context.NONE);
    }
}
