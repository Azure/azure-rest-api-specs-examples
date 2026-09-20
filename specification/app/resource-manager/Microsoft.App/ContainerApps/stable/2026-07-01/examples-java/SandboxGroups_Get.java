
/**
 * Samples for SandboxGroups GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/SandboxGroups_Get.json
     */
    /**
     * Sample code: Get a SandboxGroup.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void getASandboxGroup(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.sandboxGroups().getByResourceGroupWithResponse("examplerg", "testgroup",
            com.azure.core.util.Context.NONE);
    }
}
