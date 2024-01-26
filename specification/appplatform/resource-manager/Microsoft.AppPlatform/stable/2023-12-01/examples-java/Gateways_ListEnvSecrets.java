
/**
 * Samples for Gateways ListEnvSecrets.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2023-12-01/examples/
     * Gateways_ListEnvSecrets.json
     */
    /**
     * Sample code: Gateways_ListEnvSecrets.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void gatewaysListEnvSecrets(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.springServices().manager().serviceClient().getGateways().listEnvSecretsWithResponse("myResourceGroup",
            "myservice", "default", com.azure.core.util.Context.NONE);
    }
}
