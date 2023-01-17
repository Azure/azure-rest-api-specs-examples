/** Samples for Tasks Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/datamigration/resource-manager/Microsoft.DataMigration/stable/2018-04-19/examples/Tasks_Get.json
     */
    /**
     * Sample code: Tasks_Get.
     *
     * @param manager Entry point to DataMigrationManager.
     */
    public static void tasksGet(com.azure.resourcemanager.datamigration.DataMigrationManager manager) {
        manager
            .tasks()
            .getWithResponse(
                "DmsSdkRg", "DmsSdkService", "DmsSdkProject", "DmsSdkTask", null, com.azure.core.util.Context.NONE);
    }
}
