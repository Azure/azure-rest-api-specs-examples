
import com.azure.resourcemanager.sql.fluent.models.JobAgentInner;
import com.azure.resourcemanager.sql.models.Sku;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for JobAgents CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/CreateOrUpdateJobAgentWithSku.json
     */
    /**
     * Sample code: Create or update a job agent with sku.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void createOrUpdateAJobAgentWithSku(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getJobAgents().createOrUpdate("group1", "server1", "agent1",
            new JobAgentInner().withLocation("southeastasia").withSku(new Sku().withName("JA400")).withDatabaseId(
                "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/server1/databases/db1"),
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
