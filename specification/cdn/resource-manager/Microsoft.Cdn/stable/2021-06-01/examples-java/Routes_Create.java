import com.azure.core.management.serializer.SerializerFactory;
import com.azure.core.util.Context;
import com.azure.core.util.serializer.SerializerEncoding;
import com.azure.resourcemanager.cdn.fluent.models.RouteInner;
import com.azure.resourcemanager.cdn.models.ActivatedResourceReference;
import com.azure.resourcemanager.cdn.models.AfdEndpointProtocols;
import com.azure.resourcemanager.cdn.models.AfdQueryStringCachingBehavior;
import com.azure.resourcemanager.cdn.models.AfdRouteCacheConfiguration;
import com.azure.resourcemanager.cdn.models.EnabledState;
import com.azure.resourcemanager.cdn.models.ForwardingProtocol;
import com.azure.resourcemanager.cdn.models.HttpsRedirect;
import com.azure.resourcemanager.cdn.models.LinkToDefaultDomain;
import com.azure.resourcemanager.cdn.models.ResourceReference;
import java.io.IOException;
import java.util.Arrays;

/** Samples for Routes Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/Routes_Create.json
     */
    /**
     * Sample code: Routes_Create.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void routesCreate(com.azure.resourcemanager.AzureResourceManager azure) throws IOException {
        azure
            .cdnProfiles()
            .manager()
            .serviceClient()
            .getRoutes()
            .create(
                "RG",
                "profile1",
                "endpoint1",
                "route1",
                new RouteInner()
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
                            .withQueryStringCachingBehavior(
                                AfdQueryStringCachingBehavior.IGNORE_SPECIFIED_QUERY_STRINGS)
                            .withQueryParameters("querystring=test")
                            .withCompressionSettings(
                                SerializerFactory
                                    .createDefaultManagementSerializerAdapter()
                                    .deserialize(
                                        "{\"contentTypesToCompress\":[\"text/html\",\"application/octet-stream\"],\"isCompressionEnabled\":true}",
                                        Object.class,
                                        SerializerEncoding.JSON)))
                    .withForwardingProtocol(ForwardingProtocol.MATCH_REQUEST)
                    .withLinkToDefaultDomain(LinkToDefaultDomain.ENABLED)
                    .withHttpsRedirect(HttpsRedirect.ENABLED)
                    .withEnabledState(EnabledState.ENABLED),
                Context.NONE);
    }
}
