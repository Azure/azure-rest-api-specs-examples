
import com.azure.resourcemanager.containerregistry.fluent.models.CacheRuleInner;
import com.azure.resourcemanager.containerregistry.models.EntraFederatedIdentityCredentialAuthenticationProperties;
import com.azure.resourcemanager.containerregistry.models.IdentityProperties;
import com.azure.resourcemanager.containerregistry.models.ResourceIdentityType;
import com.azure.resourcemanager.containerregistry.models.UserIdentityProperties;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for CacheRules Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-15-preview/CacheRuleCreateUserAssignedMIEntraFederatedIdentityCredential.json
     */
    /**
     * Sample code: CacheRuleCreateUserAssignedMIAuthEntraFederatedIdentityCredential.
     * 
     * @param manager Entry point to ContainerRegistryManager.
     */
    public static void cacheRuleCreateUserAssignedMIAuthEntraFederatedIdentityCredential(
        com.azure.resourcemanager.containerregistry.ContainerRegistryManager manager) {
        manager.serviceClient().getCacheRules().create("myResourceGroup", "myRegistry", "myCacheRule",
            new CacheRuleInner().withIdentity(new IdentityProperties().withType(ResourceIdentityType.USER_ASSIGNED)
                .withUserAssignedIdentities(mapOf(
                    "/subscriptions/12345678-1234-1234-1234-123456789012/resourceGroups/myResourceGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/myUserAssignedIdentity",
                    new UserIdentityProperties())))
                .withAdditionalAuthenticationProperties(new EntraFederatedIdentityCredentialAuthenticationProperties()
                    .withSourceTenantId("72f988bf-86f1-41af-91ab-2d7cd011db47")
                    .withFederatedApplicationClientId("abcdef12-3456-7890-abcd-ef1234567890"))
                .withSourceRepository("myupstream.azurecr.io/hello-world")
                .withTargetRepository("cached-upstream/hello-world"),
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
