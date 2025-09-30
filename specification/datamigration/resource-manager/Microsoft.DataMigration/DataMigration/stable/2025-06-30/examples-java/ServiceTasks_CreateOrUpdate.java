
import com.azure.resourcemanager.datamigration.fluent.models.ProjectTaskInner;
import com.azure.resourcemanager.datamigration.models.ConnectToSourceMySqlTaskInput;
import com.azure.resourcemanager.datamigration.models.ConnectToSourceMySqlTaskProperties;
import com.azure.resourcemanager.datamigration.models.MySqlConnectionInfo;

/**
 * Samples for ServiceTasks CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/datamigration/resource-manager/Microsoft.DataMigration/DataMigration/stable/2025-06-30/examples/
     * ServiceTasks_CreateOrUpdate.json
     */
    /**
     * Sample code: Tasks_CreateOrUpdate.
     * 
     * @param manager Entry point to DataMigrationManager.
     */
    public static void tasksCreateOrUpdate(com.azure.resourcemanager.datamigration.DataMigrationManager manager) {
        manager.serviceTasks().createOrUpdateWithResponse("DmsSdkRg", "DmsSdkService", "DmsSdkTask",
            new ProjectTaskInner()
                .withProperties(new ConnectToSourceMySqlTaskProperties().withInput(new ConnectToSourceMySqlTaskInput()
                    .withSourceConnectionInfo(new MySqlConnectionInfo().withServerName("localhost").withPort(3306)))),
            com.azure.core.util.Context.NONE);
    }
}
