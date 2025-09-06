
import com.azure.resourcemanager.iotfirmwaredefense.models.SummaryType;

/**
 * Samples for Summaries Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-02/Summaries_Get_MinimumSet_Gen.json
     */
    /**
     * Sample code: Summaries_Get_MinimumSet_Gen.
     * 
     * @param manager Entry point to IoTFirmwareDefenseManager.
     */
    public static void
        summariesGetMinimumSetGen(com.azure.resourcemanager.iotfirmwaredefense.IoTFirmwareDefenseManager manager) {
        manager.summaries().getWithResponse("FirmwareAnalysisRG", "default", "00000000-0000-0000-0000-000000000000",
            SummaryType.FIRMWARE, com.azure.core.util.Context.NONE);
    }
}
