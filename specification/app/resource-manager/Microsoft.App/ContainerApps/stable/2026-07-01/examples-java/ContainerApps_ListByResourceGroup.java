
/**
 * Samples for ContainerApps ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/ContainerApps_ListByResourceGroup.json
     */
    /**
     * Sample code: List Container Apps by resource group.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void
        listContainerAppsByResourceGroup(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.containerApps().listByResourceGroup("rg", com.azure.core.util.Context.NONE);
    }
}
