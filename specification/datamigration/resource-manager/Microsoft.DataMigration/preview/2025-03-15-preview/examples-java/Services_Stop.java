
/**
 * Samples for Services Stop.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/datamigration/resource-manager/Microsoft.DataMigration/preview/2025-03-15-preview/examples/
     * Services_Stop.json
     */
    /**
     * Sample code: Services_Stop.
     * 
     * @param manager Entry point to DataMigrationManager.
     */
    public static void servicesStop(com.azure.resourcemanager.datamigration.DataMigrationManager manager) {
        manager.services().stop("DmsSdkRg", "DmsSdkService", com.azure.core.util.Context.NONE);
    }
}
