import com.azure.resourcemanager.postgresqlflexibleserver.models.ActiveDirectoryAuthEnum;
import com.azure.resourcemanager.postgresqlflexibleserver.models.AuthConfig;
import com.azure.resourcemanager.postgresqlflexibleserver.models.AzureManagedDiskPerformanceTiers;
import com.azure.resourcemanager.postgresqlflexibleserver.models.Backup;
import com.azure.resourcemanager.postgresqlflexibleserver.models.CreateModeForUpdate;
import com.azure.resourcemanager.postgresqlflexibleserver.models.PasswordAuthEnum;
import com.azure.resourcemanager.postgresqlflexibleserver.models.Server;
import com.azure.resourcemanager.postgresqlflexibleserver.models.Sku;
import com.azure.resourcemanager.postgresqlflexibleserver.models.SkuTier;
import com.azure.resourcemanager.postgresqlflexibleserver.models.Storage;
import com.azure.resourcemanager.postgresqlflexibleserver.models.StorageAutoGrow;
import java.util.HashMap;
import java.util.Map;

/** Samples for Servers Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2023-03-01-preview/examples/ServerUpdateWithAadAuthEnabled.json
     */
    /**
     * Sample code: ServerUpdateWithAadAuthEnabled.
     *
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void serverUpdateWithAadAuthEnabled(
        com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        Server resource =
            manager
                .servers()
                .getByResourceGroupWithResponse("TestGroup", "pgtestsvc4", com.azure.core.util.Context.NONE)
                .getValue();
        resource
            .update()
            .withSku(new Sku().withName("Standard_D8s_v3").withTier(SkuTier.GENERAL_PURPOSE))
            .withAdministratorLoginPassword("newpassword")
            .withStorage(
                new Storage()
                    .withStorageSizeGB(1024)
                    .withAutoGrow(StorageAutoGrow.DISABLED)
                    .withTier(AzureManagedDiskPerformanceTiers.P30))
            .withBackup(new Backup().withBackupRetentionDays(20))
            .withAuthConfig(
                new AuthConfig()
                    .withActiveDirectoryAuth(ActiveDirectoryAuthEnum.ENABLED)
                    .withPasswordAuth(PasswordAuthEnum.ENABLED)
                    .withTenantId("tttttt-tttt-tttt-tttt-tttttttttttt"))
            .withCreateMode(CreateModeForUpdate.UPDATE)
            .apply();
    }

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
