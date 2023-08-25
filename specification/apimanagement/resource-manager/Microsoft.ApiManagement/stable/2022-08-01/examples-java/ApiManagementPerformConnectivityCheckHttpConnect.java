import com.azure.resourcemanager.apimanagement.models.ConnectivityCheckProtocol;
import com.azure.resourcemanager.apimanagement.models.ConnectivityCheckRequest;
import com.azure.resourcemanager.apimanagement.models.ConnectivityCheckRequestDestination;
import com.azure.resourcemanager.apimanagement.models.ConnectivityCheckRequestProtocolConfiguration;
import com.azure.resourcemanager.apimanagement.models.ConnectivityCheckRequestProtocolConfigurationHttpConfiguration;
import com.azure.resourcemanager.apimanagement.models.ConnectivityCheckRequestSource;
import com.azure.resourcemanager.apimanagement.models.HttpHeader;
import com.azure.resourcemanager.apimanagement.models.Method;
import java.util.Arrays;

/** Samples for ResourceProvider PerformConnectivityCheckAsync. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementPerformConnectivityCheckHttpConnect.json
     */
    /**
     * Sample code: HTTP Connectivity Check.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void hTTPConnectivityCheck(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .resourceProviders()
            .performConnectivityCheckAsync(
                "rg1",
                "apimService1",
                new ConnectivityCheckRequest()
                    .withSource(new ConnectivityCheckRequestSource().withRegion("northeurope"))
                    .withDestination(
                        new ConnectivityCheckRequestDestination().withAddress("https://microsoft.com").withPort(3306L))
                    .withProtocol(ConnectivityCheckProtocol.HTTPS)
                    .withProtocolConfiguration(
                        new ConnectivityCheckRequestProtocolConfiguration()
                            .withHttpConfiguration(
                                new ConnectivityCheckRequestProtocolConfigurationHttpConfiguration()
                                    .withMethod(Method.GET)
                                    .withValidStatusCodes(Arrays.asList(200L, 204L))
                                    .withHeaders(
                                        Arrays
                                            .asList(new HttpHeader().withName("Authorization").withValue("******"))))),
                com.azure.core.util.Context.NONE);
    }
}
