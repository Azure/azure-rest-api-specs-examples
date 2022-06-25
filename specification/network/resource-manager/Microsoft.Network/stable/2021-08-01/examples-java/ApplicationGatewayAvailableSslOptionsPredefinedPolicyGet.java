import com.azure.core.util.Context;

/** Samples for ApplicationGateways GetSslPredefinedPolicy. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/ApplicationGatewayAvailableSslOptionsPredefinedPolicyGet.json
     */
    /**
     * Sample code: Get Available Ssl Predefined Policy by name.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAvailableSslPredefinedPolicyByName(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getApplicationGateways()
            .getSslPredefinedPolicyWithResponse("AppGwSslPolicy20150501", Context.NONE);
    }
}
