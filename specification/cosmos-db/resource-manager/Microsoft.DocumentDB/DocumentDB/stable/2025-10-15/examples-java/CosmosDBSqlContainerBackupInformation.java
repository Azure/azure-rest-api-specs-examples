
import com.azure.resourcemanager.cosmos.models.ContinuousBackupRestoreLocation;

/**
 * Samples for SqlResources RetrieveContinuousBackupInformation.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cosmos-db/resource-manager/Microsoft.DocumentDB/DocumentDB/stable/2025-10-15/examples/
     * CosmosDBSqlContainerBackupInformation.json
     */
    /**
     * Sample code: CosmosDBSqlContainerBackupInformation.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBSqlContainerBackupInformation(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cosmosDBAccounts().manager().serviceClient().getSqlResources().retrieveContinuousBackupInformation(
            "rgName", "ddb1", "databaseName", "containerName",
            new ContinuousBackupRestoreLocation().withLocation("North Europe"), com.azure.core.util.Context.NONE);
    }
}
