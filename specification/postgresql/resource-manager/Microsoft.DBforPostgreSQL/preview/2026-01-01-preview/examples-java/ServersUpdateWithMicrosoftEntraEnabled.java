
import com.azure.resourcemanager.postgresqlflexibleserver.models.AuthConfigForPatch;
import com.azure.resourcemanager.postgresqlflexibleserver.models.AzureManagedDiskPerformanceTier;
import com.azure.resourcemanager.postgresqlflexibleserver.models.BackupForPatch;
import com.azure.resourcemanager.postgresqlflexibleserver.models.CreateModeForPatch;
import com.azure.resourcemanager.postgresqlflexibleserver.models.MicrosoftEntraAuth;
import com.azure.resourcemanager.postgresqlflexibleserver.models.PasswordBasedAuth;
import com.azure.resourcemanager.postgresqlflexibleserver.models.Server;
import com.azure.resourcemanager.postgresqlflexibleserver.models.SkuForPatch;
import com.azure.resourcemanager.postgresqlflexibleserver.models.SkuTier;
import com.azure.resourcemanager.postgresqlflexibleserver.models.Storage;
import com.azure.resourcemanager.postgresqlflexibleserver.models.StorageAutoGrow;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Servers Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01-preview/ServersUpdateWithMicrosoftEntraEnabled.json
     */
    /**
     * Sample code: Update an existing server with Microsoft Entra authentication enabled.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void updateAnExistingServerWithMicrosoftEntraAuthenticationEnabled(
        com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        Server resource = manager.servers()
            .getByResourceGroupWithResponse("exampleresourcegroup", "exampleserver", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update().withSku(new SkuForPatch().withName("Standard_D8s_v3").withTier(SkuTier.GENERAL_PURPOSE))
            .withAdministratorLoginPassword("examplenewpassword")
            .withStorage(new Storage().withStorageSizeGB(1024).withAutoGrow(StorageAutoGrow.DISABLED)
                .withTier(AzureManagedDiskPerformanceTier.P30))
            .withBackup(new BackupForPatch().withBackupRetentionDays(20))
            .withAuthConfig(new AuthConfigForPatch().withActiveDirectoryAuth(MicrosoftEntraAuth.ENABLED)
                .withPasswordAuth(PasswordBasedAuth.ENABLED).withTenantId("tttttt-tttt-tttt-tttt-tttttttttttt"))
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
