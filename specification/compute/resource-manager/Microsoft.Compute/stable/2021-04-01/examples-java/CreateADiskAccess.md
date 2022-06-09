```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.compute.fluent.models.DiskAccessInner;

/** Samples for DiskAccesses CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-04-01/examples/CreateADiskAccess.json
     */
    /**
     * Sample code: Create a disk access resource.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createADiskAccessResource(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getDiskAccesses()
            .createOrUpdate(
                "myResourceGroup", "myDiskAccess", new DiskAccessInner().withLocation("West US"), Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.10.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.
