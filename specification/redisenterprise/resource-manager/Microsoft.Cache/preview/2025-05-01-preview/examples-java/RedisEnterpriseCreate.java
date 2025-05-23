
import com.azure.resourcemanager.redisenterprise.models.ClusterPropertiesEncryption;
import com.azure.resourcemanager.redisenterprise.models.ClusterPropertiesEncryptionCustomerManagedKeyEncryption;
import com.azure.resourcemanager.redisenterprise.models.ClusterPropertiesEncryptionCustomerManagedKeyEncryptionKeyIdentity;
import com.azure.resourcemanager.redisenterprise.models.CmkIdentityType;
import com.azure.resourcemanager.redisenterprise.models.ManagedServiceIdentity;
import com.azure.resourcemanager.redisenterprise.models.ManagedServiceIdentityType;
import com.azure.resourcemanager.redisenterprise.models.Sku;
import com.azure.resourcemanager.redisenterprise.models.SkuName;
import com.azure.resourcemanager.redisenterprise.models.TlsVersion;
import com.azure.resourcemanager.redisenterprise.models.UserAssignedIdentity;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for RedisEnterprise Create.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/redisenterprise/resource-manager/Microsoft.Cache/preview/2025-05-01-preview/examples/
     * RedisEnterpriseCreate.json
     */
    /**
     * Sample code: RedisEnterpriseCreate.
     * 
     * @param manager Entry point to RedisEnterpriseManager.
     */
    public static void redisEnterpriseCreate(com.azure.resourcemanager.redisenterprise.RedisEnterpriseManager manager) {
        manager.redisEnterprises().define("cache1").withRegion("West US").withExistingResourceGroup("rg1")
            .withSku(new Sku().withName(SkuName.ENTERPRISE_FLASH_F300).withCapacity(3))
            .withTags(mapOf("tag1", "value1")).withZones(Arrays.asList("1", "2", "3"))
            .withIdentity(new ManagedServiceIdentity().withType(ManagedServiceIdentityType.USER_ASSIGNED)
                .withUserAssignedIdentities(mapOf(
                    "/subscriptions/your-subscription/resourceGroups/your-rg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/your-identity",
                    new UserAssignedIdentity())))
            .withMinimumTlsVersion(TlsVersion.ONE_TWO)
            .withEncryption(new ClusterPropertiesEncryption().withCustomerManagedKeyEncryption(
                new ClusterPropertiesEncryptionCustomerManagedKeyEncryption().withKeyEncryptionKeyIdentity(
                    new ClusterPropertiesEncryptionCustomerManagedKeyEncryptionKeyIdentity()
                        .withUserAssignedIdentityResourceId(
                            "/subscriptions/your-subscription/resourceGroups/your-rg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/your-identity")
                        .withIdentityType(CmkIdentityType.USER_ASSIGNED_IDENTITY))
                    .withKeyEncryptionKeyUrl("fakeTokenPlaceholder")))
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
