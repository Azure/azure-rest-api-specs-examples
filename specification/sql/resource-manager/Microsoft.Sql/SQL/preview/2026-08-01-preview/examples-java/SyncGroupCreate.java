
import com.azure.resourcemanager.sql.fluent.models.SyncGroupInner;
import com.azure.resourcemanager.sql.models.SyncConflictResolutionPolicy;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for SyncGroups CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01-preview/SyncGroupCreate.json
     */
    /**
     * Sample code: Create a sync group.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void createASyncGroup(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getSyncGroups().createOrUpdate("syncgroupcrud-65440", "syncgroupcrud-8475",
            "syncgroupcrud-4328", "syncgroupcrud-3187",
            new SyncGroupInner().withInterval(-1).withConflictResolutionPolicy(SyncConflictResolutionPolicy.HUB_WIN)
                .withSyncDatabaseId(
                    "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/syncgroupcrud-3521/providers/Microsoft.Sql/servers/syncgroupcrud-8475/databases/syncgroupcrud-4328")
                .withHubDatabaseUsername("hubUser").withUsePrivateLinkConnection(true),
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
