import com.azure.core.util.Context;

/** Samples for ConnectedEnvironmentsDaprComponents Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2022-06-01-preview/examples/ConnectedEnvironmentsDaprComponents_Get.json
     */
    /**
     * Sample code: Get Dapr Component.
     *
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void getDaprComponent(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager
            .connectedEnvironmentsDaprComponents()
            .getWithResponse("examplerg", "myenvironment", "reddog", Context.NONE);
    }
}
