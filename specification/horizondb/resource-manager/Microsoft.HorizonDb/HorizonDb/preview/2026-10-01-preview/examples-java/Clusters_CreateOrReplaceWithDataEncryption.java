
import com.azure.resourcemanager.horizondb.models.BackupProperties;
import com.azure.resourcemanager.horizondb.models.ClusterProperties;
import com.azure.resourcemanager.horizondb.models.GeoRedundancyState;
import com.azure.resourcemanager.horizondb.models.HorizonDbClusterDataEncryption;
import com.azure.resourcemanager.horizondb.models.HorizonDbClusterDataEncryptionType;
import com.azure.resourcemanager.horizondb.models.HorizonDbClusterIdentityType;
import com.azure.resourcemanager.horizondb.models.HorizonDbClusterManagedServiceIdentity;
import com.azure.resourcemanager.horizondb.models.PostgreSqlVersion;
import com.azure.resourcemanager.horizondb.models.UserAssignedIdentity;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Clusters CreateOrReplace.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/Clusters_CreateOrReplaceWithDataEncryption.json
     */
    /**
     * Sample code: Create a HorizonDB cluster with customer-managed key data encryption.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void createAHorizonDBClusterWithCustomerManagedKeyDataEncryption(
        com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        manager.clusters().define("example-cluster").withRegion("eastus2")
            .withExistingResourceGroup("example-resource-group")
            .withProperties(new ClusterProperties().withPostgreSqlVersion(PostgreSqlVersion.V17)
                .withBackup(new BackupProperties().withGeoRedundancy(GeoRedundancyState.DISABLED).withRetentionDays(14))
                .withDataEncryption(new HorizonDbClusterDataEncryption()
                    .withType(HorizonDbClusterDataEncryptionType.AZURE_KEY_VAULT)
                    .withPrimaryKeyUri("fakeTokenPlaceholder").withPrimaryUserAssignedIdentityId(
                        "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/identity-resource-group/providers/Microsoft.ManagedIdentity/userAssignedIdentities/primary-encryption-identity")))
            .withIdentity(new HorizonDbClusterManagedServiceIdentity()
                .withType(HorizonDbClusterIdentityType.USER_ASSIGNED)
                .withUserAssignedIdentities(mapOf(
                    "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/identity-resource-group/providers/Microsoft.ManagedIdentity/userAssignedIdentities/primary-encryption-identity",
                    new UserAssignedIdentity())))
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
