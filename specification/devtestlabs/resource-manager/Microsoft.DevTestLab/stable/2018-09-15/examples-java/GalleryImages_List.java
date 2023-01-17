/** Samples for GalleryImages List. */
public final class Main {
    /*
     * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/GalleryImages_List.json
     */
    /**
     * Sample code: GalleryImages_List.
     *
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void galleryImagesList(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager
            .galleryImages()
            .list("resourceGroupName", "{labName}", null, null, null, null, com.azure.core.util.Context.NONE);
    }
}
