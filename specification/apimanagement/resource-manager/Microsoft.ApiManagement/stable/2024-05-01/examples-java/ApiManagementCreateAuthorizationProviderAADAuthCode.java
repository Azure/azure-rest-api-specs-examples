
import com.azure.resourcemanager.apimanagement.models.AuthorizationProviderOAuth2GrantTypes;
import com.azure.resourcemanager.apimanagement.models.AuthorizationProviderOAuth2Settings;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for AuthorizationProvider CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementCreateAuthorizationProviderAADAuthCode.json
     */
    /**
     * Sample code: ApiManagementCreateAuthorizationProviderAADAuthCode.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementCreateAuthorizationProviderAADAuthCode(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.authorizationProviders().define("aadwithauthcode").withExistingService("rg1", "apimService1")
            .withDisplayName("aadwithauthcode").withIdentityProvider("aad")
            .withOauth2(new AuthorizationProviderOAuth2Settings()
                .withRedirectUrl("https://authorization-manager.consent.azure-apim.net/redirect/apim/apimService1")
                .withGrantTypes(new AuthorizationProviderOAuth2GrantTypes()
                    .withAuthorizationCode(mapOf("clientId", "clientsecretid", "clientSecret", "fakeTokenPlaceholder",
                        "resourceUri", "https://graph.microsoft.com", "scopes", "User.Read.All Group.Read.All"))))
            .create();
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
