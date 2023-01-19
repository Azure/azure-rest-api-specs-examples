/** Samples for Clusters ListUpgradableVersions. */
public final class Main {
    /*
     * x-ms-original-file: specification/servicefabric/resource-manager/Microsoft.ServiceFabric/stable/2021-06-01/examples/ListUpgradableVersionsMinMax_example.json
     */
    /**
     * Sample code: Get minimum and maximum code versions.
     *
     * @param manager Entry point to ServiceFabricManager.
     */
    public static void getMinimumAndMaximumCodeVersions(
        com.azure.resourcemanager.servicefabric.ServiceFabricManager manager) {
        manager
            .clusters()
            .listUpgradableVersionsWithResponse("resRg", "myCluster", null, com.azure.core.util.Context.NONE);
    }
}
