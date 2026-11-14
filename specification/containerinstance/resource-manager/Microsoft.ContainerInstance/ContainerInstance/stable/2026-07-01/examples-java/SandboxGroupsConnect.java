
/**
 * Samples for SandboxGroups Connect.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/SandboxGroupsConnect.json
     */
    /**
     * Sample code: Connect to a SandboxGroup.
     * 
     * @param manager Entry point to ContainerInstanceManager.
     */
    public static void
        connectToASandboxGroup(com.azure.resourcemanager.containerinstance.ContainerInstanceManager manager) {
        manager.serviceClient().getSandboxGroups().connectWithResponse("myResourceGroup", "mySandboxGroup",
            com.azure.core.util.Context.NONE);
    }
}
