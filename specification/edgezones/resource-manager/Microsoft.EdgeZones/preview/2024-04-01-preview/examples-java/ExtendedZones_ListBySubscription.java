
/**
 * Samples for ExtendedZones List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/edgezones/resource-manager/Microsoft.EdgeZones/preview/2024-04-01-preview/examples/
     * ExtendedZones_ListBySubscription.json
     */
    /**
     * Sample code: ListExtendedZones.
     * 
     * @param manager Entry point to EdgeZonesManager.
     */
    public static void listExtendedZones(com.azure.resourcemanager.edgezones.EdgeZonesManager manager) {
        manager.extendedZones().list(com.azure.core.util.Context.NONE);
    }
}
