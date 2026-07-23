
import com.azure.resourcemanager.cosmos.models.ContinuousBackupRestoreLocation;

/**
 * Samples for SqlResources RetrieveContinuousBackupInformation.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBSqlContainerBackupInformation.json
     */
    /**
     * Sample code: CosmosDBSqlContainerBackupInformation.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBSqlContainerBackupInformation(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getSqlResources().retrieveContinuousBackupInformation("rgName", "ddb1", "databaseName",
            "containerName", new ContinuousBackupRestoreLocation().withLocation("North Europe"),
            com.azure.core.util.Context.NONE);
    }
}
