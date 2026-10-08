
import com.azure.resourcemanager.cosmos.models.MergeParameters;

/**
 * Samples for MongoDBResources MongoDBDatabasePartitionMerge.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/CosmosDBMongoDBDatabasePartitionMerge.json
     */
    /**
     * Sample code: CosmosDBMongoDBDatabasePartitionMerge.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBMongoDBDatabasePartitionMerge(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getMongoDBResources().mongoDBDatabasePartitionMerge("rgName", "ddb1", "databaseName",
            new MergeParameters().withIsDryRun(false), com.azure.core.util.Context.NONE);
    }
}
