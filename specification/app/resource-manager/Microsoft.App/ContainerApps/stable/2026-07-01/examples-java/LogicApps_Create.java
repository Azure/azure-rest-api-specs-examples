
/**
 * Samples for LogicApps CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/LogicApps_Create.json
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
