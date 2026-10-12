
/**
 * Samples for Profiles GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-09-01/Profile-GET-WithEndpointsAndRecordType.json
     */
    /**
     * Sample code: Profile-GET-WithEndpointsAndRecordType.
     * 
     * @param manager Entry point to TrafficManager.
     */
    public static void
        profileGETWithEndpointsAndRecordType(com.azure.resourcemanager.trafficmanager.TrafficManager manager) {
        manager.serviceClient().getProfiles().getByResourceGroupWithResponse("azuresdkfornetautoresttrafficmanager1323",
            "azuresdkfornetautoresttrafficmanager3880", com.azure.core.util.Context.NONE);
    }
}
