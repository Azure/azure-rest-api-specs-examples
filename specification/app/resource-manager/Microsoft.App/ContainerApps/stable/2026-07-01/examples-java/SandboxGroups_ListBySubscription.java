
/**
 * Samples for SandboxGroups List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/SandboxGroups_ListBySubscription.json
     */
    /**
     * Sample code: List SandboxGroups by subscription.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void
        listSandboxGroupsBySubscription(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.sandboxGroups().list(com.azure.core.util.Context.NONE);
    }
}
