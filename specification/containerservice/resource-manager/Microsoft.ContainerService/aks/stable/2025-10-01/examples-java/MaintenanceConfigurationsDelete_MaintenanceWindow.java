
/**
 * Samples for MaintenanceConfigurations Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2025-10-01/examples/
     * MaintenanceConfigurationsDelete_MaintenanceWindow.json
     */
    /**
     * Sample code: Delete Maintenance Configuration For Node OS Upgrade.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        deleteMaintenanceConfigurationForNodeOSUpgrade(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.kubernetesClusters().manager().serviceClient().getMaintenanceConfigurations().deleteWithResponse("rg1",
            "clustername1", "aksManagedNodeOSUpgradeSchedule", com.azure.core.util.Context.NONE);
    }
}
