
import com.azure.resourcemanager.trafficmanager.fluent.models.EndpointInner;
import com.azure.resourcemanager.trafficmanager.fluent.models.ProfileInner;
import com.azure.resourcemanager.trafficmanager.models.AllowedEndpointRecordType;
import com.azure.resourcemanager.trafficmanager.models.DnsConfig;
import com.azure.resourcemanager.trafficmanager.models.EndpointStatus;
import com.azure.resourcemanager.trafficmanager.models.MonitorConfig;
import com.azure.resourcemanager.trafficmanager.models.MonitorProtocol;
import com.azure.resourcemanager.trafficmanager.models.ProfileStatus;
import com.azure.resourcemanager.trafficmanager.models.TrafficRoutingMethod;
import java.util.Arrays;

/** Samples for Profiles CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/trafficmanager/resource-manager/Microsoft.Network/stable/2022-04-01/examples/Profile-PUT-
     * WithAliasing.json
     */
    /**
     * Sample code: Profile-PUT-WithAliasing.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void profilePUTWithAliasing(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.trafficManagerProfiles().manager().serviceClient().getProfiles().createOrUpdateWithResponse(
            "azuresdkfornetautoresttrafficmanager2583", "azuresdkfornetautoresttrafficmanager6192",
            new ProfileInner().withLocation("global").withProfileStatus(ProfileStatus.ENABLED)
                .withTrafficRoutingMethod(TrafficRoutingMethod.PERFORMANCE)
                .withDnsConfig(
                    new DnsConfig().withRelativeName("azuresdkfornetautoresttrafficmanager6192").withTtl(35L))
                .withMonitorConfig(
                    new MonitorConfig().withProtocol(MonitorProtocol.HTTP).withPort(80L).withPath("/testpath.aspx")
                        .withIntervalInSeconds(10L).withTimeoutInSeconds(5L).withToleratedNumberOfFailures(2L))
                .withEndpoints(Arrays.asList(new EndpointInner().withName("My external endpoint")
                    .withType("Microsoft.network/TrafficManagerProfiles/ExternalEndpoints")
                    .withTarget("foobar.contoso.com").withEndpointStatus(EndpointStatus.ENABLED)
                    .withEndpointLocation("North Europe")))
                .withAllowedEndpointRecordTypes(Arrays.asList(AllowedEndpointRecordType.DOMAIN_NAME)),
            com.azure.core.util.Context.NONE);
    }
}
