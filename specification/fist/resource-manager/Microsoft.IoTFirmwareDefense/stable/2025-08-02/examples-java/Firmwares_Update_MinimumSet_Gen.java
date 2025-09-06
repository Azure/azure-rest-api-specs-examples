
import com.azure.resourcemanager.iotfirmwaredefense.models.Firmware;

/**
 * Samples for Firmwares Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-02/Firmwares_Update_MinimumSet_Gen.json
     */
    /**
     * Sample code: Firmwares_Update_MinimumSet_Gen.
     * 
     * @param manager Entry point to IoTFirmwareDefenseManager.
     */
    public static void
        firmwaresUpdateMinimumSetGen(com.azure.resourcemanager.iotfirmwaredefense.IoTFirmwareDefenseManager manager) {
        Firmware resource = manager.firmwares()
            .getWithResponse("rgworkspaces-firmwares", "A7", "umrkdttp", com.azure.core.util.Context.NONE).getValue();
        resource.update().apply();
    }
}
