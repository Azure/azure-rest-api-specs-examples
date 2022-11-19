import com.azure.core.util.Context;

/** Samples for ImageVersions ListByImage. */
public final class Main {
    /*
     * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-10-12-preview/examples/ImageVersions_List.json
     */
    /**
     * Sample code: ImageVersions_ListByImage.
     *
     * @param manager Entry point to DevCenterManager.
     */
    public static void imageVersionsListByImage(com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager.imageVersions().listByImage("rg1", "Contoso", "DefaultDevGallery", "Win11", Context.NONE);
    }
}
