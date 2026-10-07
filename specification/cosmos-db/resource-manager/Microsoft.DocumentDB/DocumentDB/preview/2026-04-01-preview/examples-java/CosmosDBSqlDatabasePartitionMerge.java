
import com.azure.resourcemanager.cosmos.models.MergeParameters;

/**
 * Samples for SqlResources SqlDatabasePartitionMerge.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/CosmosDBSqlDatabasePartitionMerge.json
     */
    /**
     * Sample code: CosmosDBSqlDatabasePartitionMerge.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBSqlDatabasePartitionMerge(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getSqlResources().sqlDatabasePartitionMerge("rgName", "ddb1", "databaseName",
            new MergeParameters().withIsDryRun(false), com.azure.core.util.Context.NONE);
    }
}
