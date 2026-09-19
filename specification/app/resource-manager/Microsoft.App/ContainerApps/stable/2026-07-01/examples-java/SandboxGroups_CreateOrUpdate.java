
/**
 * Samples for SandboxGroups CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/SandboxGroups_CreateOrUpdate.json
     */
    /**
     * Sample code: Create or update a SandboxGroup.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void
        createOrUpdateASandboxGroup(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.sandboxGroups().define("testgroup").withRegion("East US").withExistingResourceGroup("examplerg")
            .create();
    }
}
