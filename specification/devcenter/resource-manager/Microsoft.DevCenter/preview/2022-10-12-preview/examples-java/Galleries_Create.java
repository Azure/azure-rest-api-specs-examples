/** Samples for Galleries CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-10-12-preview/examples/Galleries_Create.json
     */
    /**
     * Sample code: Galleries_CreateOrUpdate.
     *
     * @param manager Entry point to DevCenterManager.
     */
    public static void galleriesCreateOrUpdate(com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager
            .galleries()
            .define("{galleryName}")
            .withExistingDevcenter("rg1", "Contoso")
            .withGalleryResourceId(
                "/subscriptions/{subscriptionId}/resourceGroups/rg1/providers/Microsoft.Compute/galleries/{galleryName}")
            .create();
    }
}
