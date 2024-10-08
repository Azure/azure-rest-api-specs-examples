
import com.azure.resourcemanager.trafficmanager.fluent.models.ProfileInner;
import com.azure.resourcemanager.trafficmanager.models.MonitorConfig;
import com.azure.resourcemanager.trafficmanager.models.MonitorConfigCustomHeadersItem;
import com.azure.resourcemanager.trafficmanager.models.MonitorProtocol;
import java.util.Arrays;

/**
 * Samples for Profiles Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/trafficmanager/resource-manager/Microsoft.Network/stable/2022-04-01/examples/Profile-PATCH-
     * MonitorConfig.json
     */
    /**
     * Sample code: Profile-PATCH-MonitorConfig.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void profilePATCHMonitorConfig(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.trafficManagerProfiles().manager().serviceClient().getProfiles()
            .updateWithResponse("azuresdkfornetautoresttrafficmanager2583", "azuresdkfornetautoresttrafficmanager6192",
                new ProfileInner().withMonitorConfig(
                    new MonitorConfig().withProtocol(MonitorProtocol.HTTP).withPort(80L).withPath("/testpath.aspx")
                        .withIntervalInSeconds(30L).withTimeoutInSeconds(6L).withToleratedNumberOfFailures(4L)
                        .withCustomHeaders(Arrays.asList(
                            new MonitorConfigCustomHeadersItem().withName("header-1").withValue("value-1"),
                            new MonitorConfigCustomHeadersItem().withName("header-2").withValue("value-2")))),
                com.azure.core.util.Context.NONE);
    }
}
