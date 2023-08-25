import com.azure.resourcemanager.apimanagement.models.ConnectivityCheckRequest;
import com.azure.resourcemanager.apimanagement.models.ConnectivityCheckRequestDestination;
import com.azure.resourcemanager.apimanagement.models.ConnectivityCheckRequestSource;
import com.azure.resourcemanager.apimanagement.models.PreferredIpVersion;

/** Samples for ResourceProvider PerformConnectivityCheckAsync. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementPerformConnectivityCheck.json
     */
    /**
     * Sample code: TCP Connectivity Check.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void tCPConnectivityCheck(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .resourceProviders()
            .performConnectivityCheckAsync(
                "rg1",
                "apimService1",
                new ConnectivityCheckRequest()
                    .withSource(new ConnectivityCheckRequestSource().withRegion("northeurope"))
                    .withDestination(new ConnectivityCheckRequestDestination().withAddress("8.8.8.8").withPort(53L))
                    .withPreferredIpVersion(PreferredIpVersion.IPV4),
                com.azure.core.util.Context.NONE);
    }
}
