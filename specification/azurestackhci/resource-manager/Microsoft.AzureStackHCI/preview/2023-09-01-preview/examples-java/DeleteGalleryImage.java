/** Samples for GalleryImagesOperation Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/preview/2023-09-01-preview/examples/DeleteGalleryImage.json
     */
    /**
     * Sample code: DeleteGalleryImage.
     *
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void deleteGalleryImage(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.galleryImagesOperations().delete("test-rg", "test-gallery-image", com.azure.core.util.Context.NONE);
    }
}
