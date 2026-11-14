
import com.azure.resourcemanager.containerregistry.models.CacheRuleUpdateParameters;
import com.azure.resourcemanager.containerregistry.models.GarAuthenticationUpdateProperties;
import com.azure.resourcemanager.containerregistry.models.IdentityProperties;
import com.azure.resourcemanager.containerregistry.models.ResourceIdentityType;
import com.azure.resourcemanager.containerregistry.models.UserIdentityProperties;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for CacheRules Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-15-preview/CacheRuleUpdateConfigureGoogleAuthentication.json
     */
    /**
     * Sample code: Configure Google authentication on an existing cache rule.
     * 
     * @param manager Entry point to ContainerRegistryManager.
     */
    public static void configureGoogleAuthenticationOnAnExistingCacheRule(
        com.azure.resourcemanager.containerregistry.ContainerRegistryManager manager) {
        manager.serviceClient().getCacheRules().update("myResourceGroup", "myRegistry", "myCacheRule",
            new CacheRuleUpdateParameters().withIdentity(new IdentityProperties()
                .withType(ResourceIdentityType.USER_ASSIGNED)
                .withUserAssignedIdentities(mapOf(
                    "/subscriptions/12345678-1234-1234-1234-123456789012/resourceGroups/myResourceGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/myUserAssignedIdentity",
                    new UserIdentityProperties())))
                .withAdditionalAuthenticationProperties(new GarAuthenticationUpdateProperties()
                    .withProjectNumber("123456789012").withWorkloadIdentityPool("my-workload-identity-pool")
                    .withWorkloadIdentityProvider("my-workload-identity-provider")),
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
