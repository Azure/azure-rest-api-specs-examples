
/**
 * Samples for LogicApps CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/app/resource-manager/Microsoft.App/ContainerApps/stable/2025-07-01/examples/LogicApps_Create.json
     */
    /**
     * Sample code: Create logic app extension.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void
        createLogicAppExtension(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.logicApps().define("testcontainerApp0").withExistingContainerApp("examplerg", "testcontainerApp0")
            .create();
    }
}
