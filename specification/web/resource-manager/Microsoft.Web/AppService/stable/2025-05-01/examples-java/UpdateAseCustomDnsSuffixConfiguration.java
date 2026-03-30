
import com.azure.resourcemanager.appservice.fluent.models.CustomDnsSuffixConfigurationInner;

/**
 * Samples for AppServiceEnvironments UpdateAseCustomDnsSuffixConfiguration.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/UpdateAseCustomDnsSuffixConfiguration.json
     */
    /**
     * Sample code: Update ASE custom DNS suffix configuration.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void
        updateASECustomDNSSuffixConfiguration(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getAppServiceEnvironments().updateAseCustomDnsSuffixConfigurationWithResponse("test-rg",
            "test-ase",
            new CustomDnsSuffixConfigurationInner().withDnsSuffix("contoso.com")
                .withCertificateUrl("https://test-kv.vault.azure.net/secrets/contosocert")
                .withKeyVaultReferenceIdentity("fakeTokenPlaceholder"),
            com.azure.core.util.Context.NONE);
    }
}
