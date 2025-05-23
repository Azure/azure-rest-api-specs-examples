
import com.azure.resourcemanager.datamigration.models.AuthenticationType;
import com.azure.resourcemanager.datamigration.models.ConnectToTargetSqlDbTaskInput;
import com.azure.resourcemanager.datamigration.models.ConnectToTargetSqlDbTaskProperties;
import com.azure.resourcemanager.datamigration.models.ProjectTask;
import com.azure.resourcemanager.datamigration.models.SqlConnectionInfo;

/**
 * Samples for Tasks Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/datamigration/resource-manager/Microsoft.DataMigration/stable/2018-04-19/examples/Tasks_Update.json
     */
    /**
     * Sample code: Tasks_Update.
     * 
     * @param manager Entry point to DataMigrationManager.
     */
    public static void tasksUpdate(com.azure.resourcemanager.datamigration.DataMigrationManager manager) {
        ProjectTask resource = manager.tasks().getWithResponse("DmsSdkRg", "DmsSdkService", "DmsSdkProject",
            "DmsSdkTask", null, com.azure.core.util.Context.NONE).getValue();
        resource.update()
            .withProperties(new ConnectToTargetSqlDbTaskProperties().withInput(new ConnectToTargetSqlDbTaskInput()
                .withTargetConnectionInfo(new SqlConnectionInfo().withUsername("testuser")
                    .withPassword("fakeTokenPlaceholder").withDataSource("ssma-test-server.database.windows.net")
                    .withAuthentication(AuthenticationType.SQL_AUTHENTICATION).withEncryptConnection(true)
                    .withTrustServerCertificate(true))))
            .apply();
    }
}
