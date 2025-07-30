
/**
 * Samples for NewRegionFrontloadRelease Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2024-09-01/examples/
     * NewRegionFrontloadRelease_Get.json
     */
    /**
     * Sample code: NewRegionFrontloadRelease_Get.
     * 
     * @param manager Entry point to ProviderHubManager.
     */
    public static void newRegionFrontloadReleaseGet(com.azure.resourcemanager.providerhub.ProviderHubManager manager) {
        manager.newRegionFrontloadReleases().getWithResponse("Microsoft.Contoso", "2020week10",
            com.azure.core.util.Context.NONE);
    }
}
