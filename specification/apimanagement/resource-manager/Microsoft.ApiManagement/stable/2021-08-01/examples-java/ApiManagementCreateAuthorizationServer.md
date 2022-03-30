Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-apimanagement_1.0.0-beta.3/sdk/apimanagement/azure-resourcemanager-apimanagement/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.resourcemanager.apimanagement.models.AuthorizationMethod;
import com.azure.resourcemanager.apimanagement.models.BearerTokenSendingMethod;
import com.azure.resourcemanager.apimanagement.models.GrantType;
import java.util.Arrays;

/** Samples for AuthorizationServer CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementCreateAuthorizationServer.json
     */
    /**
     * Sample code: ApiManagementCreateAuthorizationServer.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementCreateAuthorizationServer(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .authorizationServers()
            .define("newauthServer")
            .withExistingService("rg1", "apimService1")
            .withDisplayName("test2")
            .withClientRegistrationEndpoint("https://www.contoso.com/apps")
            .withAuthorizationEndpoint("https://www.contoso.com/oauth2/auth")
            .withGrantTypes(Arrays.asList(GrantType.AUTHORIZATION_CODE, GrantType.IMPLICIT))
            .withClientId("1")
            .withClientSecret("2")
            .withDescription("test server")
            .withAuthorizationMethods(Arrays.asList(AuthorizationMethod.GET))
            .withTokenEndpoint("https://www.contoso.com/oauth2/token")
            .withSupportState(true)
            .withDefaultScope("read write")
            .withBearerTokenSendingMethods(Arrays.asList(BearerTokenSendingMethod.AUTHORIZATION_HEADER))
            .withResourceOwnerUsername("un")
            .withResourceOwnerPassword("pwd")
            .create();
    }
}
```
