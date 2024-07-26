
/**
 * Samples for Builds CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2024-02-02-preview/examples/
     * Builds_CreateOrUpdate_NoConfig.json
     */
    /**
     * Sample code: Builds_CreateOrUpdate_NoConfig.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void
        buildsCreateOrUpdateNoConfig(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.builds().define("testBuild").withExistingBuilder("rg", "testBuilder").create();
    }
}
