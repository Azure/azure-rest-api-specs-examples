/** Samples for ServiceFabrics Start. */
public final class Main {
    /*
     * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/ServiceFabrics_Start.json
     */
    /**
     * Sample code: ServiceFabrics_Start.
     *
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void serviceFabricsStart(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager
            .serviceFabrics()
            .start(
                "resourceGroupName",
                "{labName}",
                "{userName}",
                "{serviceFabricName}",
                com.azure.core.util.Context.NONE);
    }
}
