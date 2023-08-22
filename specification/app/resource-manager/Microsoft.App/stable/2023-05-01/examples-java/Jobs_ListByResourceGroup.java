/** Samples for Jobs ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/stable/2023-05-01/examples/Jobs_ListByResourceGroup.json
     */
    /**
     * Sample code: List Container Apps Jobs by resource group.
     *
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void listContainerAppsJobsByResourceGroup(
        com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.jobs().listByResourceGroup("rg", com.azure.core.util.Context.NONE);
    }
}
