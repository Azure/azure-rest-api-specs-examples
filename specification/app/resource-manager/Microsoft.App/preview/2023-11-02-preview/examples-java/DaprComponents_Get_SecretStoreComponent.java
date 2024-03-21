
/**
 * Samples for DaprComponents Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2023-11-02-preview/examples/
     * DaprComponents_Get_SecretStoreComponent.json
     */
    /**
     * Sample code: Get Dapr Component with secret store component.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void getDaprComponentWithSecretStoreComponent(
        com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.daprComponents().getWithResponse("examplerg", "myenvironment", "reddog",
            com.azure.core.util.Context.NONE);
    }

    /*
     * x-ms-original-file:
     * specification/app/resource-manager/Microsoft.App/preview/2023-11-02-preview/examples/DaprComponents_Get_Secrets.
     * json
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
