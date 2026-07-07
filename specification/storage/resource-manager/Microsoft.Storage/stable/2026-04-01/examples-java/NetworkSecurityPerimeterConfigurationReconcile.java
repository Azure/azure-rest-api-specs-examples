
/**
 * Samples for NetworkSecurityPerimeterConfigurations Reconcile.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01/NetworkSecurityPerimeterConfigurationReconcile.json
     */
    /**
     * Sample code: NetworkSecurityPerimeterConfigurationReconcile.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void
        networkSecurityPerimeterConfigurationReconcile(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getNetworkSecurityPerimeterConfigurations().reconcile("res4410", "sto8607",
            "dbedb4e0-40e6-4145-81f3-f1314c150774.resourceAssociation1", com.azure.core.util.Context.NONE);
    }
}
