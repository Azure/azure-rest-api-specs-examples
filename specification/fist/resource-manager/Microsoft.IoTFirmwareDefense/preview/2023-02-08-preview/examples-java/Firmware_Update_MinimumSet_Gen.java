import com.azure.resourcemanager.iotfirmwaredefense.models.Firmware;

/** Samples for Firmware Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/fist/resource-manager/Microsoft.IoTFirmwareDefense/preview/2023-02-08-preview/examples/Firmware_Update_MinimumSet_Gen.json
     */
    /**
     * Sample code: Firmware_Update_MinimumSet_Gen.
     *
     * @param manager Entry point to IoTFirmwareDefenseManager.
     */
    public static void firmwareUpdateMinimumSetGen(
        com.azure.resourcemanager.iotfirmwaredefense.IoTFirmwareDefenseManager manager) {
        Firmware resource =
            manager
                .firmwares()
                .getWithResponse("rgworkspaces-firmwares", "A7", "umrkdttp", com.azure.core.util.Context.NONE)
                .getValue();
        resource.update().apply();
    }
}
