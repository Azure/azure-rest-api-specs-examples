
/**
 * Samples for Dataflow Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-15-preview/Dataflow_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: Dataflow_Get.
     * 
     * @param manager Entry point to IoTOperationsManager.
     */
    public static void dataflowGet(com.azure.resourcemanager.iotoperations.IoTOperationsManager manager) {
        manager.dataflows().getWithResponse("rgiotoperations", "resource-name123", "resource-name123",
            "resource-name123", com.azure.core.util.Context.NONE);
    }
}
