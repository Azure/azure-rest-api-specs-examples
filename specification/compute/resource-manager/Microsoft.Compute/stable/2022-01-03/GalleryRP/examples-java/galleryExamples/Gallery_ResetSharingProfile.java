import com.azure.core.util.Context;
import com.azure.resourcemanager.compute.fluent.models.SharingUpdateInner;
import com.azure.resourcemanager.compute.models.SharingUpdateOperationTypes;

/** Samples for GallerySharingProfile Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-01-03/GalleryRP/examples/galleryExamples/Gallery_ResetSharingProfile.json
     */
    /**
     * Sample code: reset sharing profile of a gallery.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void resetSharingProfileOfAGallery(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getGallerySharingProfiles()
            .update(
                "myResourceGroup",
                "myGalleryName",
                new SharingUpdateInner().withOperationType(SharingUpdateOperationTypes.RESET),
                Context.NONE);
    }
}
