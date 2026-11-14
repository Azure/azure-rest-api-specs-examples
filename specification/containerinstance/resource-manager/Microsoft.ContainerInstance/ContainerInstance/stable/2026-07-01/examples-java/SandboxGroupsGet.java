
/**
 * Samples for SandboxGroups GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/SandboxGroupsGet.json
     */
    /**
     * Sample code: Get a SandboxGroup.
     * 
     * @param manager Entry point to ContainerInstanceManager.
     */
    public static void getASandboxGroup(com.azure.resourcemanager.containerinstance.ContainerInstanceManager manager) {
        manager.serviceClient().getSandboxGroups().getByResourceGroupWithResponse("myResourceGroup", "mySandboxGroup",
            com.azure.core.util.Context.NONE);
    }
}
