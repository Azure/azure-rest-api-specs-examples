
import com.azure.resourcemanager.postgresqlflexibleserver.models.CreateMode;
import com.azure.resourcemanager.postgresqlflexibleserver.models.DataEncryption;
import com.azure.resourcemanager.postgresqlflexibleserver.models.DataEncryptionType;
import com.azure.resourcemanager.postgresqlflexibleserver.models.IdentityType;
import com.azure.resourcemanager.postgresqlflexibleserver.models.UserAssignedIdentity;
import com.azure.resourcemanager.postgresqlflexibleserver.models.UserIdentity;
import java.time.OffsetDateTime;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Servers CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2025-08-01/examples/
     * ServersCreateReplica.json
     */
    /**
     * Sample code: Create a read replica of an existing server.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void createAReadReplicaOfAnExistingServer(
        com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.servers().define("exampleserver").withRegion("eastus").withExistingResourceGroup("exampleresourcegroup")
            .withIdentity(new UserAssignedIdentity().withUserAssignedIdentities(mapOf(
                "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/exampleresourcegroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/exampleprimaryidentity",
                new UserIdentity())).withType(IdentityType.USER_ASSIGNED))
            .withDataEncryption(new DataEncryption().withPrimaryKeyUri("fakeTokenPlaceholder")
                .withPrimaryUserAssignedIdentityId(
                    "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/exampleresourcegroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/exampleprimaryidentity")
                .withGeoBackupKeyUri("fakeTokenPlaceholder").withGeoBackupUserAssignedIdentityId("")
                .withType(DataEncryptionType.AZURE_KEY_VAULT))
            .withSourceServerResourceId(
                "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/exampleresourcegroup/providers/Microsoft.DBforPostgreSQL/flexibleServers/examplesourceserver")
            .withPointInTimeUtc(OffsetDateTime.parse("2025-06-01T18:35:22.123456Z")).withCreateMode(CreateMode.REPLICA)
            .create();
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
