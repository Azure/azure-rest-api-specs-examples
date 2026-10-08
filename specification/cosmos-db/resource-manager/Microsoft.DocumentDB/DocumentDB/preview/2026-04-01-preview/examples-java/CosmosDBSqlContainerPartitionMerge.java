
import com.azure.resourcemanager.cosmos.models.MergeParameters;

/**
 * Samples for SqlResources ListSqlContainerPartitionMerge.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/CosmosDBSqlContainerPartitionMerge.json
     */
    /**
     * Sample code: CosmosDBSqlContainerPartitionMerge.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBSqlContainerPartitionMerge(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getSqlResources().listSqlContainerPartitionMerge("rgName", "ddb1", "databaseName",
            "containerName", new MergeParameters().withIsDryRun(false), com.azure.core.util.Context.NONE);
    }
}
