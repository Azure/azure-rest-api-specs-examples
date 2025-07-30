
/**
 * Samples for NewRegionFrontloadRelease Stop.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2024-09-01/examples/
     * NewRegionFrontloadRelease_Stop.json
     */
    /**
     * Sample code: NewRegionFrontloadRelease_Stop.
     * 
     * @param manager Entry point to ProviderHubManager.
     */
    public static void newRegionFrontloadReleaseStop(com.azure.resourcemanager.providerhub.ProviderHubManager manager) {
        manager.newRegionFrontloadReleases().stopWithResponse("Microsoft.Contoso", "2020week10",
            com.azure.core.util.Context.NONE);
    }
}
