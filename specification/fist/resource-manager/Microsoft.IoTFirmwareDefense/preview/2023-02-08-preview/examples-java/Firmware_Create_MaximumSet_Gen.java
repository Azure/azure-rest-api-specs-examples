import com.azure.core.management.serializer.SerializerFactory;
import com.azure.core.util.serializer.SerializerEncoding;
import com.azure.resourcemanager.iotfirmwaredefense.models.Status;
import java.io.IOException;
import java.util.Arrays;

/** Samples for Firmware Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/fist/resource-manager/Microsoft.IoTFirmwareDefense/preview/2023-02-08-preview/examples/Firmware_Create_MaximumSet_Gen.json
     */
    /**
     * Sample code: Firmware_Create_MaximumSet_Gen.
     *
     * @param manager Entry point to IoTFirmwareDefenseManager.
     */
    public static void firmwareCreateMaximumSetGen(
        com.azure.resourcemanager.iotfirmwaredefense.IoTFirmwareDefenseManager manager) throws IOException {
        manager
            .firmwares()
            .define("umrkdttp")
            .withExistingWorkspace("rgworkspaces-firmwares", "A7")
            .withFileName("wresexxulcdsdd")
            .withVendor("vycmdhgtmepcptyoubztiuudpkcpd")
            .withModel("f")
            .withVersion("s")
            .withDescription("uz")
            .withFileSize(17L)
            .withStatus(Status.PENDING)
            .withStatusMessages(
                Arrays
                    .asList(
                        SerializerFactory
                            .createDefaultManagementSerializerAdapter()
                            .deserialize(
                                "{\"message\":\"ulvhmhokezathzzauiitu\"}", Object.class, SerializerEncoding.JSON)))
            .create();
    }
}
