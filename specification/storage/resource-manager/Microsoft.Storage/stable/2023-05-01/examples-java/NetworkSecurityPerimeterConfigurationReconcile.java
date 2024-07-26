
/**
 * Samples for NetworkSecurityPerimeterConfigurations Reconcile.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2023-05-01/examples/
     * NetworkSecurityPerimeterConfigurationReconcile.json
     */
    /**
     * Sample code: NetworkSecurityPerimeterConfigurationReconcile.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        networkSecurityPerimeterConfigurationReconcile(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.storageAccounts().manager().serviceClient().getNetworkSecurityPerimeterConfigurations().reconcile(
            "res4410", "sto8607", "dbedb4e0-40e6-4145-81f3-f1314c150774.resourceAssociation1",
            com.azure.core.util.Context.NONE);
    }
}
