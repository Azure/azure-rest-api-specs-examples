
import com.azure.resourcemanager.cosmos.models.MergeParameters;

/**
 * Samples for MongoDBResources ListMongoDBCollectionPartitionMerge.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/CosmosDBMongoDBCollectionPartitionMerge.json
     */
    /**
     * Sample code: CosmosDBMongoDBCollectionPartitionMerge.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBMongoDBCollectionPartitionMerge(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getMongoDBResources().listMongoDBCollectionPartitionMerge("rgName", "ddb1",
            "databaseName", "collectionName", new MergeParameters().withIsDryRun(false),
            com.azure.core.util.Context.NONE);
    }
}
