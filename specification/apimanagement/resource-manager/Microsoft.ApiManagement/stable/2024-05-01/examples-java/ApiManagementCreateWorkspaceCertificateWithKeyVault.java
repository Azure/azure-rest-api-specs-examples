
import com.azure.resourcemanager.apimanagement.models.CertificateCreateOrUpdateParameters;
import com.azure.resourcemanager.apimanagement.models.KeyVaultContractCreateProperties;

/**
 * Samples for WorkspaceCertificate CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementCreateWorkspaceCertificateWithKeyVault.json
     */
    /**
     * Sample code: ApiManagementCreateWorkspaceCertificateWithKeyVault.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementCreateWorkspaceCertificateWithKeyVault(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceCertificates().createOrUpdateWithResponse("rg1", "apimService1", "wks1", "templateCertkv",
            new CertificateCreateOrUpdateParameters()
                .withKeyVault(new KeyVaultContractCreateProperties().withSecretIdentifier("fakeTokenPlaceholder")
                    .withIdentityClientId("ceaa6b06-c00f-43ef-99ac-f53d1fe876a0")),
            null, com.azure.core.util.Context.NONE);
    }
}
