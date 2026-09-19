
/**
 * Samples for ContainerAppsLabelHistory ListLabelHistory.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/LabelHistory_List.json
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
