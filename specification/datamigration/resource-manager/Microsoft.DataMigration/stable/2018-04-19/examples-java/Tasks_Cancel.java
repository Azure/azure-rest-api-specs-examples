/** Samples for Tasks Cancel. */
public final class Main {
    /*
     * x-ms-original-file: specification/datamigration/resource-manager/Microsoft.DataMigration/stable/2018-04-19/examples/Tasks_Cancel.json
     */
    /**
     * Sample code: Tasks_Cancel.
     *
     * @param manager Entry point to DataMigrationManager.
     */
    public static void tasksCancel(com.azure.resourcemanager.datamigration.DataMigrationManager manager) {
        manager
            .tasks()
            .cancelWithResponse(
                "DmsSdkRg", "DmsSdkService", "DmsSdkProject", "DmsSdkTask", com.azure.core.util.Context.NONE);
    }
}
