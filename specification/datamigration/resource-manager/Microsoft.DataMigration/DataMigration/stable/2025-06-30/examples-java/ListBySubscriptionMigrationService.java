
/**
 * Samples for MigrationServices List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/datamigration/resource-manager/Microsoft.DataMigration/DataMigration/stable/2025-06-30/examples/
     * ListBySubscriptionMigrationService.json
     */
    /**
     * Sample code: Get Services in the Subscriptions.
     * 
     * @param manager Entry point to DataMigrationManager.
     */
    public static void
        getServicesInTheSubscriptions(com.azure.resourcemanager.datamigration.DataMigrationManager manager) {
        manager.migrationServices().list(com.azure.core.util.Context.NONE);
    }
}
