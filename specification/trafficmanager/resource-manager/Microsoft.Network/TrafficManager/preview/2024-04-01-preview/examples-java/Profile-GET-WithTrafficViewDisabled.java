
/**
 * Samples for Profiles GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-04-01-preview/Profile-GET-WithTrafficViewDisabled.json
     */
    /**
     * Sample code: Profile-GET-WithTrafficViewDisabled.
     * 
     * @param manager Entry point to TrafficManager.
     */
    public static void
        profileGETWithTrafficViewDisabled(com.azure.resourcemanager.trafficmanager.TrafficManager manager) {
        manager.serviceClient().getProfiles().getByResourceGroupWithResponse("azuresdkfornetautoresttrafficmanager1323",
            "azuresdkfornetautoresttrafficmanager3880", com.azure.core.util.Context.NONE);
    }
}
