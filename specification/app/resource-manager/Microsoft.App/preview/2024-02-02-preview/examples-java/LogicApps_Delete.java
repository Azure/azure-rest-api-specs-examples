
/**
 * Samples for LogicApps Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/app/resource-manager/Microsoft.App/preview/2024-02-02-preview/examples/LogicApps_Delete.json
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
