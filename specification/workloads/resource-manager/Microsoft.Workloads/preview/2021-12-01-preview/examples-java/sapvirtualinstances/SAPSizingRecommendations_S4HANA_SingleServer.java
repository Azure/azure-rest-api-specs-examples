import com.azure.core.util.Context;

/** Samples for ResourceProvider SapSizingRecommendations. */
public final class Main {
    /*
     * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/preview/2021-12-01-preview/examples/sapvirtualinstances/SAPSizingRecommendations_S4HANA_SingleServer.json
     */
    /**
     * Sample code: SAPSizingRecommendations_S4HANA_SingleServer.
     *
     * @param manager Entry point to WorkloadsManager.
     */
    public static void sAPSizingRecommendationsS4HANASingleServer(
        com.azure.resourcemanager.workloads.WorkloadsManager manager) {
        manager.resourceProviders().sapSizingRecommendationsWithResponse("centralus", null, Context.NONE);
    }
}
