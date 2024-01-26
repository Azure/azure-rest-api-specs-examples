
/** Samples for PolicyDefinitions List. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/resources/resource-manager/Microsoft.Authorization/stable/2021-06-01/examples/listPolicyDefinitions
     * .json
     */
    /**
     * Sample code: List policy definitions by subscription.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listPolicyDefinitionsBySubscription(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.genericResources().manager().policyClient().getPolicyDefinitions().list(null, null,
            com.azure.core.util.Context.NONE);
    }
}
