/** Samples for SpatialAnchorsAccounts ListKeys. */
public final class Main {
    /*
     * x-ms-original-file: specification/mixedreality/resource-manager/Microsoft.MixedReality/stable/2021-01-01/examples/spatial-anchors/ListKeys.json
     */
    /**
     * Sample code: List spatial anchor account key.
     *
     * @param manager Entry point to MixedRealityManager.
     */
    public static void listSpatialAnchorAccountKey(com.azure.resourcemanager.mixedreality.MixedRealityManager manager) {
        manager
            .spatialAnchorsAccounts()
            .listKeysWithResponse("MyResourceGroup", "MyAccount", com.azure.core.util.Context.NONE);
    }
}
