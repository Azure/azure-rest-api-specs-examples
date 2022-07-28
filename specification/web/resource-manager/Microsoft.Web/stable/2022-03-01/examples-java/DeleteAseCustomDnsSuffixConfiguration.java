import com.azure.core.util.Context;

/** Samples for AppServiceEnvironments DeleteAseCustomDnsSuffixConfiguration. */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2022-03-01/examples/DeleteAseCustomDnsSuffixConfiguration.json
     */
    /**
     * Sample code: Delete ASE custom DNS suffix configuration.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteASECustomDNSSuffixConfiguration(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .webApps()
            .manager()
            .serviceClient()
            .getAppServiceEnvironments()
            .deleteAseCustomDnsSuffixConfigurationWithResponse("test-rg", "test-ase", Context.NONE);
    }
}
