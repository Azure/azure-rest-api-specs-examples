
/**
 * Samples for ExtendedZones Unregister.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/ExtendedZones_Unregister.json
     */
    /**
     * Sample code: UnregisterExtendedZone.
     * 
     * @param manager Entry point to EdgeZonesManager.
     */
    public static void unregisterExtendedZone(com.azure.resourcemanager.edgezones.EdgeZonesManager manager) {
        manager.extendedZones().unregisterWithResponse("losangeles", com.azure.core.util.Context.NONE);
    }
}
