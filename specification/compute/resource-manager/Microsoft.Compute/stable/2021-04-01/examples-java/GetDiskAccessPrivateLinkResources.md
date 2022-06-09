```java
import com.azure.core.util.Context;

/** Samples for DiskAccesses GetPrivateLinkResources. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-04-01/examples/GetDiskAccessPrivateLinkResources.json
     */
    /**
     * Sample code: List all possible private link resources under disk access resource.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listAllPossiblePrivateLinkResourcesUnderDiskAccessResource(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getDiskAccesses()
            .getPrivateLinkResourcesWithResponse("myResourceGroup", "myDiskAccess", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.10.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.
