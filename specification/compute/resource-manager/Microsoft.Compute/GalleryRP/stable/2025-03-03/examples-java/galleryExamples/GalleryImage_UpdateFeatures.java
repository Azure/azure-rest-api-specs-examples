
import com.azure.resourcemanager.compute.models.GalleryImageFeature;
import com.azure.resourcemanager.compute.models.GalleryImageIdentifier;
import com.azure.resourcemanager.compute.models.GalleryImageUpdate;
import com.azure.resourcemanager.compute.models.HyperVGeneration;
import com.azure.resourcemanager.compute.models.OperatingSystemStateTypes;
import com.azure.resourcemanager.compute.models.OperatingSystemTypes;
import java.util.Arrays;

/**
 * Samples for GalleryImages Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2025-03-03/examples/galleryExamples/
     * GalleryImage_UpdateFeatures.json
     */
    /**
     * Sample code: Update a gallery image feature.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void updateAGalleryImageFeature(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getGalleryImages().update("myResourceGroup", "myGalleryName",
            "myGalleryImageName",
            new GalleryImageUpdate().withOsType(OperatingSystemTypes.WINDOWS)
                .withOsState(OperatingSystemStateTypes.GENERALIZED).withHyperVGeneration(HyperVGeneration.V2)
                .withIdentifier(new GalleryImageIdentifier().withPublisher("myPublisherName").withOffer("myOfferName")
                    .withSku("mySkuName"))
                .withFeatures(Arrays.asList(new GalleryImageFeature().withName("SecurityType")
                    .withValue("TrustedLaunch").withStartsAtVersion("2.0.0")))
                .withAllowUpdateImage(true),
            com.azure.core.util.Context.NONE);
    }
}
