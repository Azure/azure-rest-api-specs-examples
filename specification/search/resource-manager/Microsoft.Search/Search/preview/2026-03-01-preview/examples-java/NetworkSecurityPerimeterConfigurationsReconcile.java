
/**
 * Samples for NetworkSecurityPerimeterConfigurations Reconcile.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/NetworkSecurityPerimeterConfigurationsReconcile.json
     */
    /**
     * Sample code: Reconcile NSP Config.
     * 
     * @param manager Entry point to SearchServiceManager.
     */
    public static void reconcileNSPConfig(com.azure.resourcemanager.search.SearchServiceManager manager) {
        manager.serviceClient().getNetworkSecurityPerimeterConfigurations().reconcile("rg1", "mysearchservice",
            "00000001-2222-3333-4444-111144444444.assoc1", com.azure.core.util.Context.NONE);
    }
}
