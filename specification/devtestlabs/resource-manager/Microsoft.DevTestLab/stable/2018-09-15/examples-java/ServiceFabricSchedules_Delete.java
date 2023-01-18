/** Samples for ServiceFabricSchedules Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/ServiceFabricSchedules_Delete.json
     */
    /**
     * Sample code: ServiceFabricSchedules_Delete.
     *
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void serviceFabricSchedulesDelete(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager
            .serviceFabricSchedules()
            .deleteWithResponse(
                "resourceGroupName",
                "{labName}",
                "@me",
                "{serviceFrabicName}",
                "{scheduleName}",
                com.azure.core.util.Context.NONE);
    }
}
