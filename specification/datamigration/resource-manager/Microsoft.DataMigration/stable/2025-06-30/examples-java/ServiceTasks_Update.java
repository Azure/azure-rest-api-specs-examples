
import com.azure.resourcemanager.datamigration.fluent.models.ProjectTaskInner;
import com.azure.resourcemanager.datamigration.models.ConnectToSourceMySqlTaskInput;
import com.azure.resourcemanager.datamigration.models.ConnectToSourceMySqlTaskProperties;
import com.azure.resourcemanager.datamigration.models.MySqlConnectionInfo;

/**
 * Samples for ServiceTasks Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/datamigration/resource-manager/Microsoft.DataMigration/stable/2025-06-30/examples/
     * ServiceTasks_Update.json
     */
    /**
     * Sample code: Tasks_Update.
     * 
     * @param manager Entry point to DataMigrationManager.
     */
    public static void tasksUpdate(com.azure.resourcemanager.datamigration.DataMigrationManager manager) {
        manager.serviceTasks().updateWithResponse("DmsSdkRg", "DmsSdkService", "DmsSdkTask",
            new ProjectTaskInner()
                .withProperties(new ConnectToSourceMySqlTaskProperties().withInput(new ConnectToSourceMySqlTaskInput()
                    .withSourceConnectionInfo(new MySqlConnectionInfo().withServerName("localhost").withPort(3306)))),
            com.azure.core.util.Context.NONE);
    }
}
