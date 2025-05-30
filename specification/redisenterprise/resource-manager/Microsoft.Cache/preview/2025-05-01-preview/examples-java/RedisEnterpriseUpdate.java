
import com.azure.resourcemanager.redisenterprise.models.Cluster;
import com.azure.resourcemanager.redisenterprise.models.Sku;
import com.azure.resourcemanager.redisenterprise.models.SkuName;
import com.azure.resourcemanager.redisenterprise.models.TlsVersion;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for RedisEnterprise Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/redisenterprise/resource-manager/Microsoft.Cache/preview/2025-05-01-preview/examples/
     * RedisEnterpriseUpdate.json
     */
    /**
     * Sample code: RedisEnterpriseUpdate.
     * 
     * @param manager Entry point to RedisEnterpriseManager.
     */
    public static void redisEnterpriseUpdate(com.azure.resourcemanager.redisenterprise.RedisEnterpriseManager manager) {
        Cluster resource = manager.redisEnterprises()
            .getByResourceGroupWithResponse("rg1", "cache1", com.azure.core.util.Context.NONE).getValue();
        resource.update().withTags(mapOf("tag1", "value1"))
            .withSku(new Sku().withName(SkuName.ENTERPRISE_FLASH_F300).withCapacity(9))
            .withMinimumTlsVersion(TlsVersion.ONE_TWO).apply();
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
