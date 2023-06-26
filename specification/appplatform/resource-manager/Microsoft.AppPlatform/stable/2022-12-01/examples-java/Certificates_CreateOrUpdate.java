import com.azure.core.util.Context;
import com.azure.resourcemanager.appplatform.fluent.models.CertificateResourceInner;
import com.azure.resourcemanager.appplatform.models.KeyVaultCertificateProperties;

/** Samples for Certificates CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-12-01/examples/Certificates_CreateOrUpdate.json
     */
    /**
     * Sample code: Certificates_CreateOrUpdate.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void certificatesCreateOrUpdate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .springServices()
            .manager()
            .serviceClient()
            .getCertificates()
            .createOrUpdate(
                "myResourceGroup",
                "myservice",
                "mycertificate",
                new CertificateResourceInner()
                    .withProperties(
                        new KeyVaultCertificateProperties()
                            .withVaultUri("https://myvault.vault.azure.net")
                            .withKeyVaultCertName("mycert")
                            .withCertVersion("08a219d06d874795a96db47e06fbb01e")),
                Context.NONE);
    }
}
