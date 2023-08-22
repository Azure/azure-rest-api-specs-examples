/** Samples for ConfigurationAssignmentsForResourceGroup GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/maintenance/resource-manager/Microsoft.Maintenance/stable/2023-04-01/examples/ConfigurationAssignmentsForResourceGroup_Get.json
     */
    /**
     * Sample code: ConfigurationAssignmentsForResourceGroup_Get.
     *
     * @param manager Entry point to MaintenanceManager.
     */
    public static void configurationAssignmentsForResourceGroupGet(
        com.azure.resourcemanager.maintenance.MaintenanceManager manager) {
        manager
            .configurationAssignmentsForResourceGroups()
            .getByResourceGroupWithResponse("examplerg", "workervmConfiguration", com.azure.core.util.Context.NONE);
    }
}
