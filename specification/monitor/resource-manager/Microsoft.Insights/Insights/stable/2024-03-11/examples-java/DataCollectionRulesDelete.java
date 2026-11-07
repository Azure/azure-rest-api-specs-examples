
/**
 * Samples for DataCollectionRules Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-03-11/DataCollectionRulesDelete.json
     */
    /**
     * Sample code: Delete data collection rule.
     * 
     * @param manager Entry point to MonitorManager.
     */
    public static void deleteDataCollectionRule(com.azure.resourcemanager.monitor.MonitorManager manager) {
        manager.serviceClient().getDataCollectionRules().deleteWithResponse("myResourceGroup", "myCollectionRule", null,
            com.azure.core.util.Context.NONE);
    }
}
