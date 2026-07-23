
import com.azure.resourcemanager.cosmos.models.ContinuousBackupRestoreLocation;

/**
 * Samples for MongoDBResources RetrieveContinuousBackupInformation.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBMongoDBCollectionBackupInformation.json
     */
    /**
     * Sample code: CosmosDBMongoDBCollectionBackupInformation.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void
        cosmosDBMongoDBCollectionBackupInformation(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getMongoDBResources().retrieveContinuousBackupInformation("rgName", "ddb1",
            "databaseName", "collectionName", new ContinuousBackupRestoreLocation().withLocation("North Europe"),
            com.azure.core.util.Context.NONE);
    }
}
