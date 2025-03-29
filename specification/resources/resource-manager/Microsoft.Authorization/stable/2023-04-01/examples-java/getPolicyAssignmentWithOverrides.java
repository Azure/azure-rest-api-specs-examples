
/**
 * Samples for PolicyAssignments Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2023-04-01/examples/
     * getPolicyAssignmentWithOverrides.json
     */
    /**
     * Sample code: Retrieve a policy assignment with overrides.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void retrieveAPolicyAssignmentWithOverrides(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.genericResources().manager().policyClient().getPolicyAssignments().getWithResponse(
            "subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2", "CostManagement", null,
            com.azure.core.util.Context.NONE);
    }
}
