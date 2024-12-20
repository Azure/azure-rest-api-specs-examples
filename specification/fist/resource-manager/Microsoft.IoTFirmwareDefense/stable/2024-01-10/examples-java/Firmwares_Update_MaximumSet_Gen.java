
import com.azure.resourcemanager.iotfirmwaredefense.models.Firmware;
import com.azure.resourcemanager.iotfirmwaredefense.models.Status;
import com.azure.resourcemanager.iotfirmwaredefense.models.StatusMessage;
import java.util.Arrays;

/**
 * Samples for Firmwares Update.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/fist/resource-manager/Microsoft.IoTFirmwareDefense/stable/2024-01-10/examples/
     * Firmwares_Update_MaximumSet_Gen.json
     */
    /**
     * Sample code: Firmwares_Update_MaximumSet_Gen.
     * 
     * @param manager Entry point to IoTFirmwareDefenseManager.
     */
    public static void
        firmwaresUpdateMaximumSetGen(com.azure.resourcemanager.iotfirmwaredefense.IoTFirmwareDefenseManager manager) {
        Firmware resource = manager.firmwares()
            .getWithResponse("rgworkspaces-firmwares", "A7", "umrkdttp", com.azure.core.util.Context.NONE).getValue();
        resource.update().withFileName("wresexxulcdsdd").withVendor("vycmdhgtmepcptyoubztiuudpkcpd").withModel("f")
            .withVersion("s").withDescription("uz").withFileSize(17L).withStatus(Status.PENDING)
            .withStatusMessages(Arrays.asList(new StatusMessage().withMessage("ulvhmhokezathzzauiitu"))).apply();
    }
}
