```java
import com.azure.core.util.Context;

/** Samples for Origins Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/Origins_Get.json
     */
    /**
     * Sample code: Origins_Get.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void originsGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cdnProfiles()
            .manager()
            .serviceClient()
            .getOrigins()
            .getWithResponse("RG", "profile1", "endpoint1", "www-someDomain-net", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.
