/** Samples for Services ListSkus. */
public final class Main {
    /*
     * x-ms-original-file: specification/datamigration/resource-manager/Microsoft.DataMigration/stable/2018-04-19/examples/Services_ListSkus.json
     */
    /**
     * Sample code: Services_ListSkus.
     *
     * @param manager Entry point to DataMigrationManager.
     */
    public static void servicesListSkus(com.azure.resourcemanager.datamigration.DataMigrationManager manager) {
        manager.services().listSkus("DmsSdkRg", "DmsSdkService", com.azure.core.util.Context.NONE);
    }
}
