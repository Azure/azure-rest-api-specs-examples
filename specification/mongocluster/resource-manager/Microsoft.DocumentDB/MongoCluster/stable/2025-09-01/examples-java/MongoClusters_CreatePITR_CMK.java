
import com.azure.resourcemanager.mongocluster.models.CreateMode;
import com.azure.resourcemanager.mongocluster.models.CustomerManagedKeyEncryptionProperties;
import com.azure.resourcemanager.mongocluster.models.EncryptionProperties;
import com.azure.resourcemanager.mongocluster.models.KeyEncryptionKeyIdentity;
import com.azure.resourcemanager.mongocluster.models.KeyEncryptionKeyIdentityType;
import com.azure.resourcemanager.mongocluster.models.ManagedServiceIdentity;
import com.azure.resourcemanager.mongocluster.models.ManagedServiceIdentityType;
import com.azure.resourcemanager.mongocluster.models.MongoClusterProperties;
import com.azure.resourcemanager.mongocluster.models.MongoClusterRestoreParameters;
import com.azure.resourcemanager.mongocluster.models.UserAssignedIdentity;
import java.time.OffsetDateTime;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for MongoClusters CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/MongoClusters_CreatePITR_CMK.json
     */
    /**
     * Sample code: Creates a Mongo Cluster resource with Customer Managed Key encryption from a point in time restore.
     * 
     * @param manager Entry point to MongoClusterManager.
     */
    public static void createsAMongoClusterResourceWithCustomerManagedKeyEncryptionFromAPointInTimeRestore(
        com.azure.resourcemanager.mongocluster.MongoClusterManager manager) {
        manager.mongoClusters().define("myMongoCluster").withRegion("westus2")
            .withExistingResourceGroup("TestResourceGroup")
            .withProperties(new MongoClusterProperties().withCreateMode(CreateMode.POINT_IN_TIME_RESTORE)
                .withRestoreParameters(new MongoClusterRestoreParameters()
                    .withPointInTimeUTC(OffsetDateTime.parse("2023-01-13T20:07:35Z")).withSourceResourceId(
                        "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestResourceGroup/providers/Microsoft.DocumentDB/mongoClusters/myOtherMongoCluster"))
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
