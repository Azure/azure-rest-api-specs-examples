
import com.azure.resourcemanager.iotfirmwaredefense.models.Status;
import com.azure.resourcemanager.iotfirmwaredefense.models.StatusMessage;
import java.util.Arrays;

/**
 * Samples for Firmwares Create.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/fist/resource-manager/Microsoft.IoTFirmwareDefense/stable/2024-01-10/examples/
     * Firmwares_Create_MaximumSet_Gen.json
     */
    /**
     * Sample code: Firmwares_Create_MaximumSet_Gen.
     * 
     * @param manager Entry point to IoTFirmwareDefenseManager.
     */
    public static void
        firmwaresCreateMaximumSetGen(com.azure.resourcemanager.iotfirmwaredefense.IoTFirmwareDefenseManager manager) {
        manager.firmwares().define("umrkdttp").withExistingWorkspace("rgworkspaces-firmwares", "A7")
            .withFileName("wresexxulcdsdd").withVendor("vycmdhgtmepcptyoubztiuudpkcpd").withModel("f").withVersion("s")
            .withDescription("uz").withFileSize(17L).withStatus(Status.PENDING)
            .withStatusMessages(Arrays.asList(new StatusMessage().withMessage("ulvhmhokezathzzauiitu"))).create();
    }
}
