
import com.azure.resourcemanager.cosmos.models.CassandraKeyspaceCreateUpdateParameters;
import com.azure.resourcemanager.cosmos.models.CassandraKeyspaceResource;
import com.azure.resourcemanager.cosmos.models.CreateUpdateOptions;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for CassandraResources CreateUpdateCassandraKeyspace.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2025-04-15/examples/
     * CosmosDBCassandraKeyspaceCreateUpdate.json
     */
    /**
     * Sample code: CosmosDBCassandraKeyspaceCreateUpdate.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBCassandraKeyspaceCreateUpdate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cosmosDBAccounts().manager().serviceClient().getCassandraResources().createUpdateCassandraKeyspace("rg1",
            "ddb1", "keyspaceName",
            new CassandraKeyspaceCreateUpdateParameters().withLocation("West US").withTags(mapOf())
                .withResource(new CassandraKeyspaceResource().withId("keyspaceName"))
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
