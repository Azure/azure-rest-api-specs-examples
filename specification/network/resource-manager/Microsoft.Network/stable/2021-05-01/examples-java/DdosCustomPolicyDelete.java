import com.azure.core.util.Context;

/** Samples for DdosCustomPolicies Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/DdosCustomPolicyDelete.json
     */
    /**
     * Sample code: Delete DDoS custom policy.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteDDoSCustomPolicy(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getDdosCustomPolicies()
            .delete("rg1", "test-ddos-custom-policy", Context.NONE);
    }
}
