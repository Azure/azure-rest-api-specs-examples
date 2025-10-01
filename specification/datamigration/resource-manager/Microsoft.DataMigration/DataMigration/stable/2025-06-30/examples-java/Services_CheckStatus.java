
/**
 * Samples for Services CheckStatus.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/datamigration/resource-manager/Microsoft.DataMigration/DataMigration/stable/2025-06-30/examples/
     * Services_CheckStatus.json
     */
    /**
     * Sample code: Services_CheckStatus.
     * 
     * @param manager Entry point to DataMigrationManager.
     */
    public static void servicesCheckStatus(com.azure.resourcemanager.datamigration.DataMigrationManager manager) {
        manager.services().checkStatusWithResponse("DmsSdkRg", "DmsSdkService", com.azure.core.util.Context.NONE);
    }
}
