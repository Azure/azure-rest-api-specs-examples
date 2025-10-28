
/**
 * Samples for ContainerAppsSessionPools ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/ContainerApps/stable/2025-07-01/examples/
     * SessionPools_ListByResourceGroup.json
     */
    /**
     * Sample code: List Session Pools by resource group.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void
        listSessionPoolsByResourceGroup(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.containerAppsSessionPools().listByResourceGroup("rg", com.azure.core.util.Context.NONE);
    }
}
