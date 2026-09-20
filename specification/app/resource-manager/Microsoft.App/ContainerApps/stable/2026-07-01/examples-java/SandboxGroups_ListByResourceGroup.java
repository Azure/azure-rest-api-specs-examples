
/**
 * Samples for SandboxGroups ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/SandboxGroups_ListByResourceGroup.json
     */
    /**
     * Sample code: List SandboxGroups by resource group.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void
        listSandboxGroupsByResourceGroup(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.sandboxGroups().listByResourceGroup("examplerg", com.azure.core.util.Context.NONE);
    }
}
