import com.azure.core.util.Context;

/** Samples for ContainerApps ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/stable/2022-03-01/examples/ContainerApps_ListByResourceGroup.json
     */
    /**
     * Sample code: List Container Apps by resource group.
     *
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void listContainerAppsByResourceGroup(
        com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.containerApps().listByResourceGroup("rg", Context.NONE);
    }
}
