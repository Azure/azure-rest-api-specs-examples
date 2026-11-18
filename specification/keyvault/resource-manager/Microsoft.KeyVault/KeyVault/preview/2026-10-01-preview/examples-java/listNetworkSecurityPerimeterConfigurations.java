
/**
 * Samples for NetworkSecurityPerimeterConfigurations ListConfigurations.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/listNetworkSecurityPerimeterConfigurations.json
     */
    /**
     * Sample code: List the network security perimeter configurations of a vault.
     * 
     * @param manager Entry point to KeyVaultManager.
     */
    public static void listTheNetworkSecurityPerimeterConfigurationsOfAVault(
        com.azure.resourcemanager.keyvault.KeyVaultManager manager) {
        manager.serviceClient().getNetworkSecurityPerimeterConfigurations().listConfigurations("sample-resource-group",
            "sample-vault", com.azure.core.util.Context.NONE);
    }
}
