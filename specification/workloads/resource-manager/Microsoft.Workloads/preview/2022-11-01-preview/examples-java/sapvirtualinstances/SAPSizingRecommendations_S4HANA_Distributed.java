/** Samples for ResourceProvider SapSizingRecommendations. */
public final class Main {
    /*
     * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/preview/2022-11-01-preview/examples/sapvirtualinstances/SAPSizingRecommendations_S4HANA_Distributed.json
     */
    /**
     * Sample code: SAPSizingRecommendations_S4HANA_Distributed.
     *
     * @param manager Entry point to WorkloadsManager.
     */
    public static void sAPSizingRecommendationsS4HANADistributed(
        com.azure.resourcemanager.workloads.WorkloadsManager manager) {
        manager
            .resourceProviders()
            .sapSizingRecommendationsWithResponse("centralus", null, com.azure.core.util.Context.NONE);
    }
}
