
/**
 * Samples for ContainerAppsLabelHistory ListLabelHistory.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/app/resource-manager/Microsoft.App/ContainerApps/preview/2025-02-02-preview/examples/
     * LabelHistory_List.json
     */
    /**
     * Sample code: List Container App's all label history.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void
        listContainerAppSAllLabelHistory(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.containerAppsLabelHistories().listLabelHistory("rg", "testContainerApp", null,
            com.azure.core.util.Context.NONE);
    }
}
