
import com.azure.resourcemanager.hybridnetwork.models.ExecuteRequestParameters;
import com.azure.resourcemanager.hybridnetwork.models.HttpMethod;
import com.azure.resourcemanager.hybridnetwork.models.RequestMetadata;

/**
 * Samples for NetworkFunctions ExecuteRequest.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/
     * NetworkFunctionsExecuteRequest.json
     */
    /**
     * Sample code: Send request to network function services.
     * 
     * @param manager Entry point to HybridNetworkManager.
     */
    public static void
        sendRequestToNetworkFunctionServices(com.azure.resourcemanager.hybridnetwork.HybridNetworkManager manager) {
        manager.networkFunctions().executeRequest("rg", "testNetworkfunction", new ExecuteRequestParameters()
            .withServiceEndpoint("serviceEndpoint")
            .withRequestMetadata(new RequestMetadata().withRelativePath("/simProfiles/testSimProfile")
                .withHttpMethod(HttpMethod.POST)
                .withSerializedBody(
                    "{\"subscriptionProfile\":\"ChantestSubscription15\",\"permanentKey\":\"00112233445566778899AABBCCDDEEFF\",\"opcOperatorCode\":\"63bfa50ee6523365ff14c1f45f88737d\",\"staticIpAddresses\":{\"internet\":{\"ipv4Addr\":\"198.51.100.1\",\"ipv6Prefix\":\"2001:db8:abcd:12::0/64\"},\"another_network\":{\"ipv6Prefix\":\"2001:111:cdef:22::0/64\"}}}")
                .withApiVersion("apiVersionQueryString")),
            com.azure.core.util.Context.NONE);
    }
}
