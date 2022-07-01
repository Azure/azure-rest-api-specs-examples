import com.azure.core.util.Context;

/** Samples for ResourceProvider SapSizingRecommendations. */
public final class Main {
    /*
     * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/preview/2021-12-01-preview/examples/sapvirtualinstances/SAPSizingRecommendations_S4HANA_HA_AvZone.json
     */
    /**
     * Sample code: SAPSizingRecommendations_S4HANA_DistributedHA_AvZone.
     *
     * @param manager Entry point to WorkloadsManager.
     */
    public static void sAPSizingRecommendationsS4HANADistributedHAAvZone(
        com.azure.resourcemanager.workloads.WorkloadsManager manager) {
        manager.resourceProviders().sapSizingRecommendationsWithResponse("centralus", null, Context.NONE);
    }
}
