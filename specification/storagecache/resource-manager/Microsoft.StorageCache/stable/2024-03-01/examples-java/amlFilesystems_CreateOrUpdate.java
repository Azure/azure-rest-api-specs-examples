
import com.azure.resourcemanager.storagecache.models.AmlFilesystemEncryptionSettings;
import com.azure.resourcemanager.storagecache.models.AmlFilesystemHsmSettings;
import com.azure.resourcemanager.storagecache.models.AmlFilesystemIdentity;
import com.azure.resourcemanager.storagecache.models.AmlFilesystemIdentityType;
import com.azure.resourcemanager.storagecache.models.AmlFilesystemPropertiesHsm;
import com.azure.resourcemanager.storagecache.models.AmlFilesystemPropertiesMaintenanceWindow;
import com.azure.resourcemanager.storagecache.models.AmlFilesystemRootSquashSettings;
import com.azure.resourcemanager.storagecache.models.AmlFilesystemSquashMode;
import com.azure.resourcemanager.storagecache.models.KeyVaultKeyReference;
import com.azure.resourcemanager.storagecache.models.KeyVaultKeyReferenceSourceVault;
import com.azure.resourcemanager.storagecache.models.MaintenanceDayOfWeekType;
import com.azure.resourcemanager.storagecache.models.SkuName;
import com.azure.resourcemanager.storagecache.models.UserAssignedIdentitiesValue;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for AmlFilesystems CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/storagecache/resource-manager/Microsoft.StorageCache/stable/2024-03-01/examples/
     * amlFilesystems_CreateOrUpdate.json
     */
    /**
     * Sample code: amlFilesystems_CreateOrUpdate.
     * 
     * @param manager Entry point to StorageCacheManager.
     */
    public static void
        amlFilesystemsCreateOrUpdate(com.azure.resourcemanager.storagecache.StorageCacheManager manager) {
        manager.amlFilesystems().define("fs1").withRegion("eastus").withExistingResourceGroup("scgroup")
            .withTags(mapOf("Dept", "ContosoAds"))
            .withIdentity(new AmlFilesystemIdentity().withType(AmlFilesystemIdentityType.USER_ASSIGNED)
                .withUserAssignedIdentities(mapOf(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/scgroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity1",
                    new UserAssignedIdentitiesValue())))
            .withSku(new SkuName().withName("AMLFS-Durable-Premium-250")).withZones(Arrays.asList("1"))
            .withStorageCapacityTiB(16f)
            .withFilesystemSubnet(
                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/scgroup/providers/Microsoft.Network/virtualNetworks/scvnet/subnets/fsSub")
            .withEncryptionSettings(new AmlFilesystemEncryptionSettings()
                .withKeyEncryptionKey(new KeyVaultKeyReference().withKeyUrl("fakeTokenPlaceholder")
                    .withSourceVault(new KeyVaultKeyReferenceSourceVault().withId(
                        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/scgroup/providers/Microsoft.KeyVault/vaults/keyvault-cmk"))))
            .withMaintenanceWindow(new AmlFilesystemPropertiesMaintenanceWindow()
                .withDayOfWeek(MaintenanceDayOfWeekType.FRIDAY).withTimeOfDayUtc("22:00"))
            .withHsm(new AmlFilesystemPropertiesHsm().withSettings(new AmlFilesystemHsmSettings().withContainer(
                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/scgroup/providers/Microsoft.Storage/storageAccounts/storageaccountname/blobServices/default/containers/containername")
                .withLoggingContainer(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/scgroup/providers/Microsoft.Storage/storageAccounts/storageaccountname/blobServices/default/containers/loggingcontainername")
                .withImportPrefixesInitial(Arrays.asList("/"))))
            .withRootSquashSettings(new AmlFilesystemRootSquashSettings().withMode(AmlFilesystemSquashMode.ALL)
                .withNoSquashNidLists("10.0.0.[5-6]@tcp;10.0.1.2@tcp").withSquashUid(99L).withSquashGid(99L))
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
