
import com.azure.resourcemanager.apimanagement.models.ApiCreateOrUpdateParameter;
import com.azure.resourcemanager.apimanagement.models.AuthenticationSettingsContract;
import com.azure.resourcemanager.apimanagement.models.OAuth2AuthenticationSettingsContract;
import com.azure.resourcemanager.apimanagement.models.Protocol;
import com.azure.resourcemanager.apimanagement.models.SubscriptionKeyParameterNamesContract;
import java.util.Arrays;

/**
 * Samples for WorkspaceApi CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementCreateWorkspaceApi.json
     */
    /**
     * Sample code: ApiManagementCreateWorkspaceApi.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementCreateWorkspaceApi(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceApis().createOrUpdate("rg1", "apimService1", "wks1", "tempgroup",
            new ApiCreateOrUpdateParameter().withDisplayName("apiname1463")
                .withServiceUrl("http://newechoapi.cloudapp.net/api").withPath("newapiPath")
                .withProtocols(Arrays.asList(Protocol.HTTPS, Protocol.HTTP)).withDescription("apidescription5200")
                .withAuthenticationSettings(
                    new AuthenticationSettingsContract().withOAuth2(new OAuth2AuthenticationSettingsContract()
                        .withAuthorizationServerId("fakeTokenPlaceholder").withScope("oauth2scope2580")))
                .withSubscriptionKeyParameterNames(new SubscriptionKeyParameterNamesContract()
                    .withHeaderProperty("header4520").withQuery("query3037")),
            null, com.azure.core.util.Context.NONE);
    }
}
