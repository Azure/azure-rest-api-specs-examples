/** Samples for DaprComponents Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/stable/2023-05-01/examples/DaprComponents_Get_SecretStoreComponent.json
     */
    /**
     * Sample code: Get Dapr Component with secret store component.
     *
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void getDaprComponentWithSecretStoreComponent(
        com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager
            .daprComponents()
            .getWithResponse("examplerg", "myenvironment", "reddog", com.azure.core.util.Context.NONE);
    }
}
