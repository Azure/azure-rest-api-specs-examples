
import com.azure.resourcemanager.containerservice.fluent.models.MaintenanceConfigurationInner;

/**
 * Samples for MaintenanceConfigurations CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-02-preview/MaintenanceConfigurationsCreate_LinkedMaintenanceWindow.json
     */
    /**
     * Sample code: Create a Linked Maintenance Configuration.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void createALinkedMaintenanceConfiguration(
        com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getMaintenanceConfigurations().createOrUpdateWithResponse("rg1", "clustername1",
            "aksManagedAutoUpgradeSchedule",
            new MaintenanceConfigurationInner().withMaintenanceWindowId(
                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ContainerService/maintenanceWindows/myMaintenanceWindow"),
            com.azure.core.util.Context.NONE);
    }
}
