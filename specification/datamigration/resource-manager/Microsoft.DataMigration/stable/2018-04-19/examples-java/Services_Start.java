/** Samples for Services Start. */
public final class Main {
    /*
     * x-ms-original-file: specification/datamigration/resource-manager/Microsoft.DataMigration/stable/2018-04-19/examples/Services_Start.json
     */
    /**
     * Sample code: Services_Start.
     *
     * @param manager Entry point to DataMigrationManager.
     */
    public static void servicesStart(com.azure.resourcemanager.datamigration.DataMigrationManager manager) {
        manager.services().start("DmsSdkRg", "DmsSdkService", com.azure.core.util.Context.NONE);
    }
}
