
import com.azure.core.util.Context;
import com.azure.resourcemanager.sql.models.ManagedInstanceUpdate;
import java.util.HashMap;
import java.util.Map;

/** Samples for ManagedInstances Update. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/ManagedInstanceUpdateMin.json
     */
    /**
     * Sample code: Update managed instance with minimal properties.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        updateManagedInstanceWithMinimalProperties(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getManagedInstances().update("testrg", "testinstance",
            new ManagedInstanceUpdate().withTags(mapOf("tagKey1", "TagValue1")), Context.NONE);
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
