/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/stable/2023-05-01/examples/Operations_List.json
     */
    /**
     * Sample code: List all operations.
     *
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void listAllOperations(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
