import com.azure.resourcemanager.apimanagement.models.AuthorizationProviderOAuth2GrantTypes;
import com.azure.resourcemanager.apimanagement.models.AuthorizationProviderOAuth2Settings;
import java.util.HashMap;
import java.util.Map;

/** Samples for AuthorizationProvider CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementCreateAuthorizationProviderGenericOAuth2.json
     */
    /**
     * Sample code: ApiManagementCreateAuthorizationProviderGenericOAuth2.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementCreateAuthorizationProviderGenericOAuth2(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .authorizationProviders()
            .define("eventbrite")
            .withExistingService("rg1", "apimService1")
            .withDisplayName("eventbrite")
            .withIdentityProvider("oauth2")
            .withOauth2(
                new AuthorizationProviderOAuth2Settings()
                    .withRedirectUrl("https://authorization-manager.consent.azure-apim.net/redirect/apim/apimService1")
                    .withGrantTypes(
                        new AuthorizationProviderOAuth2GrantTypes()
                            .withAuthorizationCode(
                                mapOf(
                                    "authorizationUrl",
                                    "fakeTokenPlaceholder",
                                    "clientId",
                                    "ZYIJTBTABHOUQQDLZY",
                                    "clientSecret",
                                    "fakeTokenPlaceholder",
                                    "refreshUrl",
                                    "https://www.eventbrite.com/oauth/token",
                                    "scopes",
                                    "",
                                    "tokenUrl",
                                    "fakeTokenPlaceholder"))))
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
