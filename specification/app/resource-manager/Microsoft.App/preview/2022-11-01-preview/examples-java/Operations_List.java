/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2022-11-01-preview/examples/Operations_List.json
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
