
import com.azure.resourcemanager.containerregistry.fluent.models.CacheRuleInner;
import com.azure.resourcemanager.containerregistry.models.EcrAuthenticationProperties;
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
     * x-ms-original-file: 2026-10-15-preview/CacheRuleCreateUserAssignedMIAws.json
     */
    /**
     * Sample code: CacheRuleCreateUserAssignedMIAuthAws.
     * 
     * @param manager Entry point to ContainerRegistryManager.
     */
    public static void cacheRuleCreateUserAssignedMIAuthAws(
        com.azure.resourcemanager.containerregistry.ContainerRegistryManager manager) {
        manager.serviceClient().getCacheRules().create("myResourceGroup", "myRegistry", "myCacheRule",
            new CacheRuleInner().withIdentity(new IdentityProperties().withType(ResourceIdentityType.USER_ASSIGNED)
                .withUserAssignedIdentities(mapOf(
                    "/subscriptions/12345678-1234-1234-1234-123456789012/resourceGroups/myResourceGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/myUserAssignedIdentity",
                    new UserIdentityProperties())))
                .withAdditionalAuthenticationProperties(
                    new EcrAuthenticationProperties().withAccountNumber("123456789012").withRegion("us-west-2")
                        .withRoleName("my-ecr-pull-role").withRolePath("/acr/"))
                .withSourceRepository("123456789012.dkr.ecr.us-west-2.amazonaws.com/hello-world")
                .withTargetRepository("cached-ecr/hello-world"),
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
