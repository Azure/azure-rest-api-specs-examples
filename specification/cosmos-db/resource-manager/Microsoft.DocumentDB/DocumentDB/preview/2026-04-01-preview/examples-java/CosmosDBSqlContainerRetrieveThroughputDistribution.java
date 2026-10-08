
import com.azure.resourcemanager.cosmos.models.PhysicalPartitionId;
import com.azure.resourcemanager.cosmos.models.RetrieveThroughputParameters;
import com.azure.resourcemanager.cosmos.models.RetrieveThroughputPropertiesResource;
import java.util.Arrays;

/**
 * Samples for SqlResources SqlContainerRetrieveThroughputDistribution.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/CosmosDBSqlContainerRetrieveThroughputDistribution.json
     */
    /**
     * Sample code: CosmosDBSqlContainerRetrieveThroughputDistribution.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void
        cosmosDBSqlContainerRetrieveThroughputDistribution(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getSqlResources().sqlContainerRetrieveThroughputDistribution("rg1", "ddb1",
            "databaseName", "containerName",
            new RetrieveThroughputParameters()
                .withResource(new RetrieveThroughputPropertiesResource().withPhysicalPartitionIds(
                    Arrays.asList(new PhysicalPartitionId().withId("0"), new PhysicalPartitionId().withId("1")))),
            com.azure.core.util.Context.NONE);
    }
}
