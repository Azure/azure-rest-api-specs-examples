
/**
 * Samples for MaintenanceConfigurations Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2025-10-01/examples/
     * MaintenanceConfigurationsGet_MaintenanceWindow.json
     */
    /**
     * Sample code: Get Maintenance Configuration Configured With Maintenance Window.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getMaintenanceConfigurationConfiguredWithMaintenanceWindow(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.kubernetesClusters().manager().serviceClient().getMaintenanceConfigurations().getWithResponse("rg1",
            "clustername1", "aksManagedNodeOSUpgradeSchedule", com.azure.core.util.Context.NONE);
    }
}
