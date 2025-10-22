
import com.azure.resourcemanager.mongocluster.models.CustomerManagedKeyEncryptionProperties;
import com.azure.resourcemanager.mongocluster.models.EncryptionProperties;
import com.azure.resourcemanager.mongocluster.models.KeyEncryptionKeyIdentity;
import com.azure.resourcemanager.mongocluster.models.KeyEncryptionKeyIdentityType;
import com.azure.resourcemanager.mongocluster.models.ManagedServiceIdentity;
import com.azure.resourcemanager.mongocluster.models.ManagedServiceIdentityType;
import com.azure.resourcemanager.mongocluster.models.MongoCluster;
import com.azure.resourcemanager.mongocluster.models.MongoClusterUpdateProperties;
import com.azure.resourcemanager.mongocluster.models.UserAssignedIdentity;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for MongoClusters Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/MongoClusters_PatchCMK.json
     */
    /**
     * Sample code: Updates the customer managed encryption key on a mongo cluster resource.
     * 
     * @param manager Entry point to MongoClusterManager.
     */
    public static void updatesTheCustomerManagedEncryptionKeyOnAMongoClusterResource(
        com.azure.resourcemanager.mongocluster.MongoClusterManager manager) {
        MongoCluster resource = manager.mongoClusters()
            .getByResourceGroupWithResponse("TestResourceGroup", "myMongoCluster", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update().withIdentity(new ManagedServiceIdentity().withType(ManagedServiceIdentityType.USER_ASSIGNED)
            .withUserAssignedIdentities(mapOf(
                "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestResourceGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/myidentity2",
                new UserAssignedIdentity())))
            .withProperties(new MongoClusterUpdateProperties().withEncryption(
                new EncryptionProperties().withCustomerManagedKeyEncryption(new CustomerManagedKeyEncryptionProperties()
                    .withKeyEncryptionKeyIdentity(new KeyEncryptionKeyIdentity()
                        .withIdentityType(KeyEncryptionKeyIdentityType.USER_ASSIGNED_IDENTITY)
                        .withUserAssignedIdentityResourceId(
                            "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestResourceGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/myidentity2"))
                    .withKeyEncryptionKeyUrl("fakeTokenPlaceholder"))))
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
