
/**
 * Samples for ExtendedZones Register.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-04-01-preview/ExtendedZones_Register.json
     */
    /**
     * Sample code: RegisterExtendedZone.
     * 
     * @param manager Entry point to EdgeZonesManager.
     */
    public static void registerExtendedZone(com.azure.resourcemanager.edgezones.EdgeZonesManager manager) {
        manager.extendedZones().registerWithResponse("losangeles", com.azure.core.util.Context.NONE);
    }
}
