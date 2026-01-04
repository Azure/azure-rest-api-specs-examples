
import com.azure.resourcemanager.cosmos.models.ApiProperties;
import com.azure.resourcemanager.cosmos.models.ConsistencyPolicy;
import com.azure.resourcemanager.cosmos.models.ContinuousModeBackupPolicy;
import com.azure.resourcemanager.cosmos.models.ContinuousModeProperties;
import com.azure.resourcemanager.cosmos.models.ContinuousTier;
import com.azure.resourcemanager.cosmos.models.CreateMode;
import com.azure.resourcemanager.cosmos.models.DatabaseAccountCreateUpdateParameters;
import com.azure.resourcemanager.cosmos.models.DatabaseAccountKind;
import com.azure.resourcemanager.cosmos.models.DatabaseRestoreResource;
import com.azure.resourcemanager.cosmos.models.DefaultConsistencyLevel;
import com.azure.resourcemanager.cosmos.models.Location;
import com.azure.resourcemanager.cosmos.models.MinimalTlsVersion;
import com.azure.resourcemanager.cosmos.models.RestoreMode;
import com.azure.resourcemanager.cosmos.models.RestoreParameters;
import com.azure.resourcemanager.cosmos.models.ServerVersion;
import java.time.OffsetDateTime;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for DatabaseAccounts CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cosmos-db/resource-manager/Microsoft.DocumentDB/DocumentDB/stable/2025-10-15/examples/
     * CosmosDBRestoreDatabaseAccountCreateUpdate.json
     */
    /**
     * Sample code: CosmosDBRestoreDatabaseAccountCreateUpdate.json.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        cosmosDBRestoreDatabaseAccountCreateUpdateJson(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cosmosDBAccounts().manager().serviceClient().getDatabaseAccounts().createOrUpdate("rg1", "ddb1",
            new DatabaseAccountCreateUpdateParameters().withLocation("westus").withTags(mapOf())
                .withKind(DatabaseAccountKind.GLOBAL_DOCUMENT_DB)
                .withConsistencyPolicy(
                    new ConsistencyPolicy().withDefaultConsistencyLevel(DefaultConsistencyLevel.BOUNDED_STALENESS)
                        .withMaxStalenessPrefix(200L).withMaxIntervalInSeconds(10))
                .withLocations(Arrays.asList(new Location().withLocationName("southcentralus").withFailoverPriority(0)
                    .withIsZoneRedundant(false)))
                .withKeyVaultKeyUri("fakeTokenPlaceholder").withEnableFreeTier(false)
                .withApiProperties(new ApiProperties().withServerVersion(ServerVersion.THREE_TWO))
                .withEnableAnalyticalStorage(true).withCreateMode(CreateMode.RESTORE)
                .withBackupPolicy(new ContinuousModeBackupPolicy().withContinuousModeProperties(
                    new ContinuousModeProperties().withTier(ContinuousTier.CONTINUOUS30DAYS)))
                .withRestoreParameters(new RestoreParameters().withRestoreSource(
                    "/subscriptions/subid/providers/Microsoft.DocumentDB/locations/westus/restorableDatabaseAccounts/1a97b4bb-f6a0-430e-ade1-638d781830cc")
                    .withRestoreTimestampInUtc(OffsetDateTime.parse("2021-03-11T22:05:09Z"))
                    .withRestoreWithTtlDisabled(false).withRestoreMode(RestoreMode.POINT_IN_TIME)
                    .withDatabasesToRestore(Arrays.asList(
                        new DatabaseRestoreResource().withDatabaseName("db1")
                            .withCollectionNames(Arrays.asList("collection1", "collection2")),
                        new DatabaseRestoreResource().withDatabaseName("db2")
                            .withCollectionNames(Arrays.asList("collection3", "collection4"))))
                    .withSourceBackupLocation("westus"))
                .withMinimalTlsVersion(MinimalTlsVersion.TLS),
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
