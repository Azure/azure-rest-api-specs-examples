import com.azure.resourcemanager.apimanagement.models.AuthenticationSettingsContract;
import com.azure.resourcemanager.apimanagement.models.BearerTokenSendingMethods;
import com.azure.resourcemanager.apimanagement.models.OpenIdAuthenticationSettingsContract;
import com.azure.resourcemanager.apimanagement.models.Protocol;
import com.azure.resourcemanager.apimanagement.models.SubscriptionKeyParameterNamesContract;
import java.util.Arrays;

/** Samples for Api CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementCreateApiWithMultipleOpenIdConnectProviders.json
     */
    /**
     * Sample code: ApiManagementCreateApiWithMultipleOpenIdConnectProviders.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementCreateApiWithMultipleOpenIdConnectProviders(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .apis()
            .define("tempgroup")
            .withExistingService("rg1", "apimService1")
            .withDisplayName("apiname1463")
            .withServiceUrl("http://newechoapi.cloudapp.net/api")
            .withPath("newapiPath")
            .withProtocols(Arrays.asList(Protocol.HTTPS, Protocol.HTTP))
            .withDescription("apidescription5200")
            .withAuthenticationSettings(
                new AuthenticationSettingsContract()
                    .withOpenidAuthenticationSettings(
                        Arrays
                            .asList(
                                new OpenIdAuthenticationSettingsContract()
                                    .withOpenidProviderId("openidProviderId2283")
                                    .withBearerTokenSendingMethods(
                                        Arrays.asList(BearerTokenSendingMethods.AUTHORIZATION_HEADER)),
                                new OpenIdAuthenticationSettingsContract()
                                    .withOpenidProviderId("openidProviderId2284")
                                    .withBearerTokenSendingMethods(
                                        Arrays.asList(BearerTokenSendingMethods.AUTHORIZATION_HEADER)))))
            .withSubscriptionKeyParameterNames(
                new SubscriptionKeyParameterNamesContract().withHeaderProperty("header4520").withQuery("query3037"))
            .create();
    }
}
