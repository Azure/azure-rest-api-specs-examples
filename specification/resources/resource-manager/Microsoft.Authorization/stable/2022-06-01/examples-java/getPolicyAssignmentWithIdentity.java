
/** Samples for PolicyAssignments Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2022-06-01/examples/
     * getPolicyAssignmentWithIdentity.json
     */
    /**
     * Sample code: Retrieve a policy assignment with a system assigned identity.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        retrieveAPolicyAssignmentWithASystemAssignedIdentity(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.genericResources().manager().policyClient().getPolicyAssignments().getWithResponse(
            "subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2", "EnforceNaming", com.azure.core.util.Context.NONE);
    }
}
