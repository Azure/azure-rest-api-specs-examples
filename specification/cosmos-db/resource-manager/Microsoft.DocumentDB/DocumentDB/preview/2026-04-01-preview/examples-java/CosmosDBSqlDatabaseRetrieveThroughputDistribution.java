
import com.azure.resourcemanager.cosmos.models.PhysicalPartitionId;
import com.azure.resourcemanager.cosmos.models.RetrieveThroughputParameters;
import com.azure.resourcemanager.cosmos.models.RetrieveThroughputPropertiesResource;
import java.util.Arrays;

/**
 * Samples for SqlResources SqlDatabaseRetrieveThroughputDistribution.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/CosmosDBSqlDatabaseRetrieveThroughputDistribution.json
     */
    /**
     * Sample code: CosmosDBSqlDatabaseRetrieveThroughputDistribution.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void
        cosmosDBSqlDatabaseRetrieveThroughputDistribution(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getSqlResources().sqlDatabaseRetrieveThroughputDistribution("rg1", "ddb1",
            "databaseName",
            new RetrieveThroughputParameters()
                .withResource(new RetrieveThroughputPropertiesResource().withPhysicalPartitionIds(
                    Arrays.asList(new PhysicalPartitionId().withId("0"), new PhysicalPartitionId().withId("1")))),
            com.azure.core.util.Context.NONE);
    }
}
