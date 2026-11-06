
/**
 * Samples for NspScheduledQueryRule ListNSP.
 */
public final class Main {
    /*
     * x-ms-original-file: 2021-10-01/NSPForScheduledQueryRule_List.json
     */
    /**
     * Sample code: List NSP configs by Scheduled Query Rule.
     * 
     * @param manager Entry point to MonitorManager.
     */
    public static void listNSPConfigsByScheduledQueryRule(com.azure.resourcemanager.monitor.MonitorManager manager) {
        manager.serviceClient().getNspScheduledQueryRules().listNSP("exampleRG", "someRule",
            com.azure.core.util.Context.NONE);
    }
}
