import com.azure.core.util.Context;
import com.azure.resourcemanager.cdn.models.ActivatedResourceReference;
import com.azure.resourcemanager.cdn.models.AfdEndpointProtocols;
import com.azure.resourcemanager.cdn.models.AfdQueryStringCachingBehavior;
import com.azure.resourcemanager.cdn.models.AfdRouteCacheConfiguration;
import com.azure.resourcemanager.cdn.models.CompressionSettings;
import com.azure.resourcemanager.cdn.models.EnabledState;
import com.azure.resourcemanager.cdn.models.ForwardingProtocol;
import com.azure.resourcemanager.cdn.models.HttpsRedirect;
import com.azure.resourcemanager.cdn.models.LinkToDefaultDomain;
import com.azure.resourcemanager.cdn.models.ResourceReference;
import com.azure.resourcemanager.cdn.models.RouteUpdateParameters;
import java.util.Arrays;

/** Samples for Routes Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/Routes_Update.json
     */
    /**
     * Sample code: Routes_Update.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void routesUpdate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cdnProfiles()
            .manager()
            .serviceClient()
            .getRoutes()
            .update(
                "RG",
                "profile1",
                "endpoint1",
                "route1",
                new RouteUpdateParameters()
                    .withCustomDomains(
                        Arrays
                            .asList(
                                new ActivatedResourceReference()
                                    .withId(
                                        "/subscriptions/subid/resourceGroups/RG/providers/Microsoft.Cdn/profiles/profile1/customDomains/domain1")))
                    .withOriginGroup(
                        new ResourceReference()
                            .withId(
                                "/subscriptions/subid/resourceGroups/RG/providers/Microsoft.Cdn/profiles/profile1/originGroups/originGroup1"))
                    .withRuleSets(
                        Arrays
                            .asList(
                                new ResourceReference()
                                    .withId(
                                        "/subscriptions/subid/resourceGroups/RG/providers/Microsoft.Cdn/profiles/profile1/ruleSets/ruleSet1")))
                    .withSupportedProtocols(Arrays.asList(AfdEndpointProtocols.HTTPS, AfdEndpointProtocols.HTTP))
                    .withPatternsToMatch(Arrays.asList("/*"))
                    .withCacheConfiguration(
                        new AfdRouteCacheConfiguration()
                            .withQueryStringCachingBehavior(AfdQueryStringCachingBehavior.IGNORE_QUERY_STRING)
                            .withCompressionSettings(
                                new CompressionSettings()
                                    .withContentTypesToCompress(Arrays.asList("text/html", "application/octet-stream"))
                                    .withIsCompressionEnabled(true)))
                    .withForwardingProtocol(ForwardingProtocol.MATCH_REQUEST)
                    .withLinkToDefaultDomain(LinkToDefaultDomain.ENABLED)
                    .withHttpsRedirect(HttpsRedirect.ENABLED)
                    .withEnabledState(EnabledState.ENABLED),
                Context.NONE);
    }
}
