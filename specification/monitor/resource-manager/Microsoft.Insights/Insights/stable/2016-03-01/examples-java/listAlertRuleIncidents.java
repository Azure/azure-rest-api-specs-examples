
/**
 * Samples for AlertRuleIncidents ListByAlertRule.
 */
public final class Main {
    /*
     * x-ms-original-file: 2016-03-01/listAlertRuleIncidents.json
     */
    /**
     * Sample code: List alert rule incidents.
     * 
     * @param manager Entry point to MonitorManager.
     */
    public static void listAlertRuleIncidents(com.azure.resourcemanager.monitor.MonitorManager manager) {
        manager.serviceClient().getAlertRuleIncidents().listByAlertRule("Rac46PostSwapRG", "myRuleName",
            com.azure.core.util.Context.NONE);
    }
}
