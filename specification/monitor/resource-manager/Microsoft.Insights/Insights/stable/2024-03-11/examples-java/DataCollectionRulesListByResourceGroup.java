
/**
 * Samples for DataCollectionRules ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-03-11/DataCollectionRulesListByResourceGroup.json
     */
    /**
     * Sample code: List data collection rules by resource group.
     * 
     * @param manager Entry point to MonitorManager.
     */
    public static void
        listDataCollectionRulesByResourceGroup(com.azure.resourcemanager.monitor.MonitorManager manager) {
        manager.serviceClient().getDataCollectionRules().listByResourceGroup("myResourceGroup",
            com.azure.core.util.Context.NONE);
    }
}
