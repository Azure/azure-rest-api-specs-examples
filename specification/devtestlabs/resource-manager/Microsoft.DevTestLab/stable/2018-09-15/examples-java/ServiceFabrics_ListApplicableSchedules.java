/** Samples for ServiceFabrics ListApplicableSchedules. */
public final class Main {
    /*
     * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/ServiceFabrics_ListApplicableSchedules.json
     */
    /**
     * Sample code: ServiceFabrics_ListApplicableSchedules.
     *
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void serviceFabricsListApplicableSchedules(
        com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager
            .serviceFabrics()
            .listApplicableSchedulesWithResponse(
                "resourceGroupName",
                "{labName}",
                "{userName}",
                "{serviceFabricName}",
                com.azure.core.util.Context.NONE);
    }
}
