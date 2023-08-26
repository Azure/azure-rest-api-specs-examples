import com.azure.resourcemanager.apimanagement.models.AuthorizationProviderOAuth2GrantTypes;
import com.azure.resourcemanager.apimanagement.models.AuthorizationProviderOAuth2Settings;
import java.util.HashMap;
import java.util.Map;

/** Samples for AuthorizationProvider CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementCreateAuthorizationProviderOOBGoogle.json
     */
    /**
     * Sample code: ApiManagementCreateAuthorizationProviderOOBGoogle.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementCreateAuthorizationProviderOOBGoogle(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .authorizationProviders()
            .define("google")
            .withExistingService("rg1", "apimService1")
            .withDisplayName("google")
            .withIdentityProvider("google")
            .withOauth2(
                new AuthorizationProviderOAuth2Settings()
                    .withRedirectUrl("https://authorization-manager.consent.azure-apim.net/redirect/apim/apimService1")
                    .withGrantTypes(
                        new AuthorizationProviderOAuth2GrantTypes()
                            .withAuthorizationCode(
                                mapOf(
                                    "clientId",
                                    "508791967882-5qv6o2i99a75un7329vlegtk78kr766h.apps.googleusercontent.com",
                                    "clientSecret",
                                    "fakeTokenPlaceholder",
                                    "scopes",
                                    "openid https://www.googleapis.com/auth/userinfo.profile"
                                        + " https://www.googleapis.com/auth/userinfo.email"))))
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
