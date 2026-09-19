
/**
 * Samples for DaprComponents Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/DaprComponents_Get_Secrets.json
     */
    /**
     * Sample code: Get Dapr Component with secrets.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void
        getDaprComponentWithSecrets(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.daprComponents().getWithResponse("examplerg", "myenvironment", "reddog",
            com.azure.core.util.Context.NONE);
    }
}
