
import com.azure.resourcemanager.appservice.fluent.models.SiteAuthSettingsV2Inner;
import com.azure.resourcemanager.appservice.models.AllowedAudiencesValidation;
import com.azure.resourcemanager.appservice.models.AuthPlatform;
import com.azure.resourcemanager.appservice.models.ClientRegistration;
import com.azure.resourcemanager.appservice.models.CookieExpiration;
import com.azure.resourcemanager.appservice.models.CookieExpirationConvention;
import com.azure.resourcemanager.appservice.models.FileSystemTokenStore;
import com.azure.resourcemanager.appservice.models.ForwardProxy;
import com.azure.resourcemanager.appservice.models.ForwardProxyConvention;
import com.azure.resourcemanager.appservice.models.GlobalValidation;
import com.azure.resourcemanager.appservice.models.Google;
import com.azure.resourcemanager.appservice.models.HttpSettings;
import com.azure.resourcemanager.appservice.models.HttpSettingsRoutes;
import com.azure.resourcemanager.appservice.models.IdentityProviders;
import com.azure.resourcemanager.appservice.models.Login;
import com.azure.resourcemanager.appservice.models.LoginRoutes;
import com.azure.resourcemanager.appservice.models.LoginScopes;
import com.azure.resourcemanager.appservice.models.Nonce;
import com.azure.resourcemanager.appservice.models.TokenStore;
import com.azure.resourcemanager.appservice.models.UnauthenticatedClientActionV2;
import java.util.Arrays;

/**
 * Samples for WebApps UpdateAuthSettingsV2.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.Web/stable/2024-11-01/examples/UpdateAuthSettingsV2.json
     */
    /**
     * Sample code: Update Auth Settings V2.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void updateAuthSettingsV2(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getWebApps().updateAuthSettingsV2WithResponse("testrg123",
            "sitef6141",
            new SiteAuthSettingsV2Inner()
                .withPlatform(new AuthPlatform().withEnabled(true).withRuntimeVersion("~1")
                    .withConfigFilePath("/auth/config.json"))
                .withGlobalValidation(new GlobalValidation().withRequireAuthentication(true)
                    .withUnauthenticatedClientAction(UnauthenticatedClientActionV2.RETURN403)
                    .withExcludedPaths(Arrays.asList("/nosecrets/Path")))
                .withIdentityProviders(new IdentityProviders().withGoogle(new Google().withEnabled(true)
                    .withRegistration(new ClientRegistration()
                        .withClientId("42d795a9-8abb-4d06-8534-39528af40f8e.apps.googleusercontent.com")
                        .withClientSecretSettingName("fakeTokenPlaceholder"))
                    .withLogin(new LoginScopes().withScopes(Arrays.asList("admin"))).withValidation(
                        new AllowedAudiencesValidation().withAllowedAudiences(Arrays.asList("https://example.com")))))
                .withLogin(new Login().withRoutes(new LoginRoutes().withLogoutEndpoint("https://app.com/logout"))
                    .withTokenStore(new TokenStore().withEnabled(true).withTokenRefreshExtensionHours(96.0D)
                        .withFileSystem(new FileSystemTokenStore().withDirectory("/wwwroot/sites/example")))
                    .withPreserveUrlFragmentsForLogins(true)
                    .withAllowedExternalRedirectUrls(Arrays.asList("https://someurl.com"))
                    .withCookieExpiration(
                        new CookieExpiration().withConvention(CookieExpirationConvention.IDENTITY_PROVIDER_DERIVED)
                            .withTimeToExpiration("2022:09-01T00:00Z"))
                    .withNonce(new Nonce().withValidateNonce(true)))
                .withHttpSettings(new HttpSettings().withRequireHttps(true)
                    .withRoutes(new HttpSettingsRoutes().withApiPrefix("/authv2/"))
                    .withForwardProxy(new ForwardProxy().withConvention(ForwardProxyConvention.STANDARD)
                        .withCustomHostHeaderName("authHeader").withCustomProtoHeaderName("customProtoHeader"))),
            com.azure.core.util.Context.NONE);
    }
}
