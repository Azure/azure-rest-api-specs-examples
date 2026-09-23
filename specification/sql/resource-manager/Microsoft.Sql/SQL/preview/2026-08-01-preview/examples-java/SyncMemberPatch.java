
import com.azure.resourcemanager.sql.fluent.models.SyncMemberInner;
import com.azure.resourcemanager.sql.models.SyncDirection;
import com.azure.resourcemanager.sql.models.SyncMemberDbType;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for SyncMembers Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01-preview/SyncMemberPatch.json
     */
    /**
     * Sample code: Update an existing sync member.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void updateAnExistingSyncMember(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getSyncMembers().update("syncgroupcrud-65440", "syncgroupcrud-8475",
            "syncgroupcrud-4328", "syncgroupcrud-3187", "syncmembercrud-4879",
            new SyncMemberInner().withDatabaseType(SyncMemberDbType.AZURE_SQL_DATABASE)
                .withSyncMemberAzureDatabaseResourceId(
                    "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/syncgroupcrud-65440/providers/Microsoft.Sql/servers/syncgroupcrud-8475/databases/syncgroupcrud-4328")
                .withUsePrivateLinkConnection(true).withServerName("syncgroupcrud-3379.database.windows.net")
                .withDatabaseName("syncgroupcrud-7421").withUsername("myUser")
                .withSyncDirection(SyncDirection.BIDIRECTIONAL),
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
