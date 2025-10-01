
/**
 * Samples for Usages List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/datamigration/resource-manager/Microsoft.DataMigration/DataMigration/stable/2025-06-30/examples/
     * Usages_List.json
     */
    /**
     * Sample code: Services_Usages.
     * 
     * @param manager Entry point to DataMigrationManager.
     */
    public static void servicesUsages(com.azure.resourcemanager.datamigration.DataMigrationManager manager) {
        manager.usages().list("westus", com.azure.core.util.Context.NONE);
    }
}
