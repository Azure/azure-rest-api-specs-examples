Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.13.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.compute.fluent.models.RestorePointCollectionInner;
import com.azure.resourcemanager.compute.models.RestorePointCollectionSourceProperties;
import java.util.HashMap;
import java.util.Map;

/** Samples for RestorePointCollections CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/CreateOrUpdateARestorePointCollectionForCrossRegionCopy.json
     */
    /**
     * Sample code: Create or update a restore point collection for cross region copy.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createOrUpdateARestorePointCollectionForCrossRegionCopy(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getRestorePointCollections()
            .createOrUpdateWithResponse(
                "myResourceGroup",
                "myRpc",
                new RestorePointCollectionInner()
                    .withLocation("norwayeast")
                    .withTags(mapOf("myTag1", "tagValue1"))
                    .withSource(
                        new RestorePointCollectionSourceProperties()
                            .withId(
                                "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/restorePointCollections/sourceRpcName")),
                Context.NONE);
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
