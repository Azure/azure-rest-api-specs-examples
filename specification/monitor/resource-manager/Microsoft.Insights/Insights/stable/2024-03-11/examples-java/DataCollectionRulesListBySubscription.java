
/**
 * Samples for DataCollectionRules List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-03-11/DataCollectionRulesListBySubscription.json
     */
    /**
     * Sample code: List data collection rules by subscription.
     * 
     * @param manager Entry point to MonitorManager.
     */
    public static void listDataCollectionRulesBySubscription(com.azure.resourcemanager.monitor.MonitorManager manager) {
        manager.serviceClient().getDataCollectionRules().list(com.azure.core.util.Context.NONE);
    }
}
