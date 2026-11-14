
/**
 * Samples for SandboxGroups List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/SandboxGroupsListBySubscription.json
     */
    /**
     * Sample code: List SandboxGroups by subscription.
     * 
     * @param manager Entry point to ContainerInstanceManager.
     */
    public static void
        listSandboxGroupsBySubscription(com.azure.resourcemanager.containerinstance.ContainerInstanceManager manager) {
        manager.serviceClient().getSandboxGroups().list(com.azure.core.util.Context.NONE);
    }
}
