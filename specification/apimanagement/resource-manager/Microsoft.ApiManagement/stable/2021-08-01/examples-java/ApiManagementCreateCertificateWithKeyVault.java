import com.azure.resourcemanager.apimanagement.models.KeyVaultContractCreateProperties;

/** Samples for Certificate CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementCreateCertificateWithKeyVault.json
     */
    /**
     * Sample code: ApiManagementCreateCertificateWithKeyVault.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementCreateCertificateWithKeyVault(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .certificates()
            .define("templateCertkv")
            .withExistingService("rg1", "apimService1")
            .withKeyVault(
                new KeyVaultContractCreateProperties()
                    .withSecretIdentifier(
                        "https://rpbvtkeyvaultintegration.vault-int.azure-int.net/secrets/msitestingCert")
                    .withIdentityClientId("ceaa6b06-c00f-43ef-99ac-f53d1fe876a0"))
            .create();
    }
}
