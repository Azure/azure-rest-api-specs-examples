
import com.azure.resourcemanager.apimanagement.models.AuthorizationType;
import com.azure.resourcemanager.apimanagement.models.OAuth2GrantType;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Authorization CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementCreateAuthorizationAADAuthCode.json
     */
    /**
     * Sample code: ApiManagementCreateAuthorizationAADAuthCode.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementCreateAuthorizationAADAuthCode(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.authorizations().define("authz2")
            .withExistingAuthorizationProvider("rg1", "apimService1", "aadwithauthcode")
            .withAuthorizationType(AuthorizationType.OAUTH2).withOAuth2GrantType(OAuth2GrantType.AUTHORIZATION_CODE)
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
