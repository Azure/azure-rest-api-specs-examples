
/**
 * Samples for ConfigurationAssignmentsForSubscriptions Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2023-10-01-preview/ConfigurationAssignmentsForSubscriptions_Get.json
     */
    /**
     * Sample code: ConfigurationAssignments_GetParent.
     * 
     * @param manager Entry point to MaintenanceManager.
     */
    public static void
        configurationAssignmentsGetParent(com.azure.resourcemanager.maintenance.MaintenanceManager manager) {
        manager.configurationAssignmentsForSubscriptions().getWithResponse("workervmConfiguration",
            com.azure.core.util.Context.NONE);
    }
}
