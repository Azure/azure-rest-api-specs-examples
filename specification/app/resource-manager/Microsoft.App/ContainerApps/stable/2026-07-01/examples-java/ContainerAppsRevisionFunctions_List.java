
/**
 * Samples for ContainerAppsRevisionFunctions List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/ContainerAppsRevisionFunctions_List.json
     */
    /**
     * Sample code: List Container App Revision's functions.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void
        listContainerAppRevisionSFunctions(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.containerAppsRevisionFunctions().list("myResourceGroup", "myContainerApp", "myContainerApp-abc123",
            com.azure.core.util.Context.NONE);
    }
}
