
import com.azure.resourcemanager.sql.models.JobAgentUpdate;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for JobAgents Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/UpdateJobAgent.json
     */
    /**
     * Sample code: Update a job agent's tags.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void updateAJobAgentSTags(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getJobAgents().update("group1", "server1", "agent1",
            new JobAgentUpdate().withTags(mapOf("mytag1", "myvalue1")), com.azure.core.util.Context.NONE);
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
