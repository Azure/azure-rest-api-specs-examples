
/**
 * Samples for ActionGroups ReconcileNSP.
 */
public final class Main {
    /*
     * x-ms-original-file: 2021-10-01/NSPForActionGroups_Reconcile.json
     */
    /**
     * Sample code: Reconcile NSP config by name for an action group.
     * 
     * @param manager Entry point to MonitorManager.
     */
    public static void
        reconcileNSPConfigByNameForAnActionGroup(com.azure.resourcemanager.monitor.MonitorManager manager) {
        manager.serviceClient().getActionGroups().reconcileNSP("exampleRG", "someActionGroup",
            "somePerimeterConfiguration", com.azure.core.util.Context.NONE);
    }
}
