
/**
 * Samples for SandboxGroups Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/SandboxGroupsDelete.json
     */
    /**
     * Sample code: Delete a SandboxGroup.
     * 
     * @param manager Entry point to ContainerInstanceManager.
     */
    public static void
        deleteASandboxGroup(com.azure.resourcemanager.containerinstance.ContainerInstanceManager manager) {
        manager.serviceClient().getSandboxGroups().delete("myResourceGroup", "mySandboxGroup",
            com.azure.core.util.Context.NONE);
    }
}
