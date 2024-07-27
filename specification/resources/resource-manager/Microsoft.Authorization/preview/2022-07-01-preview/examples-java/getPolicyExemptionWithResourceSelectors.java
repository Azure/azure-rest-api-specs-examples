
/**
 * Samples for PolicyExemptions Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/resources/resource-manager/Microsoft.Authorization/preview/2022-07-01-preview/examples/
     * getPolicyExemptionWithResourceSelectors.json
     */
    /**
     * Sample code: Retrieve a policy exemption with resource selectors.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        retrieveAPolicyExemptionWithResourceSelectors(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.genericResources().manager().policyClient().getPolicyExemptions().getWithResponse(
            "subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/resourceGroups/demoCluster", "DemoExpensiveVM",
            com.azure.core.util.Context.NONE);
    }
}
