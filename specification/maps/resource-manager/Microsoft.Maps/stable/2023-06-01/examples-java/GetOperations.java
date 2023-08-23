/** Samples for Maps ListOperations. */
public final class Main {
    /*
     * x-ms-original-file: specification/maps/resource-manager/Microsoft.Maps/stable/2023-06-01/examples/GetOperations.json
     */
    /**
     * Sample code: Get Operations.
     *
     * @param manager Entry point to AzureMapsManager.
     */
    public static void getOperations(com.azure.resourcemanager.maps.AzureMapsManager manager) {
        manager.maps().listOperations(com.azure.core.util.Context.NONE);
    }
}
