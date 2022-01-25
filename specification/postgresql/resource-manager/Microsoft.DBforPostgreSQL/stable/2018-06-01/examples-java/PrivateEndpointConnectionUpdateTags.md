Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-postgresql_1.0.2/sdk/postgresql/azure-resourcemanager-postgresql/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.postgresql.models.PrivateEndpointConnection;
import java.util.HashMap;
import java.util.Map;

/** Samples for PrivateEndpointConnections UpdateTags. */
public final class Main {
    /*
     * x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2018-06-01/examples/PrivateEndpointConnectionUpdateTags.json
     */
    /**
     * Sample code: Update private endpoint connection Tags.
     *
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void updatePrivateEndpointConnectionTags(
        com.azure.resourcemanager.postgresql.PostgreSqlManager manager) {
        PrivateEndpointConnection resource =
            manager
                .privateEndpointConnections()
                .getWithResponse("Default", "test-svr", "private-endpoint-connection-name", Context.NONE)
                .getValue();
        resource.update().withTags(mapOf("key1", "val1", "key2", "val2", "key3", "val3")).apply();
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
