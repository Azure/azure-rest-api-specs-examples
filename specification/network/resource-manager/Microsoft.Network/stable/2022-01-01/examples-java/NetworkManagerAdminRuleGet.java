import com.azure.core.util.Context;

/** Samples for AdminRules Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/NetworkManagerAdminRuleGet.json
     */
    /**
     * Sample code: Gets security admin rule.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getsSecurityAdminRule(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getAdminRules()
            .getWithResponse(
                "rg1",
                "testNetworkManager",
                "myTestSecurityConfig",
                "testRuleCollection",
                "SampleAdminRule",
                Context.NONE);
    }
}
