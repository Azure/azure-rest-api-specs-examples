/** Samples for Firmware ListByWorkspace. */
public final class Main {
    /*
     * x-ms-original-file: specification/fist/resource-manager/Microsoft.IoTFirmwareDefense/preview/2023-02-08-preview/examples/Firmware_ListByWorkspace_MinimumSet_Gen.json
     */
    /**
     * Sample code: Firmware_ListByWorkspace_MinimumSet_Gen.
     *
     * @param manager Entry point to IoTFirmwareDefenseManager.
     */
    public static void firmwareListByWorkspaceMinimumSetGen(
        com.azure.resourcemanager.iotfirmwaredefense.IoTFirmwareDefenseManager manager) {
        manager.firmwares().listByWorkspace("rgworkspaces-firmwares", "A7", com.azure.core.util.Context.NONE);
    }
}
