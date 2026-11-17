
/**
 * Samples for NetworkSecurityPerimeterConfigurations ReconcileConfiguration.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/reconcileNetworkSecurityPerimeter.json
     */
    /**
     * Sample code: Reconcile the network security perimeter configuration of a vault.
     * 
     * @param manager Entry point to KeyVaultManager.
     */
    public static void reconcileTheNetworkSecurityPerimeterConfigurationOfAVault(
        com.azure.resourcemanager.keyvault.KeyVaultManager manager) {
        manager.serviceClient().getNetworkSecurityPerimeterConfigurations().reconcileConfigurationWithResponse(
            "sample-resource-group", "sample-vault", "00000000-0000-0000-0000-000000000000.sampleAssociation",
            com.azure.core.util.Context.NONE);
    }
}
