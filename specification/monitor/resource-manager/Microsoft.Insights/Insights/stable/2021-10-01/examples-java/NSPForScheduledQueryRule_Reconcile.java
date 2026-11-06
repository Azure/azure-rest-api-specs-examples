
/**
 * Samples for NspScheduledQueryRule ReconcileNSP.
 */
public final class Main {
    /*
     * x-ms-original-file: 2021-10-01/NSPForScheduledQueryRule_Reconcile.json
     */
    /**
     * Sample code: Reconcile NSP config for Scheduled Query Rule.
     * 
     * @param manager Entry point to MonitorManager.
     */
    public static void
        reconcileNSPConfigForScheduledQueryRule(com.azure.resourcemanager.monitor.MonitorManager manager) {
        manager.serviceClient().getNspScheduledQueryRules().reconcileNSP("exampleRG", "someRule",
            "somePerimeterConfiguration", com.azure.core.util.Context.NONE);
    }
}
