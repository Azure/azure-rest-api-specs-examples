
/**
 * Samples for GalleryImages List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2018-09-15/GalleryImages_List.json
     */
    /**
     * Sample code: GalleryImages_List.
     * 
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void galleryImagesList(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager.galleryImages().list("resourceGroupName", "{labName}", null, null, null, null,
            com.azure.core.util.Context.NONE);
    }
}
