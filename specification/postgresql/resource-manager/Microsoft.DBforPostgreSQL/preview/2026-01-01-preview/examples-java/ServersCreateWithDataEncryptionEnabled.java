
import com.azure.resourcemanager.postgresqlflexibleserver.models.AzureManagedDiskPerformanceTier;
import com.azure.resourcemanager.postgresqlflexibleserver.models.Backup;
import com.azure.resourcemanager.postgresqlflexibleserver.models.CreateMode;
import com.azure.resourcemanager.postgresqlflexibleserver.models.DataEncryption;
import com.azure.resourcemanager.postgresqlflexibleserver.models.DataEncryptionType;
import com.azure.resourcemanager.postgresqlflexibleserver.models.GeographicallyRedundantBackup;
import com.azure.resourcemanager.postgresqlflexibleserver.models.HighAvailability;
import com.azure.resourcemanager.postgresqlflexibleserver.models.HighAvailabilityMode;
import com.azure.resourcemanager.postgresqlflexibleserver.models.IdentityType;
import com.azure.resourcemanager.postgresqlflexibleserver.models.Network;
import com.azure.resourcemanager.postgresqlflexibleserver.models.PostgresMajorVersion;
import com.azure.resourcemanager.postgresqlflexibleserver.models.Sku;
import com.azure.resourcemanager.postgresqlflexibleserver.models.SkuTier;
import com.azure.resourcemanager.postgresqlflexibleserver.models.Storage;
import com.azure.resourcemanager.postgresqlflexibleserver.models.StorageAutoGrow;
import com.azure.resourcemanager.postgresqlflexibleserver.models.UserAssignedIdentity;
import com.azure.resourcemanager.postgresqlflexibleserver.models.UserIdentity;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Servers CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01-preview/ServersCreateWithDataEncryptionEnabled.json
     */
    /**
     * Sample code: Create a new server with data encryption based on customer managed key.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void createANewServerWithDataEncryptionBasedOnCustomerManagedKey(
        com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.servers().define("exampleserver").withRegion("eastus").withExistingResourceGroup("exampleresourcegroup")
            .withSku(new Sku().withName("Standard_D4ds_v5").withTier(SkuTier.GENERAL_PURPOSE))
            .withIdentity(new UserAssignedIdentity().withUserAssignedIdentities(mapOf(
                "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/exampleresourcegroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/exampleprimaryidentity",
                new UserIdentity())).withType(IdentityType.USER_ASSIGNED))
            .withAdministratorLogin("exampleadministratorlogin").withAdministratorLoginPassword("examplepassword")
            .withVersion(PostgresMajorVersion.ONE_SEVEN)
            .withStorage(new Storage().withStorageSizeGB(512).withAutoGrow(StorageAutoGrow.DISABLED)
                .withTier(AzureManagedDiskPerformanceTier.P20))
            .withDataEncryption(new DataEncryption().withPrimaryKeyUri("fakeTokenPlaceholder")
                .withPrimaryUserAssignedIdentityId(
                    "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/exampleresourcegroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/exampleprimaryidentity")
                .withGeoBackupKeyUri("fakeTokenPlaceholder").withGeoBackupUserAssignedIdentityId("")
                .withType(DataEncryptionType.AZURE_KEY_VAULT))
            .withBackup(
                new Backup().withBackupRetentionDays(7).withGeoRedundantBackup(GeographicallyRedundantBackup.DISABLED))
            .withNetwork(new Network().withDelegatedSubnetResourceId(
                "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/exampleresourcegroup/providers/Microsoft.Network/virtualNetworks/examplevirtualnetwork/subnets/examplesubnet")
                .withPrivateDnsZoneArmResourceId(
                    "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourcegroups/exampleresourcegroup/providers/Microsoft.Network/privateDnsZones/exampleprivatednszone.postgres.database.azure.com"))
            .withHighAvailability(new HighAvailability().withMode(HighAvailabilityMode.ZONE_REDUNDANT))
            .withAvailabilityZone("1").withCreateMode(CreateMode.CREATE).create();
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
