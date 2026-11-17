
/**
 * Samples for NetworkSecurityPerimeterConfigurations GetConfiguration.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/getNetworkSecurityPerimeterConfiguration.json
     */
    /**
     * Sample code: Retrieve the network security perimeter configuration of a vault.
     * 
     * @param manager Entry point to KeyVaultManager.
     */
    public static void retrieveTheNetworkSecurityPerimeterConfigurationOfAVault(
        com.azure.resourcemanager.keyvault.KeyVaultManager manager) {
        manager.serviceClient().getNetworkSecurityPerimeterConfigurations().getConfigurationWithResponse(
            "sample-resource-group", "sample-vault", "00000000-0000-0000-0000-000000000000.sampleAssociation",
            com.azure.core.util.Context.NONE);
    }
}
