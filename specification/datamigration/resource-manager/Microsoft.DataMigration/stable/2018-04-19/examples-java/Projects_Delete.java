/** Samples for Projects Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/datamigration/resource-manager/Microsoft.DataMigration/stable/2018-04-19/examples/Projects_Delete.json
     */
    /**
     * Sample code: Projects_Delete.
     *
     * @param manager Entry point to DataMigrationManager.
     */
    public static void projectsDelete(com.azure.resourcemanager.datamigration.DataMigrationManager manager) {
        manager
            .projects()
            .deleteWithResponse("DmsSdkRg", "DmsSdkService", "DmsSdkProject", null, com.azure.core.util.Context.NONE);
    }
}
