
import com.azure.resourcemanager.sql.models.AuthMetadataLookupModes;
import com.azure.resourcemanager.sql.models.AvailabilityZoneType;
import com.azure.resourcemanager.sql.models.BackupStorageRedundancy;
import com.azure.resourcemanager.sql.models.HybridSecondaryUsage;
import com.azure.resourcemanager.sql.models.ManagedInstanceDatabaseFormat;
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
     * x-ms-original-file: 2025-01-01/ManagedInstanceUpdateMax.json
     */
    /**
     * Sample code: Update managed instance with all properties.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void updateManagedInstanceWithAllProperties(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getManagedInstances().update("testrg", "testinstance", new ManagedInstanceUpdate()
            .withSku(new Sku().withName("GP_Gen5").withTier("GeneralPurpose").withCapacity(8))
            .withTags(mapOf("tagKey1", "fakeTokenPlaceholder")).withAdministratorLogin("dummylogin")
            .withAdministratorLoginPassword("fakeTokenPlaceholder")
            .withLicenseType(ManagedInstanceLicenseType.BASE_PRICE)
            .withHybridSecondaryUsage(HybridSecondaryUsage.PASSIVE).withVCores(8).withStorageSizeInGB(448)
            .withCollation("SQL_Latin1_General_CP1_CI_AS").withPublicDataEndpointEnabled(false)
            .withProxyOverride(ManagedInstanceProxyOverride.REDIRECT)
            .withMaintenanceConfigurationId(
                "/subscriptions/00000000-1111-2222-3333-444444444444/providers/Microsoft.Maintenance/publicMaintenanceConfigurations/SQL_JapanEast_MI_1")
            .withMinimalTlsVersion("1.2").withRequestedBackupStorageRedundancy(BackupStorageRedundancy.GEO)
            .withAuthenticationMetadata(AuthMetadataLookupModes.WINDOWS)
            .withDatabaseFormat(ManagedInstanceDatabaseFormat.ALWAYS_UP_TO_DATE)
            .withRequestedLogicalAvailabilityZone(AvailabilityZoneType.ONE), com.azure.core.util.Context.NONE);
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
