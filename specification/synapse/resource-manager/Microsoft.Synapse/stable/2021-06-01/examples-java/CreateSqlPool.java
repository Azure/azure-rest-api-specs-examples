import com.azure.resourcemanager.synapse.models.CreateMode;
import com.azure.resourcemanager.synapse.models.Sku;
import com.azure.resourcemanager.synapse.models.StorageAccountType;
import java.util.HashMap;
import java.util.Map;

/** Samples for SqlPools Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/CreateSqlPool.json
     */
    /**
     * Sample code: Create a SQL Analytics pool.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void createASQLAnalyticsPool(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .sqlPools()
            .define("ExampleSqlPool")
            .withRegion("Southeast Asia")
            .withExistingWorkspace("ExampleResourceGroup", "ExampleWorkspace")
            .withTags(mapOf())
            .withSku(new Sku().withTier("").withName(""))
            .withMaxSizeBytes(0L)
            .withCollation("")
            .withSourceDatabaseId("")
            .withRecoverableDatabaseId("")
            .withCreateMode(CreateMode.fromString(""))
            .withStorageAccountType(StorageAccountType.LRS)
            .create();
    }

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
