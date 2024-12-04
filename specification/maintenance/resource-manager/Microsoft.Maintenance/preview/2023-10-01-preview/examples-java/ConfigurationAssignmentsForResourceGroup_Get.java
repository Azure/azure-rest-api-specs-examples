
/**
 * Samples for ConfigurationAssignmentsForResourceGroup GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/maintenance/resource-manager/Microsoft.Maintenance/preview/2023-10-01-preview/examples/
     * ConfigurationAssignmentsForResourceGroup_Get.json
     */
    /**
     * Sample code: ConfigurationAssignmentsForResourceGroup_Get.
     * 
     * @param manager Entry point to MaintenanceManager.
     */
    public static void
        configurationAssignmentsForResourceGroupGet(com.azure.resourcemanager.maintenance.MaintenanceManager manager) {
        manager.configurationAssignmentsForResourceGroups().getByResourceGroupWithResponse("examplerg",
            "workervmConfiguration", com.azure.core.util.Context.NONE);
    }
}
