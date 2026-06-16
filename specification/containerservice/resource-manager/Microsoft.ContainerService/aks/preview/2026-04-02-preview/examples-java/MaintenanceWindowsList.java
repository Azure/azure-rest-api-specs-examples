
/**
 * Samples for MaintenanceWindows ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-02-preview/MaintenanceWindowsList.json
     */
    /**
     * Sample code: List Maintenance Windows by Resource Group.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void listMaintenanceWindowsByResourceGroup(
        com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getMaintenanceWindows().listByResourceGroup("rg-maintenance",
            com.azure.core.util.Context.NONE);
    }
}
