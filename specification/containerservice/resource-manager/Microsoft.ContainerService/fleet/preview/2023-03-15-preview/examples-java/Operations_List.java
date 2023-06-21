/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/fleet/preview/2023-03-15-preview/examples/Operations_List.json
     */
    /**
     * Sample code: List the operations for the provider.
     *
     * @param manager Entry point to ContainerServiceFleetManager.
     */
    public static void listTheOperationsForTheProvider(
        com.azure.resourcemanager.containerservicefleet.ContainerServiceFleetManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
