
/**
 * Samples for MaintenanceConfigurations GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/maintenance/resource-manager/Microsoft.Maintenance/preview/2023-10-01-preview/examples/
     * MaintenanceConfigurations_GetForResource_GuestOSPatchLinux.json
     */
    /**
     * Sample code: MaintenanceConfigurations_GetForResource_GuestOSPatchLinux.
     * 
     * @param manager Entry point to MaintenanceManager.
     */
    public static void maintenanceConfigurationsGetForResourceGuestOSPatchLinux(
        com.azure.resourcemanager.maintenance.MaintenanceManager manager) {
        manager.maintenanceConfigurations().getByResourceGroupWithResponse("examplerg", "configuration1",
            com.azure.core.util.Context.NONE);
    }
}
