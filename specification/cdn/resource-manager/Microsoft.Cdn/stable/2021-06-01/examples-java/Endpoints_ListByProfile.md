Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for Endpoints ListByProfile. */
public final class Main {
    /*
     * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/Endpoints_ListByProfile.json
     */
    /**
     * Sample code: Endpoints_ListByProfile.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void endpointsListByProfile(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cdnProfiles().manager().serviceClient().getEndpoints().listByProfile("RG", "profile1", Context.NONE);
    }
}
```
