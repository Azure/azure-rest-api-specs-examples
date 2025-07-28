
/**
 * Samples for ResourceSkus List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/datamigration/resource-manager/Microsoft.DataMigration/preview/2025-03-15-preview/examples/
     * ResourceSkus_ListSkus.json
     */
    /**
     * Sample code: ListSkus.
     * 
     * @param manager Entry point to DataMigrationManager.
     */
    public static void listSkus(com.azure.resourcemanager.datamigration.DataMigrationManager manager) {
        manager.resourceSkus().list(com.azure.core.util.Context.NONE);
    }
}
