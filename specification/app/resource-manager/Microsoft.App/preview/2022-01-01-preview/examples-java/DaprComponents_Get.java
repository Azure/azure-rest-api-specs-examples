import com.azure.core.util.Context;

/** Samples for DaprComponents Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2022-01-01-preview/examples/DaprComponents_Get.json
     */
    /**
     * Sample code: Get Dapr Component.
     *
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void getDaprComponent(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.daprComponents().getWithResponse("examplerg", "myenvironment", "reddog", Context.NONE);
    }
}
