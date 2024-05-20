
/**
 * Samples for ConfigurationAssignmentsForSubscriptions Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/maintenance/resource-manager/Microsoft.Maintenance/preview/2023-10-01-preview/examples/
     * ConfigurationAssignmentsForSubscriptions_Delete.json
     */
    /**
     * Sample code: ConfigurationAssignmentsForSubscriptions_Delete.
     * 
     * @param manager Entry point to MaintenanceManager.
     */
    public static void configurationAssignmentsForSubscriptionsDelete(
        com.azure.resourcemanager.maintenance.MaintenanceManager manager) {
        manager.configurationAssignmentsForSubscriptions().deleteWithResponse("workervmConfiguration",
            com.azure.core.util.Context.NONE);
    }
}
