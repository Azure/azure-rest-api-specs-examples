
import com.azure.resourcemanager.cosmos.models.CreateUpdateOptions;
import com.azure.resourcemanager.cosmos.models.MongoDBDatabaseCreateUpdateParameters;
import com.azure.resourcemanager.cosmos.models.MongoDBDatabaseResource;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for MongoDBResources CreateUpdateMongoDBDatabase.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBMongoDBDatabaseCreateUpdate.json
     */
    /**
     * Sample code: CosmosDBMongoDBDatabaseCreateUpdate.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBMongoDBDatabaseCreateUpdate(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getMongoDBResources().createUpdateMongoDBDatabase("rg1", "ddb1", "databaseName",
            new MongoDBDatabaseCreateUpdateParameters().withLocation("West US").withTags(mapOf())
                .withResource(new MongoDBDatabaseResource().withId("databaseName"))
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
