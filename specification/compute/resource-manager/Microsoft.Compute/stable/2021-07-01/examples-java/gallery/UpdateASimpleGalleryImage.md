Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.11.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.compute.models.GalleryImageIdentifier;
import com.azure.resourcemanager.compute.models.GalleryImageUpdate;
import com.azure.resourcemanager.compute.models.HyperVGeneration;
import com.azure.resourcemanager.compute.models.OperatingSystemStateTypes;
import com.azure.resourcemanager.compute.models.OperatingSystemTypes;

/** Samples for GalleryImages Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-07-01/examples/gallery/UpdateASimpleGalleryImage.json
     */
    /**
     * Sample code: Update a simple gallery image.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void updateASimpleGalleryImage(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getGalleryImages()
            .update(
                "myResourceGroup",
                "myGalleryName",
                "myGalleryImageName",
                new GalleryImageUpdate()
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
