
/**
 * Samples for ActionGroups List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-10-01-preview/listActionGroups.json
     */
    /**
     * Sample code: List action groups at subscription level.
     * 
     * @param manager Entry point to MonitorManager.
     */
    public static void listActionGroupsAtSubscriptionLevel(com.azure.resourcemanager.monitor.MonitorManager manager) {
        manager.serviceClient().getActionGroups().list(com.azure.core.util.Context.NONE);
    }
}
