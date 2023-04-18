/** Samples for ResourceProvider SapSizingRecommendations. */
public final class Main {
    /*
     * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/stable/2023-04-01/examples/sapvirtualinstances/SAPSizingRecommendations_S4HANA_HA_AvSet.json
     */
    /**
     * Sample code: SAPSizingRecommendations_S4HANA_DistributedHA_AvSet.
     *
     * @param manager Entry point to WorkloadsManager.
     */
    public static void sAPSizingRecommendationsS4HANADistributedHAAvSet(
        com.azure.resourcemanager.workloads.WorkloadsManager manager) {
        manager
            .resourceProviders()
            .sapSizingRecommendationsWithResponse("centralus", null, com.azure.core.util.Context.NONE);
    }
}
