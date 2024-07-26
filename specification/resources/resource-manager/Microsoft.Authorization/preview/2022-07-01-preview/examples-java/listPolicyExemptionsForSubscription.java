
/**
 * Samples for PolicyExemptions List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/resources/resource-manager/Microsoft.Authorization/preview/2022-07-01-preview/examples/
     * listPolicyExemptionsForSubscription.json
     */
    /**
     * Sample code: List policy exemptions that apply to a subscription.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        listPolicyExemptionsThatApplyToASubscription(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.genericResources().manager().policyClient().getPolicyExemptions().list("atScope()",
            com.azure.core.util.Context.NONE);
    }
}
