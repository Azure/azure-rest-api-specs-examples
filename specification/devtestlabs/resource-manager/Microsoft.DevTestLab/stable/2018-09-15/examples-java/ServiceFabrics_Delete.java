/** Samples for ServiceFabrics Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/ServiceFabrics_Delete.json
     */
    /**
     * Sample code: ServiceFabrics_Delete.
     *
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void serviceFabricsDelete(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager
            .serviceFabrics()
            .delete(
                "resourceGroupName",
                "{labName}",
                "{userName}",
                "{serviceFabricName}",
                com.azure.core.util.Context.NONE);
    }
}
