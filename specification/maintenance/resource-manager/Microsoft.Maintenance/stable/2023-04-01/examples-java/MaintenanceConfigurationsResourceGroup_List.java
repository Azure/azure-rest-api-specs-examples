/** Samples for MaintenanceConfigurationsForResourceGroup ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/maintenance/resource-manager/Microsoft.Maintenance/stable/2023-04-01/examples/MaintenanceConfigurationsResourceGroup_List.json
     */
    /**
     * Sample code: MaintenanceConfigurationsResourceGroup_List.
     *
     * @param manager Entry point to MaintenanceManager.
     */
    public static void maintenanceConfigurationsResourceGroupList(
        com.azure.resourcemanager.maintenance.MaintenanceManager manager) {
        manager
            .maintenanceConfigurationsForResourceGroups()
            .listByResourceGroup("examplerg", com.azure.core.util.Context.NONE);
    }
}
