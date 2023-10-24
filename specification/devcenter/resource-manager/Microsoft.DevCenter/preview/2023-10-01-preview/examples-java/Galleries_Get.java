/** Samples for Galleries Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2023-10-01-preview/examples/Galleries_Get.json
     */
    /**
     * Sample code: Galleries_Get.
     *
     * @param manager Entry point to DevCenterManager.
     */
    public static void galleriesGet(com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager.galleries().getWithResponse("rg1", "Contoso", "StandardGallery", com.azure.core.util.Context.NONE);
    }
}
