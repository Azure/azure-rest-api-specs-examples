Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for GalleryImages Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-10-01/examples/gallery/GetAGalleryImage.json
     */
    /**
     * Sample code: Get a gallery image.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAGalleryImage(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getGalleryImages()
            .getWithResponse("myResourceGroup", "myGalleryName", "myGalleryImageName", Context.NONE);
    }
}
```
