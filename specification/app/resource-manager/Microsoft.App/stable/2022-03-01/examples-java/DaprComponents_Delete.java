import com.azure.core.util.Context;

/** Samples for DaprComponents Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/stable/2022-03-01/examples/DaprComponents_Delete.json
     */
    /**
     * Sample code: Delete dapr component.
     *
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void deleteDaprComponent(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.daprComponents().deleteWithResponse("examplerg", "myenvironment", "reddog", Context.NONE);
    }
}
