/** Samples for ImageVersions ListByImage. */
public final class Main {
    /*
     * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/stable/2023-04-01/examples/ImageVersions_List.json
     */
    /**
     * Sample code: ImageVersions_ListByImage.
     *
     * @param manager Entry point to DevCenterManager.
     */
    public static void imageVersionsListByImage(com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager
            .imageVersions()
            .listByImage("rg1", "Contoso", "DefaultDevGallery", "Win11", com.azure.core.util.Context.NONE);
    }
}
