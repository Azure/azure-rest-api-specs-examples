
/**
 * Samples for SpatialAnchorsAccounts GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/mixedreality/resource-manager/Microsoft.MixedReality/stable/2021-01-01/examples/spatial-anchors/Get
     * .json
     */
    /**
     * Sample code: Get spatial anchors account.
     * 
     * @param manager Entry point to MixedRealityManager.
     */
    public static void getSpatialAnchorsAccount(com.azure.resourcemanager.mixedreality.MixedRealityManager manager) {
        manager.spatialAnchorsAccounts().getByResourceGroupWithResponse("MyResourceGroup", "MyAccount",
            com.azure.core.util.Context.NONE);
    }
}
