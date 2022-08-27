import com.azure.core.util.Context;
import com.azure.resourcemanager.keyvault.fluent.models.ManagedHsmInner;
import com.azure.resourcemanager.keyvault.models.ManagedHsmProperties;
import com.azure.resourcemanager.keyvault.models.ManagedHsmSku;
import com.azure.resourcemanager.keyvault.models.ManagedHsmSkuFamily;
import com.azure.resourcemanager.keyvault.models.ManagedHsmSkuName;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;
import java.util.UUID;

/** Samples for ManagedHsms CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/keyvault/resource-manager/Microsoft.KeyVault/stable/2021-10-01/examples/ManagedHsm_CreateOrUpdate.json
     */
    /**
     * Sample code: Create a new managed HSM Pool or update an existing managed HSM Pool.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createANewManagedHSMPoolOrUpdateAnExistingManagedHSMPool(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .vaults()
            .manager()
            .serviceClient()
            .getManagedHsms()
            .createOrUpdate(
                "hsm-group",
                "hsm1",
                new ManagedHsmInner()
                    .withLocation("westus")
                    .withTags(mapOf("Dept", "hsm", "Environment", "dogfood"))
                    .withSku(
                        new ManagedHsmSku().withFamily(ManagedHsmSkuFamily.B).withName(ManagedHsmSkuName.STANDARD_B1))
                    .withProperties(
                        new ManagedHsmProperties()
                            .withTenantId(UUID.fromString("00000000-0000-0000-0000-000000000000"))
                            .withInitialAdminObjectIds(Arrays.asList("00000000-0000-0000-0000-000000000000"))
                            .withEnableSoftDelete(true)
                            .withSoftDeleteRetentionInDays(90)
                            .withEnablePurgeProtection(true)),
                Context.NONE);
    }

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
