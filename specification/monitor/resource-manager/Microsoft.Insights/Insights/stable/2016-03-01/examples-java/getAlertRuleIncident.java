
/**
 * Samples for AlertRuleIncidents Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2016-03-01/getAlertRuleIncident.json
     */
    /**
     * Sample code: Get a single alert rule incident.
     * 
     * @param manager Entry point to MonitorManager.
     */
    public static void getASingleAlertRuleIncident(com.azure.resourcemanager.monitor.MonitorManager manager) {
        manager.serviceClient().getAlertRuleIncidents().getWithResponse("Rac46PostSwapRG", "myRuleName",
            "Website_started", com.azure.core.util.Context.NONE);
    }
}
