import com.azure.resourcemanager.apimanagement.models.AuthenticationSettingsContract;
import com.azure.resourcemanager.apimanagement.models.BearerTokenSendingMethods;
import com.azure.resourcemanager.apimanagement.models.OpenIdAuthenticationSettingsContract;
import com.azure.resourcemanager.apimanagement.models.Protocol;
import com.azure.resourcemanager.apimanagement.models.SubscriptionKeyParameterNamesContract;
import java.util.Arrays;

/** Samples for Api CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementCreateApiWithOpenIdConnect.json
     */
    /**
     * Sample code: ApiManagementCreateApiWithOpenIdConnect.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementCreateApiWithOpenIdConnect(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .apis()
            .define("tempgroup")
            .withExistingService("rg1", "apimService1")
            .withDisplayName("Swagger Petstore")
            .withServiceUrl("http://petstore.swagger.io/v2")
            .withPath("petstore")
            .withProtocols(Arrays.asList(Protocol.HTTPS))
            .withDescription(
                "This is a sample server Petstore server.  You can find out more about Swagger at"
                    + " [http://swagger.io](http://swagger.io) or on [irc.freenode.net,"
                    + " #swagger](http://swagger.io/irc/).  For this sample, you can use the api key `special-key` to"
                    + " test the authorization filters.")
            .withAuthenticationSettings(
                new AuthenticationSettingsContract()
                    .withOpenid(
                        new OpenIdAuthenticationSettingsContract()
                            .withOpenidProviderId("testopenid")
                            .withBearerTokenSendingMethods(
                                Arrays.asList(BearerTokenSendingMethods.AUTHORIZATION_HEADER))))
            .withSubscriptionKeyParameterNames(
                new SubscriptionKeyParameterNamesContract()
                    .withHeaderProperty("Ocp-Apim-Subscription-Key")
                    .withQuery("subscription-key"))
            .create();
    }
}
