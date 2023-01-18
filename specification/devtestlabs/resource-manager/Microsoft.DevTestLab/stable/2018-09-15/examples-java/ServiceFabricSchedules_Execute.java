/** Samples for ServiceFabricSchedules Execute. */
public final class Main {
    /*
     * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/ServiceFabricSchedules_Execute.json
     */
    /**
     * Sample code: ServiceFabricSchedules_Execute.
     *
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void serviceFabricSchedulesExecute(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager
            .serviceFabricSchedules()
            .execute(
                "resourceGroupName",
                "{labName}",
                "@me",
                "{serviceFrabicName}",
                "{scheduleName}",
                com.azure.core.util.Context.NONE);
    }
}
