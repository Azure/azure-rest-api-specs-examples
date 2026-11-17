
import com.azure.resourcemanager.keyvault.models.AccessPolicyEntry;
import com.azure.resourcemanager.keyvault.models.CertificatePermissions;
import com.azure.resourcemanager.keyvault.models.KeyPermissions;
import com.azure.resourcemanager.keyvault.models.Permissions;
import com.azure.resourcemanager.keyvault.models.SecretPermissions;
import com.azure.resourcemanager.keyvault.models.Sku;
import com.azure.resourcemanager.keyvault.models.SkuFamily;
import com.azure.resourcemanager.keyvault.models.SkuName;
import com.azure.resourcemanager.keyvault.models.VaultPatchParameters;
import com.azure.resourcemanager.keyvault.models.VaultPatchProperties;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;
import java.util.UUID;

/**
 * Samples for Vaults Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/updateVault.json
     */
    /**
     * Sample code: Update an existing vault.
     * 
     * @param manager Entry point to KeyVaultManager.
     */
    public static void updateAnExistingVault(com.azure.resourcemanager.keyvault.KeyVaultManager manager) {
        manager.serviceClient().getVaults().updateWithResponse("sample-resource-group", "sample-vault",
            new VaultPatchParameters().withProperties(new VaultPatchProperties()
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

    // Use "Map.of" if available
    @SuppressWarnings("unchecked")
    private static <T> Map<String, T> mapOf(Object... inputs) {
        Map<String, T> map = new HashMap<>();
        for (int i = 0; i < inputs.length; i += 2) {
            String key = (String) inputs[i];
            T value = (T) inputs[i + 1];
            map.put(key, value);
        }
        return map;
    }
}
