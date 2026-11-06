
/**
 * Samples for DataCollectionRuleAssociations ListByRule.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-03-11/DataCollectionRuleAssociationsListByRule.json
     */
    /**
     * Sample code: List associations for specified data collection rule.
     * 
     * @param manager Entry point to MonitorManager.
     */
    public static void
        listAssociationsForSpecifiedDataCollectionRule(com.azure.resourcemanager.monitor.MonitorManager manager) {
        manager.serviceClient().getDataCollectionRuleAssociations().listByRule("myResourceGroup", "myCollectionRule",
            null, null, com.azure.core.util.Context.NONE);
    }
}
