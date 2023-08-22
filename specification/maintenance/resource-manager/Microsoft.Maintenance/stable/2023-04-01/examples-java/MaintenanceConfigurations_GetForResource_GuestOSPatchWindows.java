/** Samples for MaintenanceConfigurations GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/maintenance/resource-manager/Microsoft.Maintenance/stable/2023-04-01/examples/MaintenanceConfigurations_GetForResource_GuestOSPatchWindows.json
     */
    /**
     * Sample code: MaintenanceConfigurations_GetForResource_GuestOSPatchWindows.
     *
     * @param manager Entry point to MaintenanceManager.
     */
    public static void maintenanceConfigurationsGetForResourceGuestOSPatchWindows(
        com.azure.resourcemanager.maintenance.MaintenanceManager manager) {
        manager
            .maintenanceConfigurations()
            .getByResourceGroupWithResponse("examplerg", "configuration1", com.azure.core.util.Context.NONE);
    }
}
