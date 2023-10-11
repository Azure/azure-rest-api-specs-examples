import com.azure.resourcemanager.postgresqlflexibleserver.models.ArmServerKeyType;
import com.azure.resourcemanager.postgresqlflexibleserver.models.Backup;
import com.azure.resourcemanager.postgresqlflexibleserver.models.CreateModeForUpdate;
import com.azure.resourcemanager.postgresqlflexibleserver.models.DataEncryption;
import com.azure.resourcemanager.postgresqlflexibleserver.models.IdentityType;
import com.azure.resourcemanager.postgresqlflexibleserver.models.Server;
import com.azure.resourcemanager.postgresqlflexibleserver.models.Sku;
import com.azure.resourcemanager.postgresqlflexibleserver.models.SkuTier;
import com.azure.resourcemanager.postgresqlflexibleserver.models.Storage;
import com.azure.resourcemanager.postgresqlflexibleserver.models.UserAssignedIdentity;
import com.azure.resourcemanager.postgresqlflexibleserver.models.UserIdentity;
import java.util.HashMap;
import java.util.Map;

/** Samples for Servers Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2022-12-01/examples/ServerUpdateWithDataEncryptionEnabled.json
     */
    /**
     * Sample code: ServerUpdateWithDataEncryptionEnabled.
     *
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void serverUpdateWithDataEncryptionEnabled(
        com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        Server resource =
            manager
                .servers()
                .getByResourceGroupWithResponse("TestGroup", "pgtestsvc4", com.azure.core.util.Context.NONE)
                .getValue();
        resource
            .update()
            .withSku(new Sku().withName("Standard_D8s_v3").withTier(SkuTier.GENERAL_PURPOSE))
            .withIdentity(
                new UserAssignedIdentity()
                    .withUserAssignedIdentities(
                        mapOf(
                            "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testresourcegroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/test-usermanagedidentity",
                            new UserIdentity()))
                    .withType(IdentityType.USER_ASSIGNED))
            .withAdministratorLoginPassword("newpassword")
            .withStorage(new Storage().withStorageSizeGB(1024))
            .withBackup(new Backup().withBackupRetentionDays(20))
            .withDataEncryption(
                new DataEncryption()
                    .withPrimaryKeyUri("fakeTokenPlaceholder")
                    .withPrimaryUserAssignedIdentityId(
                        "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testresourcegroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/test-usermanagedidentity")
                    .withType(ArmServerKeyType.AZURE_KEY_VAULT))
            .withCreateMode(CreateModeForUpdate.UPDATE)
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
