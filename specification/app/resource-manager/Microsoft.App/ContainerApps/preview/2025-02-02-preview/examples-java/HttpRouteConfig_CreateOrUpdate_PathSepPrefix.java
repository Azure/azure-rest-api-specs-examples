
import com.azure.resourcemanager.appcontainers.models.BindingType;
import com.azure.resourcemanager.appcontainers.models.CustomDomain;
import com.azure.resourcemanager.appcontainers.models.HttpRoute;
import com.azure.resourcemanager.appcontainers.models.HttpRouteAction;
import com.azure.resourcemanager.appcontainers.models.HttpRouteConfigProperties;
import com.azure.resourcemanager.appcontainers.models.HttpRouteMatch;
import com.azure.resourcemanager.appcontainers.models.HttpRouteRule;
import com.azure.resourcemanager.appcontainers.models.HttpRouteTarget;
import java.util.Arrays;

/**
 * Samples for HttpRouteConfig CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/app/resource-manager/Microsoft.App/ContainerApps/preview/2025-02-02-preview/examples/
     * HttpRouteConfig_CreateOrUpdate_PathSepPrefix.json
     */
    /**
     * Sample code: Create or Update Http Route Path Separated Prefix Rule.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void createOrUpdateHttpRoutePathSeparatedPrefixRule(
        com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.httpRouteConfigs().define("httproutefriendlyname")
            .withExistingManagedEnvironment("examplerg", "testcontainerenv")
            .withProperties(new HttpRouteConfigProperties().withCustomDomains(Arrays.asList(
                new CustomDomain().withName("example.com").withBindingType(BindingType.DISABLED).withCertificateId(
                    "/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/examplerg/providers/Microsoft.App/managedEnvironments/testcontainerenv/certificates/certificate-1")))
                .withRules(Arrays.asList(new HttpRouteRule()
                    .withTargets(Arrays.asList(new HttpRouteTarget().withContainerApp("capp-1").withLabel("label-1")))
                    .withRoutes(Arrays.asList(new HttpRoute()
                        .withMatch(new HttpRouteMatch().withPathSeparatedPrefix("/v1").withCaseSensitive(true))
                        .withAction(new HttpRouteAction().withPrefixRewrite("/v1/api"))))
                    .withDescription("random-description"))))
            .create();
    }
}
