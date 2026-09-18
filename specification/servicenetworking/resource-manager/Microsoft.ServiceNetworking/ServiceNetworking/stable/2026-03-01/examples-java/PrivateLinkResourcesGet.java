
/**
 * Samples for PrivateLinkResourcesInterface ListByTrafficController.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/PrivateLinkResourcesGet.json
     */
    /**
     * Sample code: Get Private Link Resources.
     * 
     * @param manager Entry point to TrafficControllerManager.
     */
    public static void
        getPrivateLinkResources(com.azure.resourcemanager.servicenetworking.TrafficControllerManager manager) {
        manager.privateLinkResourcesInterfaces().listByTrafficController("rg1", "tc1",
            com.azure.core.util.Context.NONE);
    }
}
