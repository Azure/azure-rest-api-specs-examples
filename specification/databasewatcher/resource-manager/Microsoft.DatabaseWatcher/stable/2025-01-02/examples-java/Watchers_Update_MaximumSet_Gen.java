
import com.azure.resourcemanager.databasewatcher.models.Datastore;
import com.azure.resourcemanager.databasewatcher.models.KustoOfferingType;
import com.azure.resourcemanager.databasewatcher.models.ManagedServiceIdentityType;
import com.azure.resourcemanager.databasewatcher.models.ManagedServiceIdentityV4;
import com.azure.resourcemanager.databasewatcher.models.Watcher;
import com.azure.resourcemanager.databasewatcher.models.WatcherUpdateProperties;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Watchers Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-02/Watchers_Update_MaximumSet_Gen.json
     */
    /**
     * Sample code: Watchers_Update_MaximumSet.
     * 
     * @param manager Entry point to DatabaseWatcherManager.
     */
    public static void
        watchersUpdateMaximumSet(com.azure.resourcemanager.databasewatcher.DatabaseWatcherManager manager) {
        Watcher resource = manager.watchers()
            .getByResourceGroupWithResponse("rgWatcher", "testWatcher", com.azure.core.util.Context.NONE).getValue();
        resource.update().withTags(mapOf())
            .withIdentity(new ManagedServiceIdentityV4().withType(ManagedServiceIdentityType.SYSTEM_ASSIGNED))
            .withProperties(new WatcherUpdateProperties().withDatastore(new Datastore().withAdxClusterResourceId(
                "/subscriptions/49e0fbd3-75e8-44e7-96fd-5b64d9ad818d/resourceGroups/apiTest/providers/Microsoft.Kusto/clusters/apiTestKusto")
                .withKustoClusterDisplayName("kustoUri-adx")
                .withKustoClusterUri("https://kustouri-adx.eastus.kusto.windows.net")
                .withKustoDataIngestionUri("https://ingest-kustouri-adx.eastus.kusto.windows.net")
                .withKustoDatabaseName("kustoDatabaseName1").withKustoManagementUrl("https://portal.azure.com/")
                .withKustoOfferingType(KustoOfferingType.ADX)).withDefaultAlertRuleIdentityResourceId(
                    "/subscriptions/469DD77C-C8DB-47B7-B9E1-72D29F8C878B/resourceGroups/rgWatcher/providers/Microsoft.ManagedIdentity/userAssignedIdentities/newtest"))
            .apply();
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
