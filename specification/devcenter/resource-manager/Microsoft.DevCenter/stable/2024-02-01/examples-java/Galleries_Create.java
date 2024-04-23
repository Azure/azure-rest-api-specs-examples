
/**
 * Samples for Galleries CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/devcenter/resource-manager/Microsoft.DevCenter/stable/2024-02-01/examples/Galleries_Create.json
     */
    /**
     * Sample code: Galleries_CreateOrUpdate.
     * 
     * @param manager Entry point to DevCenterManager.
     */
    public static void galleriesCreateOrUpdate(com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager.galleries().define("StandardGallery").withExistingDevcenter("rg1", "Contoso").withGalleryResourceId(
            "/subscriptions/0ac520ee-14c0-480f-b6c9-0a90c58ffff/resourceGroups/rg1/providers/Microsoft.Compute/galleries/StandardGallery")
            .create();
    }
}
