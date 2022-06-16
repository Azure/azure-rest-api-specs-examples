import com.azure.core.util.Context;

/** Samples for Remediations ListForResource. */
public final class Main {
    /*
     * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2021-10-01/examples/Remediations_ListResourceScope.json
     */
    /**
     * Sample code: List remediations at individual resource scope.
     *
     * @param manager Entry point to PolicyInsightsManager.
     */
    public static void listRemediationsAtIndividualResourceScope(
        com.azure.resourcemanager.policyinsights.PolicyInsightsManager manager) {
        manager
            .remediations()
            .listForResource(
                "subscriptions/35ee058e-5fa0-414c-8145-3ebb8d09b6e2/resourcegroups/myResourceGroup/providers/microsoft.storage/storageaccounts/storAc1",
                null,
                null,
                Context.NONE);
    }
}
