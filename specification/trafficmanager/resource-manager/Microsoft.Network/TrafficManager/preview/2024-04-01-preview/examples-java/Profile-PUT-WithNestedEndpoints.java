
import com.azure.resourcemanager.trafficmanager.fluent.models.EndpointInner;
import com.azure.resourcemanager.trafficmanager.fluent.models.ProfileInner;
import com.azure.resourcemanager.trafficmanager.models.DnsConfig;
import com.azure.resourcemanager.trafficmanager.models.EndpointStatus;
import com.azure.resourcemanager.trafficmanager.models.MonitorConfig;
import com.azure.resourcemanager.trafficmanager.models.MonitorProtocol;
import com.azure.resourcemanager.trafficmanager.models.ProfileStatus;
import com.azure.resourcemanager.trafficmanager.models.TrafficRoutingMethod;
import java.util.Arrays;

/**
 * Samples for Profiles CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-04-01-preview/Profile-PUT-WithNestedEndpoints.json
     */
    /**
     * Sample code: Profile-PUT-WithNestedEndpoints.
     * 
     * @param manager Entry point to TrafficManager.
     */
    public static void profilePUTWithNestedEndpoints(com.azure.resourcemanager.trafficmanager.TrafficManager manager) {
        manager.serviceClient().getProfiles()
            .createOrUpdateWithResponse("myresourcegroup", "parentprofile",
                new ProfileInner().withLocation("global").withProfileStatus(ProfileStatus.ENABLED)
                    .withTrafficRoutingMethod(TrafficRoutingMethod.PRIORITY)
                    .withDnsConfig(
                        new DnsConfig().withRelativeName("parentprofile").withTtl(35L))
                    .withMonitorConfig(new MonitorConfig().withProtocol(MonitorProtocol.HTTP).withPort(80L)
                        .withPath("/testpath.aspx").withIntervalInSeconds(
                            10L)
                        .withTimeoutInSeconds(5L).withToleratedNumberOfFailures(2L))
                    .withEndpoints(Arrays.asList(
                        new EndpointInner().withName("MyFirstNestedEndpoint")
                            .withType("Microsoft.Network/trafficManagerProfiles/nestedEndpoints")
                            .withTarget("firstnestedprofile.tmpreview.watmtest.azure-test.net")
                            .withEndpointStatus(EndpointStatus.ENABLED).withWeight(1L).withPriority(1L)
                            .withMinChildEndpoints(2L).withMinChildEndpointsIPv4(1L).withMinChildEndpointsIPv6(2L),
                        new EndpointInner().withName("MySecondNestedEndpoint")
                            .withType("Microsoft.Network/trafficManagerProfiles/nestedEndpoints")
                            .withTarget("secondnestedprofile.tmpreview.watmtest.azure-test.net")
                            .withEndpointStatus(EndpointStatus.ENABLED).withWeight(1L).withPriority(2L)
                            .withMinChildEndpoints(2L).withMinChildEndpointsIPv4(2L).withMinChildEndpointsIPv6(1L))),
                com.azure.core.util.Context.NONE);
    }
}
