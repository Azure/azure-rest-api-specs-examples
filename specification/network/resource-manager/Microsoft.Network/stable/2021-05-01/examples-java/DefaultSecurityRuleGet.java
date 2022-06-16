import com.azure.core.util.Context;

/** Samples for DefaultSecurityRules Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/DefaultSecurityRuleGet.json
     */
    /**
     * Sample code: DefaultSecurityRuleGet.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void defaultSecurityRuleGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getDefaultSecurityRules()
            .getWithResponse("testrg", "nsg1", "AllowVnetInBound", Context.NONE);
    }
}
