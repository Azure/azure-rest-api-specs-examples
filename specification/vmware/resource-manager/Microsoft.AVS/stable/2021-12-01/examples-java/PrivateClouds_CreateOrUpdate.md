Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-avs_1.0.0-beta.3/sdk/avs/azure-resourcemanager-avs/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.resourcemanager.avs.models.ManagementCluster;
import com.azure.resourcemanager.avs.models.PrivateCloudIdentity;
import com.azure.resourcemanager.avs.models.ResourceIdentityType;
import com.azure.resourcemanager.avs.models.Sku;
import java.util.HashMap;
import java.util.Map;

/** Samples for PrivateClouds CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/PrivateClouds_CreateOrUpdate.json
     */
    /**
     * Sample code: PrivateClouds_CreateOrUpdate.
     *
     * @param manager Entry point to AvsManager.
     */
    public static void privateCloudsCreateOrUpdate(com.azure.resourcemanager.avs.AvsManager manager) {
        manager
            .privateClouds()
            .define("cloud1")
            .withRegion("eastus2")
            .withExistingResourceGroup("group1")
            .withSku(new Sku().withName("AV36"))
            .withTags(mapOf())
            .withIdentity(new PrivateCloudIdentity().withType(ResourceIdentityType.SYSTEM_ASSIGNED))
            .withNetworkBlock("192.168.48.0/22")
            .withManagementCluster(new ManagementCluster().withClusterSize(4))
            .create();
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
