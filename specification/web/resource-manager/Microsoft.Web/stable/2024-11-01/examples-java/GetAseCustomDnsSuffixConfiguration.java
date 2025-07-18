
/**
 * Samples for AppServiceEnvironments GetAseCustomDnsSuffixConfiguration.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.Web/stable/2024-11-01/examples/GetAseCustomDnsSuffixConfiguration.
     * json
     */
    /**
     * Sample code: Get ASE custom DNS suffix configuration.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getASECustomDNSSuffixConfiguration(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getAppServiceEnvironments()
            .getAseCustomDnsSuffixConfigurationWithResponse("test-rg", "test-ase", com.azure.core.util.Context.NONE);
    }
}
