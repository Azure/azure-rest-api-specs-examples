
/**
 * Samples for NetworkSecurityPerimeterConfigurations Reconcile.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/NetworkSecurityPerimeterConfigurationReconcile.json
     */
    /**
     * Sample code: NetworkSecurityPerimeterConfigurationList.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void
        networkSecurityPerimeterConfigurationList(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getNetworkSecurityPerimeterConfigurations().reconcile("res4410", "sto8607",
            "dbedb4e0-40e6-4145-81f3-f1314c150774.resourceAssociation1", com.azure.core.util.Context.NONE);
    }
}
