
/**
 * Samples for ContainerAppsLabelHistory DeleteLabelHistory.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/LabelHistory_Delete.json
     */
    /**
     * Sample code: Delete Container App's label history for a given label.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void deleteContainerAppSLabelHistoryForAGivenLabel(
        com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.containerAppsLabelHistories().deleteLabelHistoryWithResponse("rg", "testContainerApp", "dev",
            com.azure.core.util.Context.NONE);
    }
}
