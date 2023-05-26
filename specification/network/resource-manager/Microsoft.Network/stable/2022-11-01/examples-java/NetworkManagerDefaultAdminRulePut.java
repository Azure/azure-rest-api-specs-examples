import com.azure.resourcemanager.network.models.DefaultAdminRule;

/** Samples for AdminRules CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-11-01/examples/NetworkManagerDefaultAdminRulePut.json
     */
    /**
     * Sample code: Create a default admin rule.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createADefaultAdminRule(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getAdminRules()
            .createOrUpdateWithResponse(
                "rg1",
                "testNetworkManager",
                "myTestSecurityConfig",
                "testRuleCollection",
                "SampleDefaultAdminRule",
                new DefaultAdminRule().withFlag("AllowVnetInbound"),
                com.azure.core.util.Context.NONE);
    }
}
