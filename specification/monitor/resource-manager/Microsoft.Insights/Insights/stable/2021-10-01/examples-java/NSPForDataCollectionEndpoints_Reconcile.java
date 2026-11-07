
/**
 * Samples for DataCollectionEndpoints ReconcileNSP.
 */
public final class Main {
    /*
     * x-ms-original-file: 2021-10-01/NSPForDataCollectionEndpoints_Reconcile.json
     */
    /**
     * Sample code: Reconcile NSP config for data collection endpoint.
     * 
     * @param manager Entry point to MonitorManager.
     */
    public static void
        reconcileNSPConfigForDataCollectionEndpoint(com.azure.resourcemanager.monitor.MonitorManager manager) {
        manager.serviceClient().getDataCollectionEndpoints().reconcileNSP("exampleRG", "someDataCollectionEndpoint",
            "somePerimeterConfiguration", com.azure.core.util.Context.NONE);
    }
}
