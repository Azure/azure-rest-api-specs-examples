
/**
 * Samples for SandboxGroups Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/SandboxGroups_Delete.json
     */
    /**
     * Sample code: Delete a SandboxGroup.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void deleteASandboxGroup(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.sandboxGroups().delete("examplerg", "testgroup", com.azure.core.util.Context.NONE);
    }
}
