
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
     * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2025-04-15/examples/
     * CosmosDBMongoDBDatabaseCreateUpdate.json
     */
    /**
     * Sample code: CosmosDBMongoDBDatabaseCreateUpdate.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBMongoDBDatabaseCreateUpdate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cosmosDBAccounts().manager().serviceClient().getMongoDBResources().createUpdateMongoDBDatabase("rg1",
            "ddb1", "databaseName",
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
