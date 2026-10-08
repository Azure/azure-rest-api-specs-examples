
import com.azure.resourcemanager.cosmos.models.PhysicalPartitionId;
import com.azure.resourcemanager.cosmos.models.RetrieveThroughputParameters;
import com.azure.resourcemanager.cosmos.models.RetrieveThroughputPropertiesResource;
import java.util.Arrays;

/**
 * Samples for MongoDBResources MongoDBContainerRetrieveThroughputDistribution.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/CosmosDBMongoDBCollectionRetrieveThroughputDistribution.json
     */
    /**
     * Sample code: CosmosDBMongoDBCollectionRetrieveThroughputDistribution.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBMongoDBCollectionRetrieveThroughputDistribution(
        com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getMongoDBResources().mongoDBContainerRetrieveThroughputDistribution("rg1", "ddb1",
            "databaseName", "collectionName",
            new RetrieveThroughputParameters()
                .withResource(new RetrieveThroughputPropertiesResource().withPhysicalPartitionIds(
                    Arrays.asList(new PhysicalPartitionId().withId("0"), new PhysicalPartitionId().withId("1")))),
            com.azure.core.util.Context.NONE);
    }
}
