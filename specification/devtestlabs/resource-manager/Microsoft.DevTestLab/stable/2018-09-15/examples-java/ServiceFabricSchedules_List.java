/** Samples for ServiceFabricSchedules List. */
public final class Main {
    /*
     * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/ServiceFabricSchedules_List.json
     */
    /**
     * Sample code: ServiceFabricSchedules_List.
     *
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void serviceFabricSchedulesList(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager
            .serviceFabricSchedules()
            .list(
                "resourceGroupName",
                "{labName}",
                "@me",
                "{serviceFrabicName}",
                null,
                null,
                null,
                null,
                com.azure.core.util.Context.NONE);
    }
}
