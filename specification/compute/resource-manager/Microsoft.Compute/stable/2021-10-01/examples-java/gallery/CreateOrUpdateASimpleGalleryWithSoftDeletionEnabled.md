Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.compute.fluent.models.GalleryInner;
import com.azure.resourcemanager.compute.models.SoftDeletePolicy;

/** Samples for Galleries CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-10-01/examples/gallery/CreateOrUpdateASimpleGalleryWithSoftDeletionEnabled.json
     */
    /**
     * Sample code: Create or update a simple gallery with soft deletion enabled.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createOrUpdateASimpleGalleryWithSoftDeletionEnabled(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getGalleries()
            .createOrUpdate(
                "myResourceGroup",
                "myGalleryName",
                new GalleryInner()
                    .withLocation("West US")
                    .withDescription("This is the gallery description.")
                    .withSoftDeletePolicy(new SoftDeletePolicy().withIsSoftDeleteEnabled(true)),
                Context.NONE);
    }
}
```
