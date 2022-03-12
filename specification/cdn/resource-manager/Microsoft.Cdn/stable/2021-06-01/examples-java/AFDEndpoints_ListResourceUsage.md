Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.13.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for AfdEndpoints ListResourceUsage. */
public final class Main {
    /*
     * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/AFDEndpoints_ListResourceUsage.json
     */
    /**
     * Sample code: AFDEndpoints_ListResourceUsage.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void aFDEndpointsListResourceUsage(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cdnProfiles()
            .manager()
            .serviceClient()
            .getAfdEndpoints()
            .listResourceUsage("RG", "profile1", "endpoint1", Context.NONE);
    }
}
```
