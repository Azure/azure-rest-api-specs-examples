
/**
 * Samples for ConfigurationAssignmentsForResourceGroup Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/ConfigurationAssignmentsForResourceGroup_Delete.json
     */
    /**
     * Sample code: ConfigurationAssignmentsForResourceGroup_Delete.
     * 
     * @param manager Entry point to MaintenanceManager.
     */
    public static void configurationAssignmentsForResourceGroupDelete(
        com.azure.resourcemanager.maintenance.MaintenanceManager manager) {
        manager.configurationAssignmentsForResourceGroups().deleteByResourceGroupWithResponse("examplerg",
            "workervmConfiguration", com.azure.core.util.Context.NONE);
    }
}
