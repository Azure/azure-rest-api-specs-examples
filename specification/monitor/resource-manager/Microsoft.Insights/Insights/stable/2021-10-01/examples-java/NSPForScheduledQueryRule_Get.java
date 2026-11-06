
/**
 * Samples for NspScheduledQueryRule GetNSP.
 */
public final class Main {
    /*
     * x-ms-original-file: 2021-10-01/NSPForScheduledQueryRule_Get.json
     */
    /**
     * Sample code: Get NSP config by name for Scheduled Query Rule.
     * 
     * @param manager Entry point to MonitorManager.
     */
    public static void
        getNSPConfigByNameForScheduledQueryRule(com.azure.resourcemanager.monitor.MonitorManager manager) {
        manager.serviceClient().getNspScheduledQueryRules().getNSPWithResponse("exampleRG", "someRule",
            "somePerimeterConfiguration", com.azure.core.util.Context.NONE);
    }
}
