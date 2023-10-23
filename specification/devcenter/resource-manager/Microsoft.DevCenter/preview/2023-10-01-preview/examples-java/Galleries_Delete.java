/** Samples for Galleries Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2023-10-01-preview/examples/Galleries_Delete.json
     */
    /**
     * Sample code: Galleries_Delete.
     *
     * @param manager Entry point to DevCenterManager.
     */
    public static void galleriesDelete(com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager.galleries().delete("rg1", "Contoso", "StandardGallery", com.azure.core.util.Context.NONE);
    }
}
