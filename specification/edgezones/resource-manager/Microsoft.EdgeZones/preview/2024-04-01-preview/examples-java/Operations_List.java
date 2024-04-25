
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/edgezones/resource-manager/Microsoft.EdgeZones/preview/2024-04-01-preview/examples/Operations_List.
     * json
     */
    /**
     * Sample code: ListOperations.
     * 
     * @param manager Entry point to EdgeZonesManager.
     */
    public static void listOperations(com.azure.resourcemanager.edgezones.EdgeZonesManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
