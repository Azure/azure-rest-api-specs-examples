
import com.azure.resourcemanager.cosmos.models.ContinuousBackupRestoreLocation;

/**
 * Samples for GremlinResources RetrieveContinuousBackupInformation.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cosmos-db/resource-manager/Microsoft.DocumentDB/DocumentDB/stable/2025-10-15/examples/
     * CosmosDBGremlinGraphBackupInformation.json
     */
    /**
     * Sample code: CosmosDBGremlinGraphBackupInformation.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBGremlinGraphBackupInformation(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cosmosDBAccounts().manager().serviceClient().getGremlinResources().retrieveContinuousBackupInformation(
            "rgName", "ddb1", "databaseName", "graphName",
            new ContinuousBackupRestoreLocation().withLocation("North Europe"), com.azure.core.util.Context.NONE);
    }
}
