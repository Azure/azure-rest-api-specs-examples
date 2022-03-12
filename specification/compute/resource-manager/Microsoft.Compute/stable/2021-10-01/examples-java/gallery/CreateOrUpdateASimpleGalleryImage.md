Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.13.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.compute.fluent.models.GalleryImageInner;
import com.azure.resourcemanager.compute.models.GalleryImageIdentifier;
import com.azure.resourcemanager.compute.models.HyperVGeneration;
import com.azure.resourcemanager.compute.models.OperatingSystemStateTypes;
import com.azure.resourcemanager.compute.models.OperatingSystemTypes;

/** Samples for GalleryImages CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-10-01/examples/gallery/CreateOrUpdateASimpleGalleryImage.json
     */
    /**
     * Sample code: Create or update a simple gallery image.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createOrUpdateASimpleGalleryImage(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getGalleryImages()
            .createOrUpdate(
                "myResourceGroup",
                "myGalleryName",
                "myGalleryImageName",
                new GalleryImageInner()
                    .withLocation("West US")
                    .withOsType(OperatingSystemTypes.WINDOWS)
                    .withOsState(OperatingSystemStateTypes.GENERALIZED)
                    .withHyperVGeneration(HyperVGeneration.V1)
                    .withIdentifier(
                        new GalleryImageIdentifier()
                            .withPublisher("myPublisherName")
                            .withOffer("myOfferName")
                            .withSku("mySkuName")),
                Context.NONE);
    }
}
```
