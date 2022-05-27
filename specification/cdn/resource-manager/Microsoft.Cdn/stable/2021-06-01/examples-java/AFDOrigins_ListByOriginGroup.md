Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for AfdOrigins ListByOriginGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/AFDOrigins_ListByOriginGroup.json
     */
    /**
     * Sample code: AFDOrigins_ListByOriginGroup.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void aFDOriginsListByOriginGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cdnProfiles()
            .manager()
            .serviceClient()
            .getAfdOrigins()
            .listByOriginGroup("RG", "profile1", "origingroup1", Context.NONE);
    }
}
```
