
/**
 * Samples for MaintenanceWindows List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-02-preview/MaintenanceWindowsListBySubscription.json
     */
    /**
     * Sample code: List Maintenance Windows by Subscription.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void listMaintenanceWindowsBySubscription(
        com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getMaintenanceWindows().list(com.azure.core.util.Context.NONE);
    }
}
