
/**
 * Samples for Services List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/datamigration/resource-manager/Microsoft.DataMigration/stable/2025-06-30/examples/Services_List.
     * json
     */
    /**
     * Sample code: Services_List.
     * 
     * @param manager Entry point to DataMigrationManager.
     */
    public static void servicesList(com.azure.resourcemanager.datamigration.DataMigrationManager manager) {
        manager.services().list(com.azure.core.util.Context.NONE);
    }
}
