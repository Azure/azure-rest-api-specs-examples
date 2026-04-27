
/**
 * Samples for OsImages Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/OsImages_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: OsImages_Get_MaximumSet.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void osImagesGetMaximumSet(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.osImages().getWithResponse("arowdcr", "10.2408.0.1", com.azure.core.util.Context.NONE);
    }
}
