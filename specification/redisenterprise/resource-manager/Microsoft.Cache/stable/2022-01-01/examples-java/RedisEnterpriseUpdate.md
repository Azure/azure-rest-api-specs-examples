Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-redisenterprise_1.1.0-beta.1/sdk/redisenterprise/azure-resourcemanager-redisenterprise/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.redisenterprise.models.Cluster;
import com.azure.resourcemanager.redisenterprise.models.Sku;
import com.azure.resourcemanager.redisenterprise.models.SkuName;
import com.azure.resourcemanager.redisenterprise.models.TlsVersion;
import java.util.HashMap;
import java.util.Map;

/** Samples for RedisEnterprise Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/redisenterprise/resource-manager/Microsoft.Cache/stable/2022-01-01/examples/RedisEnterpriseUpdate.json
     */
    /**
     * Sample code: RedisEnterpriseUpdate.
     *
     * @param manager Entry point to RedisEnterpriseManager.
     */
    public static void redisEnterpriseUpdate(com.azure.resourcemanager.redisenterprise.RedisEnterpriseManager manager) {
        Cluster resource =
            manager.redisEnterprises().getByResourceGroupWithResponse("rg1", "cache1", Context.NONE).getValue();
        resource
            .update()
            .withTags(mapOf("tag1", "value1"))
            .withSku(new Sku().withName(SkuName.ENTERPRISE_FLASH_F300).withCapacity(9))
            .withMinimumTlsVersion(TlsVersion.ONE_TWO)
            .apply();
    }

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
```
