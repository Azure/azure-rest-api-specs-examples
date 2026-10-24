
/**
 * Samples for Capabilities Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/Capabilities_Get.json
     */
    /**
     * Sample code: Get HorizonDB capabilities by subscription and region.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void
        getHorizonDBCapabilitiesBySubscriptionAndRegion(com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        manager.capabilities().getWithResponse("eastus2", com.azure.core.util.Context.NONE);
    }
}
