Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-datafactory_1.0.0-beta.7/sdk/datafactory/azure-resourcemanager-datafactory/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.resourcemanager.datafactory.models.ManagedPrivateEndpoint;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/** Samples for ManagedPrivateEndpoints CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/ManagedPrivateEndpoints_Create.json
     */
    /**
     * Sample code: ManagedVirtualNetworks_Create.
     *
     * @param manager Entry point to DataFactoryManager.
     */
    public static void managedVirtualNetworksCreate(com.azure.resourcemanager.datafactory.DataFactoryManager manager) {
        manager
            .managedPrivateEndpoints()
            .define("exampleManagedPrivateEndpointName")
            .withExistingManagedVirtualNetwork(
                "exampleResourceGroup", "exampleFactoryName", "exampleManagedVirtualNetworkName")
            .withProperties(
                new ManagedPrivateEndpoint()
                    .withFqdns(Arrays.asList())
                    .withGroupId("blob")
                    .withPrivateLinkResourceId(
                        "/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.Storage/storageAccounts/exampleBlobStorage")
                    .withAdditionalProperties(mapOf()))
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
