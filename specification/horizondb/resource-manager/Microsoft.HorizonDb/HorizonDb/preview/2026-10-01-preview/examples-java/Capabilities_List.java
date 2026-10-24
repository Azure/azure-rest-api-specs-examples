
/**
 * Samples for Capabilities List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/Capabilities_List.json
     */
    /**
     * Sample code: List HorizonDB capabilities by subscription and region.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void
        listHorizonDBCapabilitiesBySubscriptionAndRegion(com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        manager.capabilities().list("eastus2", com.azure.core.util.Context.NONE);
    }
}
