/** Samples for SpatialAnchorsAccounts Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/mixedreality/resource-manager/Microsoft.MixedReality/stable/2021-01-01/examples/spatial-anchors/Put.json
     */
    /**
     * Sample code: Create spatial anchor account.
     *
     * @param manager Entry point to MixedRealityManager.
     */
    public static void createSpatialAnchorAccount(com.azure.resourcemanager.mixedreality.MixedRealityManager manager) {
        manager
            .spatialAnchorsAccounts()
            .define("MyAccount")
            .withRegion("eastus2euap")
            .withExistingResourceGroup("MyResourceGroup")
            .create();
    }
}
