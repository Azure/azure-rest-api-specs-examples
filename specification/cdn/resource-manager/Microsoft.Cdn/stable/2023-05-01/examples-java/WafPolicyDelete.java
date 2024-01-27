
/** Samples for Policies Delete. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cdn/resource-manager/Microsoft.Cdn/stable/2023-05-01/examples/WafPolicyDelete.json
     */
    /**
     * Sample code: Delete protection policy.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteProtectionPolicy(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cdnProfiles().manager().serviceClient().getPolicies().deleteWithResponse("rg1", "Policy1",
            com.azure.core.util.Context.NONE);
    }
}
