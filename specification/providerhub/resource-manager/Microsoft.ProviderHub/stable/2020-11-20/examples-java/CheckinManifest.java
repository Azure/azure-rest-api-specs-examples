import com.azure.resourcemanager.providerhub.models.CheckinManifestParams;

/** Samples for ResourceProvider CheckinManifest. */
public final class Main {
    /*
     * x-ms-original-file: specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2020-11-20/examples/CheckinManifest.json
     */
    /**
     * Sample code: CheckinManifest.
     *
     * @param manager Entry point to ProviderHubManager.
     */
    public static void checkinManifest(com.azure.resourcemanager.providerhub.ProviderHubManager manager) {
        manager
            .resourceProviders()
            .checkinManifestWithResponse(
                "Microsoft.Contoso",
                new CheckinManifestParams().withEnvironment("Prod").withBaselineArmManifestLocation("EastUS2EUAP"),
                com.azure.core.util.Context.NONE);
    }
}
