
import com.azure.resourcemanager.cosmos.models.CreateUpdateOptions;
import com.azure.resourcemanager.cosmos.models.MongoDBCollectionCreateUpdateParameters;
import com.azure.resourcemanager.cosmos.models.MongoDBCollectionResource;
import com.azure.resourcemanager.cosmos.models.MongoIndex;
import com.azure.resourcemanager.cosmos.models.MongoIndexKeys;
import com.azure.resourcemanager.cosmos.models.MongoIndexOptions;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for MongoDBResources CreateUpdateMongoDBCollection.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2025-04-15/examples/
     * CosmosDBMongoDBCollectionCreateUpdate.json
     */
    /**
     * Sample code: CosmosDBMongoDBCollectionCreateUpdate.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBMongoDBCollectionCreateUpdate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cosmosDBAccounts().manager().serviceClient().getMongoDBResources().createUpdateMongoDBCollection("rg1",
            "ddb1", "databaseName", "collectionName",
            new MongoDBCollectionCreateUpdateParameters().withLocation("West US").withTags(mapOf())
                .withResource(new MongoDBCollectionResource().withId("collectionName")
                    .withShardKey(mapOf("testKey", "fakeTokenPlaceholder"))
                    .withIndexes(Arrays.asList(
                        new MongoIndex().withKey(new MongoIndexKeys().withKeys(Arrays.asList("_ts")))
                            .withOptions(new MongoIndexOptions().withExpireAfterSeconds(100).withUnique(true)),
                        new MongoIndex().withKey(new MongoIndexKeys().withKeys(Arrays.asList("_id"))))))
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
