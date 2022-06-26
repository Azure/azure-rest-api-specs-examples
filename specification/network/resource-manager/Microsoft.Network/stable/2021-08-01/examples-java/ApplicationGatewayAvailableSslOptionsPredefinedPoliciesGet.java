import com.azure.core.util.Context;

/** Samples for ApplicationGateways ListAvailableSslPredefinedPolicies. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/ApplicationGatewayAvailableSslOptionsPredefinedPoliciesGet.json
     */
    /**
     * Sample code: Get Available Ssl Predefined Policies.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAvailableSslPredefinedPolicies(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getApplicationGateways()
            .listAvailableSslPredefinedPolicies(Context.NONE);
    }
}
