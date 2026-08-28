
/**
 * Samples for ExtendedZones Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/ExtendedZones_Get.json
     */
    /**
     * Sample code: GetExtendedZone.
     * 
     * @param manager Entry point to EdgeZonesManager.
     */
    public static void getExtendedZone(com.azure.resourcemanager.edgezones.EdgeZonesManager manager) {
        manager.extendedZones().getWithResponse("losangeles", com.azure.core.util.Context.NONE);
    }
}
