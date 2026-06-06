
/**
 * Samples for Profiles Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-04-01-preview/Profile-DELETE.json
     */
    /**
     * Sample code: Profile-DELETE.
     * 
     * @param manager Entry point to TrafficManager.
     */
    public static void profileDELETE(com.azure.resourcemanager.trafficmanager.TrafficManager manager) {
        manager.serviceClient().getProfiles().deleteWithResponse("azuresdkfornetautoresttrafficmanager1323",
            "azuresdkfornetautoresttrafficmanager3880", com.azure.core.util.Context.NONE);
    }
}
