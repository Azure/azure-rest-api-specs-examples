
import com.azure.resourcemanager.iotfirmwaredefense.models.SummaryName;

/**
 * Samples for Summaries Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/fist/resource-manager/Microsoft.IoTFirmwareDefense/stable/2024-01-10/examples/
     * Summaries_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: Summaries_Get_MaximumSet_Gen.
     * 
     * @param manager Entry point to IoTFirmwareDefenseManager.
     */
    public static void
        summariesGetMaximumSetGen(com.azure.resourcemanager.iotfirmwaredefense.IoTFirmwareDefenseManager manager) {
        manager.summaries().getWithResponse("FirmwareAnalysisRG", "default", "109a9886-50bf-85a8-9d75-000000000000",
            SummaryName.FIRMWARE, com.azure.core.util.Context.NONE);
    }
}
