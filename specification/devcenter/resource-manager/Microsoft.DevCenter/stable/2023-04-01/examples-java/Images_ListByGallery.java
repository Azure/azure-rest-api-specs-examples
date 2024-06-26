/** Samples for Images ListByGallery. */
public final class Main {
    /*
     * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/stable/2023-04-01/examples/Images_ListByGallery.json
     */
    /**
     * Sample code: Images_ListByGallery.
     *
     * @param manager Entry point to DevCenterManager.
     */
    public static void imagesListByGallery(com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager.images().listByGallery("rg1", "Contoso", "DevGallery", null, com.azure.core.util.Context.NONE);
    }
}
