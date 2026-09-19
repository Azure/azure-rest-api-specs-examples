
/**
 * Samples for LogicApps Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/LogicApps_Delete.json
     */
    /**
     * Sample code: Create logic app extension.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void
        createLogicAppExtension(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.logicApps().deleteWithResponse("examplerg", "testcontainerApp0", "testcontainerApp0",
            com.azure.core.util.Context.NONE);
    }
}
