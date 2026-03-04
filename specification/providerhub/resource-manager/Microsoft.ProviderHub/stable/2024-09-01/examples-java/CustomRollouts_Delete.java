
/**
 * Samples for CustomRollouts Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/CustomRollouts_Delete.json
     */
    /**
     * Sample code: providerReleases_Delete.
     * 
     * @param manager Entry point to ProviderHubManager.
     */
    public static void providerReleasesDelete(com.azure.resourcemanager.providerhub.ProviderHubManager manager) {
        manager.customRollouts().deleteByResourceGroupWithResponse("Microsoft.Contoso", "2020week10",
            com.azure.core.util.Context.NONE);
    }
}
