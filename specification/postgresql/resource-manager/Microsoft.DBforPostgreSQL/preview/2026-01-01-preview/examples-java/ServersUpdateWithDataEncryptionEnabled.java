
import com.azure.resourcemanager.postgresqlflexibleserver.models.BackupForPatch;
import com.azure.resourcemanager.postgresqlflexibleserver.models.CreateModeForPatch;
import com.azure.resourcemanager.postgresqlflexibleserver.models.DataEncryption;
import com.azure.resourcemanager.postgresqlflexibleserver.models.DataEncryptionType;
import com.azure.resourcemanager.postgresqlflexibleserver.models.IdentityType;
import com.azure.resourcemanager.postgresqlflexibleserver.models.Server;
import com.azure.resourcemanager.postgresqlflexibleserver.models.SkuForPatch;
import com.azure.resourcemanager.postgresqlflexibleserver.models.SkuTier;
import com.azure.resourcemanager.postgresqlflexibleserver.models.UserAssignedIdentity;
import com.azure.resourcemanager.postgresqlflexibleserver.models.UserIdentity;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Servers Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01-preview/ServersUpdateWithDataEncryptionEnabled.json
     */
    /**
     * Sample code: Update an existing server with data encryption based on customer managed key.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void updateAnExistingServerWithDataEncryptionBasedOnCustomerManagedKey(
        com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        Server resource = manager.servers()
            .getByResourceGroupWithResponse("exampleresourcegroup", "exampleserver", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update().withSku(new SkuForPatch().withName("Standard_D8s_v3").withTier(SkuTier.GENERAL_PURPOSE))
            .withIdentity(new UserAssignedIdentity().withUserAssignedIdentities(mapOf(
                "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/exampleresourcegroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/examplegeoredundantidentity",
                new UserIdentity(),
                "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/exampleresourcegroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/exampleprimaryidentity",
                new UserIdentity())).withType(IdentityType.USER_ASSIGNED))
            .withAdministratorLoginPassword("examplenewpassword")
            .withBackup(new BackupForPatch().withBackupRetentionDays(20))
            .withDataEncryption(new DataEncryption().withPrimaryKeyUri("fakeTokenPlaceholder")
                .withPrimaryUserAssignedIdentityId(
                    "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/exampleresourcegroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/exampleprimaryidentity")
                .withGeoBackupKeyUri("fakeTokenPlaceholder")
                .withGeoBackupUserAssignedIdentityId(
                    "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/exampleresourcegroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/examplegeoredundantidentity")
                .withType(DataEncryptionType.AZURE_KEY_VAULT))
            .withCreateMode(CreateModeForPatch.UPDATE).apply();
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
