import com.azure.resourcemanager.trafficmanager.fluent.models.ProfileInner;
import com.azure.resourcemanager.trafficmanager.models.DnsConfig;
import com.azure.resourcemanager.trafficmanager.models.MonitorConfig;
import com.azure.resourcemanager.trafficmanager.models.MonitorProtocol;
import com.azure.resourcemanager.trafficmanager.models.ProfileStatus;
import com.azure.resourcemanager.trafficmanager.models.TrafficRoutingMethod;
import com.azure.resourcemanager.trafficmanager.models.TrafficViewEnrollmentStatus;

/** Samples for Profiles CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/trafficmanager/resource-manager/Microsoft.Network/stable/2022-04-01/examples/Profile-PUT-MultiValue.json
     */
    /**
     * Sample code: Profile-PUT-MultiValue.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void profilePUTMultiValue(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .trafficManagerProfiles()
            .manager()
            .serviceClient()
            .getProfiles()
            .createOrUpdateWithResponse(
                "azuresdkfornetautoresttrafficmanager1421",
                "azsmnet6386",
                new ProfileInner()
                    .withLocation("global")
                    .withProfileStatus(ProfileStatus.ENABLED)
                    .withTrafficRoutingMethod(TrafficRoutingMethod.MULTI_VALUE)
                    .withDnsConfig(new DnsConfig().withRelativeName("azsmnet6386").withTtl(35L))
                    .withMonitorConfig(
                        new MonitorConfig().withProtocol(MonitorProtocol.HTTP).withPort(80L).withPath("/testpath.aspx"))
                    .withTrafficViewEnrollmentStatus(TrafficViewEnrollmentStatus.DISABLED)
                    .withMaxReturn(2L),
                com.azure.core.util.Context.NONE);
    }
}
