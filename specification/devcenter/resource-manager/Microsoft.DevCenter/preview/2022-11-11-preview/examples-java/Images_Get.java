import com.azure.core.util.Context;

/** Samples for Images Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-11-11-preview/examples/Images_Get.json
     */
    /**
     * Sample code: Images_Get.
     *
     * @param manager Entry point to DevCenterManager.
     */
    public static void imagesGet(com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager.images().getWithResponse("rg1", "Contoso", "DefaultDevGallery", "ContosoBaseImage", Context.NONE);
    }
}
