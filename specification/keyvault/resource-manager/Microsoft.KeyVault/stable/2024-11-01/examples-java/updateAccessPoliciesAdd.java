
import com.azure.resourcemanager.keyvault.fluent.models.VaultAccessPolicyParametersInner;
import com.azure.resourcemanager.keyvault.models.AccessPolicyEntry;
import com.azure.resourcemanager.keyvault.models.AccessPolicyUpdateKind;
import com.azure.resourcemanager.keyvault.models.CertificatePermissions;
import com.azure.resourcemanager.keyvault.models.KeyPermissions;
import com.azure.resourcemanager.keyvault.models.Permissions;
import com.azure.resourcemanager.keyvault.models.SecretPermissions;
import com.azure.resourcemanager.keyvault.models.VaultAccessPolicyProperties;
import java.util.Arrays;
import java.util.UUID;

/**
 * Samples for Vaults UpdateAccessPolicy.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/keyvault/resource-manager/Microsoft.KeyVault/stable/2024-11-01/examples/updateAccessPoliciesAdd.
     * json
     */
    /**
     * Sample code: Add an access policy, or update an access policy with new permissions.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void addAnAccessPolicyOrUpdateAnAccessPolicyWithNewPermissions(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.vaults().manager().serviceClient().getVaults().updateAccessPolicyWithResponse("sample-group",
            "sample-vault", AccessPolicyUpdateKind.ADD,
            new VaultAccessPolicyParametersInner()
                .withProperties(new VaultAccessPolicyProperties().withAccessPolicies(Arrays.asList(
                    new AccessPolicyEntry().withTenantId(UUID.fromString("00000000-0000-0000-0000-000000000000"))
                        .withObjectId("00000000-0000-0000-0000-000000000000")
                        .withPermissions(new Permissions().withKeys(Arrays.asList(KeyPermissions.ENCRYPT))
                            .withSecrets(Arrays.asList(SecretPermissions.GET))
                            .withCertificates(Arrays.asList(CertificatePermissions.GET)))))),
            com.azure.core.util.Context.NONE);
    }
}
