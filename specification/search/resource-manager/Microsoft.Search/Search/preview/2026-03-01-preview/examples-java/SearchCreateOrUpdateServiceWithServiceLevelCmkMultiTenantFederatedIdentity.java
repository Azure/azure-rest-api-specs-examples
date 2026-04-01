
import com.azure.resourcemanager.search.fluent.models.SearchServiceInner;
import com.azure.resourcemanager.search.models.ComputeType;
import com.azure.resourcemanager.search.models.DataUserAssignedIdentity;
import com.azure.resourcemanager.search.models.EncryptionWithCmk;
import com.azure.resourcemanager.search.models.HostingMode;
import com.azure.resourcemanager.search.models.Identity;
import com.azure.resourcemanager.search.models.IdentityType;
import com.azure.resourcemanager.search.models.SearchEncryptionWithCmk;
import com.azure.resourcemanager.search.models.SearchResourceEncryptionKey;
import com.azure.resourcemanager.search.models.Sku;
import com.azure.resourcemanager.search.models.SkuName;
import com.azure.resourcemanager.search.models.UserAssignedIdentity;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Services CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * 2026-03-01-preview/SearchCreateOrUpdateServiceWithServiceLevelCmkMultiTenantFederatedIdentity.json
     */
    /**
     * Sample code: SearchCreateOrUpdateServiceWithServiceLevelCmkMultiTenantFederatedIdentity.
     * 
     * @param manager Entry point to SearchServiceManager.
     */
    public static void searchCreateOrUpdateServiceWithServiceLevelCmkMultiTenantFederatedIdentity(
        com.azure.resourcemanager.search.SearchServiceManager manager) {
        manager.serviceClient().getServices().createOrUpdate("rg1", "mysearchservice", new SearchServiceInner()
            .withLocation("westus").withTags(mapOf("app-name", "My e-commerce app"))
            .withSku(new Sku().withName(SkuName.STANDARD))
            .withIdentity(new Identity().withType(IdentityType.SYSTEM_ASSIGNED_USER_ASSIGNED)
                .withUserAssignedIdentities(mapOf(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/rg1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/user-mi",
                    new UserAssignedIdentity())))
            .withReplicaCount(3).withPartitionCount(1).withHostingMode(HostingMode.DEFAULT)
            .withComputeType(ComputeType.DEFAULT)
            .withEncryptionWithCmk(new EncryptionWithCmk().withEnforcement(SearchEncryptionWithCmk.ENABLED)
                .withServiceLevelEncryptionKey(new SearchResourceEncryptionKey().withKeyName("fakeTokenPlaceholder")
                    .withKeyVersion("fakeTokenPlaceholder").withVaultUri("fakeTokenPlaceholder")
                    .withIdentity(new DataUserAssignedIdentity().withUserAssignedIdentity(
                        "/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/rg1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/user-mi")
                        .withFederatedIdentityClientId("f83c6b1b-4d34-47e4-bb34-9d83df58b540")))),
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
