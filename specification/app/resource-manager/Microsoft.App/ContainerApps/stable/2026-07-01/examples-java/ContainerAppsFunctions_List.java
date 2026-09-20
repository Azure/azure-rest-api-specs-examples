
/**
 * Samples for ContainerAppsFunctions List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/ContainerAppsFunctions_List.json
     */
    /**
     * Sample code: List Container App's functions.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void
        listContainerAppSFunctions(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.containerAppsFunctions().list("myResourceGroup", "myContainerApp", com.azure.core.util.Context.NONE);
    }
}
