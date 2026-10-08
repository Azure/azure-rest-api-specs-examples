
import com.azure.resourcemanager.cosmos.models.ContainerPartitionKey;
import com.azure.resourcemanager.cosmos.models.CreateUpdateOptions;
import com.azure.resourcemanager.cosmos.models.DataType;
import com.azure.resourcemanager.cosmos.models.IncludedPath;
import com.azure.resourcemanager.cosmos.models.IndexKind;
import com.azure.resourcemanager.cosmos.models.Indexes;
import com.azure.resourcemanager.cosmos.models.IndexingMode;
import com.azure.resourcemanager.cosmos.models.IndexingPolicy;
import com.azure.resourcemanager.cosmos.models.MaterializedViewDefinition;
import com.azure.resourcemanager.cosmos.models.PartitionKind;
import com.azure.resourcemanager.cosmos.models.SqlContainerCreateUpdateParameters;
import com.azure.resourcemanager.cosmos.models.SqlContainerResource;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for SqlResources CreateUpdateSqlContainer.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/CosmosDBSqlMaterializedViewCreateUpdate.json
     */
    /**
     * Sample code: CosmosDBSqlMaterializedViewCreateUpdate.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBSqlMaterializedViewCreateUpdate(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getSqlResources().createUpdateSqlContainer("rg1", "ddb1", "databaseName",
            "mvContainerName",
            new SqlContainerCreateUpdateParameters().withLocation("West US").withTags(mapOf()).withResource(
                new SqlContainerResource().withId("mvContainerName").withIndexingPolicy(new IndexingPolicy()
                    .withAutomatic(true).withIndexingMode(IndexingMode.CONSISTENT)
                    .withIncludedPaths(Arrays.asList(new IncludedPath().withPath("/*")
                        .withIndexes(Arrays.asList(
                            new Indexes().withDataType(DataType.STRING).withPrecision(-1).withKind(IndexKind.RANGE),
                            new Indexes().withDataType(DataType.NUMBER).withPrecision(-1).withKind(IndexKind.RANGE)))))
                    .withExcludedPaths(Arrays.asList()))
                    .withPartitionKey(
                        new ContainerPartitionKey().withPaths(Arrays.asList("/mvpk")).withKind(PartitionKind.HASH))
                    .withMaterializedViewDefinition(
                        new MaterializedViewDefinition().withSourceCollectionId("sourceContainerName")
                            .withDefinition("select * from ROOT").withThroughputBucketForBuild(1)))
                .withOptions(new CreateUpdateOptions()),
            com.azure.core.util.Context.NONE);
    }

    // Use "Map.of" if available
    @SuppressWarnings("unchecked")
    private static <T> Map<String, T> mapOf(Object... inputs) {
        Map<String, T> map = new HashMap<>();
        for (int i = 0; i < inputs.length; i += 2) {
            String key = (String) inputs[i];
            T value = (T) inputs[i + 1];
            map.put(key, value);
        }
        return map;
    }
}
