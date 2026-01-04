
import com.azure.resourcemanager.cosmos.models.ContinuousBackupRestoreLocation;

/**
 * Samples for MongoDBResources RetrieveContinuousBackupInformation.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cosmos-db/resource-manager/Microsoft.DocumentDB/DocumentDB/stable/2025-10-15/examples/
     * CosmosDBMongoDBCollectionBackupInformation.json
     */
    /**
     * Sample code: CosmosDBMongoDBCollectionBackupInformation.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        cosmosDBMongoDBCollectionBackupInformation(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cosmosDBAccounts().manager().serviceClient().getMongoDBResources().retrieveContinuousBackupInformation(
            "rgName", "ddb1", "databaseName", "collectionName",
            new ContinuousBackupRestoreLocation().withLocation("North Europe"), com.azure.core.util.Context.NONE);
    }
}
