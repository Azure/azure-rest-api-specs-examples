
/**
 * Samples for Profiles ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-04-01-preview/Profile-GET-ByResourceGroup.json
     */
    /**
     * Sample code: ListProfilesByResourceGroup.
     * 
     * @param manager Entry point to TrafficManager.
     */
    public static void listProfilesByResourceGroup(com.azure.resourcemanager.trafficmanager.TrafficManager manager) {
        manager.serviceClient().getProfiles().listByResourceGroup("azuresdkfornetautoresttrafficmanager3640",
            com.azure.core.util.Context.NONE);
    }
}
