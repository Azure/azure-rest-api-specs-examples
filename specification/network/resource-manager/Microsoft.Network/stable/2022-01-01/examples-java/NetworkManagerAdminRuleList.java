import com.azure.core.util.Context;

/** Samples for AdminRules List. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/NetworkManagerAdminRuleList.json
     */
    /**
     * Sample code: List security admin rules.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listSecurityAdminRules(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getAdminRules()
            .list("rg1", "testNetworkManager", "myTestSecurityConfig", "testRuleCollection", null, null, Context.NONE);
    }
}
