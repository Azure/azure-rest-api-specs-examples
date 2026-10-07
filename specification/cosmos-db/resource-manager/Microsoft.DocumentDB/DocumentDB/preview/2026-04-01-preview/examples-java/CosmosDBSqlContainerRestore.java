
import com.azure.resourcemanager.cosmos.models.CreateMode;
import com.azure.resourcemanager.cosmos.models.CreateUpdateOptions;
import com.azure.resourcemanager.cosmos.models.ResourceRestoreParameters;
import com.azure.resourcemanager.cosmos.models.SqlContainerCreateUpdateParameters;
import com.azure.resourcemanager.cosmos.models.SqlContainerResource;
import java.time.OffsetDateTime;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for SqlResources CreateUpdateSqlContainer.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/CosmosDBSqlContainerRestore.json
     */
    /**
     * Sample code: CosmosDBSqlContainerRestore.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBSqlContainerRestore(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getSqlResources().createUpdateSqlContainer("rg1", "ddb1", "databaseName",
            "containerName",
            new SqlContainerCreateUpdateParameters().withLocation("West US").withTags(mapOf())
                .withResource(new SqlContainerResource().withId("containerName")
                    .withRestoreParameters(new ResourceRestoreParameters().withRestoreSource(
                        "/subscriptions/00000000-1111-2222-3333-444444444444/providers/Microsoft.DocumentDB/locations/WestUS/restorableDatabaseAccounts/restorableDatabaseAccountId")
                        .withRestoreTimestampInUtc(OffsetDateTime.parse("2022-07-20T18:28:00Z"))
                        .withRestoreWithTtlDisabled(true))
                    .withCreateMode(CreateMode.RESTORE))
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
