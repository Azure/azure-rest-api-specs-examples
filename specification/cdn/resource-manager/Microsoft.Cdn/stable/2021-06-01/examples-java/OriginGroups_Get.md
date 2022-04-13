Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.14.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for OriginGroups Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/OriginGroups_Get.json
     */
    /**
     * Sample code: OriginGroups_Get.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void originGroupsGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cdnProfiles()
            .manager()
            .serviceClient()
            .getOriginGroups()
            .getWithResponse("RG", "profile1", "endpoint1", "originGroup1", Context.NONE);
    }
}
```
