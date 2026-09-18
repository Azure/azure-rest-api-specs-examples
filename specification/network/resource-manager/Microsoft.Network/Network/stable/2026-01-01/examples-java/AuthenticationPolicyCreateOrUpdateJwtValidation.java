
import com.azure.resourcemanager.network.fluent.models.AuthenticationPolicyInner;
import com.azure.resourcemanager.network.models.AuthenticationPolicyPropertiesFormat;
import com.azure.resourcemanager.network.models.AuthenticationProviderProperties;
import com.azure.resourcemanager.network.models.OnUnauthenticatedRequest;
import com.azure.resourcemanager.network.models.UserTrustProviderType;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for AuthenticationPolicies CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/AuthenticationPolicyCreateOrUpdateJwtValidation.json
     */
    /**
     * Sample code: Creates or updates a JWT validation authentication policy within a resource group.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void createsOrUpdatesAJWTValidationAuthenticationPolicyWithinAResourceGroup(
        com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getAuthenticationPolicies().createOrUpdate("rg1", "jwtValidationPolicy",
            new AuthenticationPolicyInner().withLocation("eastus")
                .withProperties(new AuthenticationPolicyPropertiesFormat()
                    .withUserTrustProviderType(UserTrustProviderType.ENTRA)
                    .withOnUnauthenticatedRequest(OnUnauthenticatedRequest.DENY)
                    .withAuthenticationProperties(new AuthenticationProviderProperties()
                        .withIssuer("https://login.microsoftonline.com/11111111-1111-1111-1111-111111111111/")
                        .withJwksUri(
                            "https://login.microsoftonline.com/11111111-1111-1111-1111-111111111111/discovery/v2.0/keys")
                        .withAudience("api://myapp").withClientId("00000000-0000-0000-0000-000000000001"))),
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
