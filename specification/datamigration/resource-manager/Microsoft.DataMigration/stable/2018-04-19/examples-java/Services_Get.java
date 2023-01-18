/** Samples for Services GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/datamigration/resource-manager/Microsoft.DataMigration/stable/2018-04-19/examples/Services_Get.json
     */
    /**
     * Sample code: Services_CreateOrUpdate.
     *
     * @param manager Entry point to DataMigrationManager.
     */
    public static void servicesCreateOrUpdate(com.azure.resourcemanager.datamigration.DataMigrationManager manager) {
        manager
            .services()
            .getByResourceGroupWithResponse("DmsSdkRg", "DmsSdkService", com.azure.core.util.Context.NONE);
    }
}
