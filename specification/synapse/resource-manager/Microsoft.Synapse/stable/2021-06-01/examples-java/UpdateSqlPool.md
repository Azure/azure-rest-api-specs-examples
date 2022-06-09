```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.synapse.models.Sku;
import com.azure.resourcemanager.synapse.models.SqlPool;
import java.time.OffsetDateTime;
import java.util.HashMap;
import java.util.Map;

/** Samples for SqlPools Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/UpdateSqlPool.json
     */
    /**
     * Sample code: Update a SQL Analytics pool.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void updateASQLAnalyticsPool(com.azure.resourcemanager.synapse.SynapseManager manager) {
        SqlPool resource =
            manager
                .sqlPools()
                .getWithResponse("ExampleResourceGroup", "ExampleWorkspace", "ExampleSqlPool", Context.NONE)
                .getValue();
        resource
            .update()
            .withTags(mapOf())
            .withSku(new Sku().withTier("").withName(""))
            .withMaxSizeBytes(0L)
            .withCollation("")
            .withRestorePointInTime(OffsetDateTime.parse("1970-01-01T00:00:00.000Z"))
            .apply();
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
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-synapse_1.0.0-beta.6/sdk/synapse/azure-resourcemanager-synapse/README.md) on how to add the SDK to your project and authenticate.
