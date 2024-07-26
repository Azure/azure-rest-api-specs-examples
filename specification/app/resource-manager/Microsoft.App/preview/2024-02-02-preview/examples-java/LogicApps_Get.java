
/**
 * Samples for LogicApps Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/app/resource-manager/Microsoft.App/preview/2024-02-02-preview/examples/LogicApps_Get.json
     */
    /**
     * Sample code: Get logic app extension by name.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void
        getLogicAppExtensionByName(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.logicApps().getWithResponse("examplerg", "testcontainerApp0", "testcontainerApp0",
            com.azure.core.util.Context.NONE);
    }
}
