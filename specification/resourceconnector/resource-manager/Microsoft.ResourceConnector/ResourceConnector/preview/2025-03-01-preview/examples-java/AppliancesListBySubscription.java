
/**
 * Samples for Appliances List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-01-preview/AppliancesListBySubscription.json
     */
    /**
     * Sample code: List Appliances by subscription.
     * 
     * @param manager Entry point to ResourceConnectorManager.
     */
    public static void
        listAppliancesBySubscription(com.azure.resourcemanager.resourceconnector.ResourceConnectorManager manager) {
        manager.appliances().list(com.azure.core.util.Context.NONE);
    }
}
