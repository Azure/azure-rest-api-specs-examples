/** Samples for Projects ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/datamigration/resource-manager/Microsoft.DataMigration/stable/2018-04-19/examples/Projects_List.json
     */
    /**
     * Sample code: Projects_ListByResourceGroup.
     *
     * @param manager Entry point to DataMigrationManager.
     */
    public static void projectsListByResourceGroup(
        com.azure.resourcemanager.datamigration.DataMigrationManager manager) {
        manager.projects().listByResourceGroup("DmsSdkRg", "DmsSdkService", com.azure.core.util.Context.NONE);
    }
}
