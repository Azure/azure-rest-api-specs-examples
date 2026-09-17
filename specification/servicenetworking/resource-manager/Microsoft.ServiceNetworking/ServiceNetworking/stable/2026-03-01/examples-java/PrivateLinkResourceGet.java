
/**
 * Samples for PrivateLinkResourcesInterface Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/PrivateLinkResourceGet.json
     */
    /**
     * Sample code: Get Private Link Resource.
     * 
     * @param manager Entry point to TrafficControllerManager.
     */
    public static void
        getPrivateLinkResource(com.azure.resourcemanager.servicenetworking.TrafficControllerManager manager) {
        manager.privateLinkResourcesInterfaces().getWithResponse("rg1", "tc1", "fe1", com.azure.core.util.Context.NONE);
    }
}
