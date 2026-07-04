
/**
 * Samples for VpnServerConfigurations ListRadiusSecrets.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/AllVpnServerConfigurationRadiusServerSecretsList.json
     */
    /**
     * Sample code: ListAllVpnServerConfigurationRadiusServerSecrets.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        listAllVpnServerConfigurationRadiusServerSecrets(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVpnServerConfigurations().listRadiusSecretsWithResponse("rg1", "vpnserverconfig",
            com.azure.core.util.Context.NONE);
    }
}
