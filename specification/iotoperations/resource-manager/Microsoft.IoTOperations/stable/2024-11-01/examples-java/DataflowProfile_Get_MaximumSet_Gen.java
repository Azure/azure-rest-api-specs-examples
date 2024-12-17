
/**
 * Samples for DataflowProfile Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-11-01/DataflowProfile_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: DataflowProfile_Get.
     * 
     * @param manager Entry point to IoTOperationsManager.
     */
    public static void dataflowProfileGet(com.azure.resourcemanager.iotoperations.IoTOperationsManager manager) {
        manager.dataflowProfiles().getWithResponse("rgiotoperations", "resource-name123", "resource-name123",
            com.azure.core.util.Context.NONE);
    }
}
