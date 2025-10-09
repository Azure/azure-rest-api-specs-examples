
import com.azure.resourcemanager.mongocluster.models.AdministratorProperties;
import com.azure.resourcemanager.mongocluster.models.ComputeProperties;
import com.azure.resourcemanager.mongocluster.models.CustomerManagedKeyEncryptionProperties;
import com.azure.resourcemanager.mongocluster.models.EncryptionProperties;
import com.azure.resourcemanager.mongocluster.models.HighAvailabilityMode;
import com.azure.resourcemanager.mongocluster.models.HighAvailabilityProperties;
import com.azure.resourcemanager.mongocluster.models.KeyEncryptionKeyIdentity;
import com.azure.resourcemanager.mongocluster.models.KeyEncryptionKeyIdentityType;
import com.azure.resourcemanager.mongocluster.models.ManagedServiceIdentity;
import com.azure.resourcemanager.mongocluster.models.ManagedServiceIdentityType;
import com.azure.resourcemanager.mongocluster.models.MongoClusterProperties;
import com.azure.resourcemanager.mongocluster.models.ShardingProperties;
import com.azure.resourcemanager.mongocluster.models.StorageProperties;
import com.azure.resourcemanager.mongocluster.models.UserAssignedIdentity;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for MongoClusters CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-01-preview/MongoClusters_Create_CMK.json
     */
    /**
     * Sample code: Creates a new Mongo Cluster resource with Customer Managed Key encryption.
     * 
     * @param manager Entry point to MongoClusterManager.
     */
    public static void createsANewMongoClusterResourceWithCustomerManagedKeyEncryption(
        com.azure.resourcemanager.mongocluster.MongoClusterManager manager) {
        manager.mongoClusters().define("myMongoCluster").withRegion("westus2")
            .withExistingResourceGroup("TestResourceGroup")
            .withProperties(new MongoClusterProperties()
                .withAdministrator(
                    new AdministratorProperties().withUserName("mongoAdmin").withPassword("fakeTokenPlaceholder"))
                .withHighAvailability(new HighAvailabilityProperties().withTargetMode(HighAvailabilityMode.DISABLED))
                .withStorage(new StorageProperties().withSizeGb(32L))
                .withSharding(new ShardingProperties().withShardCount(1))
                .withCompute(new ComputeProperties().withTier("M30"))
                .withEncryption(new EncryptionProperties()
                    .withCustomerManagedKeyEncryption(new CustomerManagedKeyEncryptionProperties()
                        .withKeyEncryptionKeyIdentity(new KeyEncryptionKeyIdentity()
                            .withIdentityType(KeyEncryptionKeyIdentityType.USER_ASSIGNED_IDENTITY)
                            .withUserAssignedIdentityResourceId(
                                "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestResourceGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/myidentity"))
                        .withKeyEncryptionKeyUrl("fakeTokenPlaceholder"))))
            .withIdentity(new ManagedServiceIdentity().withType(ManagedServiceIdentityType.USER_ASSIGNED)
                .withUserAssignedIdentities(mapOf(
                    "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestResourceGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/myidentity",
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
