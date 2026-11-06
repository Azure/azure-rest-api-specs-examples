
/**
 * Samples for DataCollectionRules GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-03-11/DataCollectionRulesGet.json
     */
    /**
     * Sample code: Get data collection rule.
     * 
     * @param manager Entry point to MonitorManager.
     */
    public static void getDataCollectionRule(com.azure.resourcemanager.monitor.MonitorManager manager) {
        manager.serviceClient().getDataCollectionRules().getByResourceGroupWithResponse("myResourceGroup",
            "myCollectionRule", com.azure.core.util.Context.NONE);
    }
}
