
import com.azure.resourcemanager.trafficmanager.fluent.models.ProfileInner;
import com.azure.resourcemanager.trafficmanager.models.DnsConfig;
import com.azure.resourcemanager.trafficmanager.models.MonitorConfig;
import com.azure.resourcemanager.trafficmanager.models.MonitorProtocol;
import com.azure.resourcemanager.trafficmanager.models.ProfileStatus;
import com.azure.resourcemanager.trafficmanager.models.TrafficRoutingMethod;

/**
 * Samples for Profiles CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-04-01-preview/Profile-PUT-NoEndpoints.json
     */
    /**
     * Sample code: Profile-PUT-NoEndpoints.
     * 
     * @param manager Entry point to TrafficManager.
     */
    public static void profilePUTNoEndpoints(com.azure.resourcemanager.trafficmanager.TrafficManager manager) {
        manager.serviceClient().getProfiles().createOrUpdateWithResponse("azuresdkfornetautoresttrafficmanager1421",
            "azsmnet6386",
            new ProfileInner().withLocation("global").withProfileStatus(ProfileStatus.ENABLED)
                .withTrafficRoutingMethod(TrafficRoutingMethod.PERFORMANCE)
                .withDnsConfig(new DnsConfig().withRelativeName("azsmnet6386").withTtl(35L)).withMonitorConfig(
                    new MonitorConfig().withProtocol(MonitorProtocol.HTTP).withPort(80L).withPath("/testpath.aspx")),
            com.azure.core.util.Context.NONE);
    }
}
