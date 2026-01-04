
import com.azure.resourcemanager.keyvault.models.AccessPolicyEntry;
import com.azure.resourcemanager.keyvault.models.CertificatePermissions;
import com.azure.resourcemanager.keyvault.models.KeyPermissions;
import com.azure.resourcemanager.keyvault.models.Permissions;
import com.azure.resourcemanager.keyvault.models.SecretPermissions;
import com.azure.resourcemanager.keyvault.models.Sku;
import com.azure.resourcemanager.keyvault.models.SkuFamily;
import com.azure.resourcemanager.keyvault.models.SkuName;
import com.azure.resourcemanager.keyvault.models.VaultCreateOrUpdateParameters;
import com.azure.resourcemanager.keyvault.models.VaultProperties;
import java.util.Arrays;
import java.util.UUID;

/**
 * Samples for Vaults CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/createVault.json
     */
    /**
     * Sample code: Create a new vault or update an existing vault.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createANewVaultOrUpdateAnExistingVault(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.vaults().manager().serviceClient().getVaults().createOrUpdate("sample-resource-group", "sample-vault",
            new VaultCreateOrUpdateParameters().withLocation("westus").withProperties(new VaultProperties()
                .withTenantId(UUID.fromString("00000000-0000-0000-0000-000000000000"))
                .withSku(new Sku().withFamily(SkuFamily.A).withName(SkuName.STANDARD))
                .withAccessPolicies(Arrays.asList(new AccessPolicyEntry()
                    .withTenantId(UUID.fromString("00000000-0000-0000-0000-000000000000"))
                    .withObjectId("00000000-0000-0000-0000-000000000000")
                    .withPermissions(new Permissions()
                        .withKeys(Arrays.asList(KeyPermissions.ENCRYPT, KeyPermissions.DECRYPT, KeyPermissions.WRAP_KEY,
                            KeyPermissions.UNWRAP_KEY, KeyPermissions.SIGN, KeyPermissions.VERIFY, KeyPermissions.GET,
                            KeyPermissions.LIST, KeyPermissions.CREATE, KeyPermissions.UPDATE, KeyPermissions.IMPORT,
                            KeyPermissions.DELETE, KeyPermissions.BACKUP, KeyPermissions.RESTORE,
                            KeyPermissions.RECOVER, KeyPermissions.PURGE))
                        .withSecrets(Arrays.asList(SecretPermissions.GET, SecretPermissions.LIST, SecretPermissions.SET,
                            SecretPermissions.DELETE, SecretPermissions.BACKUP, SecretPermissions.RESTORE,
                            SecretPermissions.RECOVER, SecretPermissions.PURGE))
                        .withCertificates(Arrays.asList(CertificatePermissions.GET, CertificatePermissions.LIST,
                            CertificatePermissions.DELETE, CertificatePermissions.CREATE, CertificatePermissions.IMPORT,
                            CertificatePermissions.UPDATE, CertificatePermissions.MANAGECONTACTS,
                            CertificatePermissions.GETISSUERS, CertificatePermissions.LISTISSUERS,
                            CertificatePermissions.SETISSUERS, CertificatePermissions.DELETEISSUERS,
                            CertificatePermissions.MANAGEISSUERS, CertificatePermissions.RECOVER,
                            CertificatePermissions.PURGE)))))
                .withEnabledForDeployment(true).withEnabledForDiskEncryption(true)
                .withEnabledForTemplateDeployment(true).withPublicNetworkAccess("Enabled")),
            com.azure.core.util.Context.NONE);
    }
}
