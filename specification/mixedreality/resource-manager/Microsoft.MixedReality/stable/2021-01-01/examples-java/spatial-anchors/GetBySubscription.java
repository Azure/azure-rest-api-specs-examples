/** Samples for SpatialAnchorsAccounts List. */
public final class Main {
    /*
     * x-ms-original-file: specification/mixedreality/resource-manager/Microsoft.MixedReality/stable/2021-01-01/examples/spatial-anchors/GetBySubscription.json
     */
    /**
     * Sample code: List spatial anchors accounts by subscription.
     *
     * @param manager Entry point to MixedRealityManager.
     */
    public static void listSpatialAnchorsAccountsBySubscription(
        com.azure.resourcemanager.mixedreality.MixedRealityManager manager) {
        manager.spatialAnchorsAccounts().list(com.azure.core.util.Context.NONE);
    }
}
