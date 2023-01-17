/** Samples for ServiceFabrics List. */
public final class Main {
    /*
     * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/ServiceFabrics_List.json
     */
    /**
     * Sample code: ServiceFabrics_List.
     *
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void serviceFabricsList(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager
            .serviceFabrics()
            .list(
                "resourceGroupName",
                "{labName}",
                "{userName}",
                null,
                null,
                null,
                null,
                com.azure.core.util.Context.NONE);
    }
}
