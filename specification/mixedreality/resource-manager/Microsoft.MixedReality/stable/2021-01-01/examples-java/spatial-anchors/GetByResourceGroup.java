/** Samples for SpatialAnchorsAccounts ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/mixedreality/resource-manager/Microsoft.MixedReality/stable/2021-01-01/examples/spatial-anchors/GetByResourceGroup.json
     */
    /**
     * Sample code: List spatial anchor accounts by resource group.
     *
     * @param manager Entry point to MixedRealityManager.
     */
    public static void listSpatialAnchorAccountsByResourceGroup(
        com.azure.resourcemanager.mixedreality.MixedRealityManager manager) {
        manager.spatialAnchorsAccounts().listByResourceGroup("MyResourceGroup", com.azure.core.util.Context.NONE);
    }
}
