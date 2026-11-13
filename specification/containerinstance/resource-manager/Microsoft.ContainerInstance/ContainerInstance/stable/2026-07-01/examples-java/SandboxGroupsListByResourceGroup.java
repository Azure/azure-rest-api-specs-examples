
/**
 * Samples for SandboxGroups ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/SandboxGroupsListByResourceGroup.json
     */
    /**
     * Sample code: List SandboxGroups by resource group.
     * 
     * @param manager Entry point to ContainerInstanceManager.
     */
    public static void
        listSandboxGroupsByResourceGroup(com.azure.resourcemanager.containerinstance.ContainerInstanceManager manager) {
        manager.serviceClient().getSandboxGroups().listByResourceGroup("myResourceGroup",
            com.azure.core.util.Context.NONE);
    }
}
