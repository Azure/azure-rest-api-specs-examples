
/**
 * Samples for PolicyAssignments List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/resources/resource-manager/Microsoft.Authorization/stable/2022-06-01/examples/listPolicyAssignments
     * .json
     */
    /**
     * Sample code: List policy assignments that apply to a subscription.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        listPolicyAssignmentsThatApplyToASubscription(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.genericResources().manager().policyClient().getPolicyAssignments().list("atScope()", null,
            com.azure.core.util.Context.NONE);
    }
}
