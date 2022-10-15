import com.azure.core.util.Context;

/** Samples for Attestations ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2022-09-01/examples/Attestations_ListResourceGroupScope.json
     */
    /**
     * Sample code: List attestations at resource group scope.
     *
     * @param manager Entry point to PolicyInsightsManager.
     */
    public static void listAttestationsAtResourceGroupScope(
        com.azure.resourcemanager.policyinsights.PolicyInsightsManager manager) {
        manager.attestations().listByResourceGroup("myRg", null, null, Context.NONE);
    }
}
