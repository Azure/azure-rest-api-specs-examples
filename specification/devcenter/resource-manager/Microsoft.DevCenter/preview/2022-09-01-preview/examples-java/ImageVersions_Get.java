import com.azure.core.util.Context;

/** Samples for ImageVersions Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-09-01-preview/examples/ImageVersions_Get.json
     */
    /**
     * Sample code: Versions_Get.
     *
     * @param manager Entry point to DevCenterManager.
     */
    public static void versionsGet(com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager
            .imageVersions()
            .getWithResponse("rg1", "Contoso", "DefaultDevGallery", "Win11", "{versionName}", Context.NONE);
    }
}
