
import com.azure.resourcemanager.sql.models.BackupStorageRedundancy;
import com.azure.resourcemanager.sql.models.ManagedInstanceLicenseType;
import com.azure.resourcemanager.sql.models.ManagedInstanceProxyOverride;
import com.azure.resourcemanager.sql.models.ManagedInstanceUpdate;
import com.azure.resourcemanager.sql.models.Sku;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for ManagedInstances Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/ManagedInstanceUpdateMax.json
     */
    /**
     * Sample code: Update managed instance with all properties.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void updateManagedInstanceWithAllProperties(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getManagedInstances().update("testrg", "testinstance",
            new ManagedInstanceUpdate()
                .withSku(new Sku().withName("GP_Gen4").withTier("GeneralPurpose").withCapacity(8))
                .withTags(mapOf("tagKey1", "fakeTokenPlaceholder")).withAdministratorLogin("dummylogin")
                .withAdministratorLoginPassword("fakeTokenPlaceholder")
                .withLicenseType(ManagedInstanceLicenseType.BASE_PRICE).withVCores(8).withStorageSizeInGB(448)
                .withCollation("SQL_Latin1_General_CP1_CI_AS").withPublicDataEndpointEnabled(false)
                .withProxyOverride(ManagedInstanceProxyOverride.REDIRECT)
                .withMaintenanceConfigurationId(
                    "/subscriptions/20D7082A-0FC7-4468-82BD-542694D5042B/providers/Microsoft.Maintenance/publicMaintenanceConfigurations/SQL_JapanEast_MI_1")
                .withMinimalTlsVersion("1.2").withRequestedBackupStorageRedundancy(BackupStorageRedundancy.GEO),
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
