import com.azure.core.util.Context;

/** Samples for DaprComponents List. */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2022-01-01-preview/examples/DaprComponents_List.json
     */
    /**
     * Sample code: List Dapr Components.
     *
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void listDaprComponents(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.daprComponents().list("examplerg", "myenvironment", Context.NONE);
    }
}
