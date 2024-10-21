
/**
 * Samples for Instance GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-15-preview/Instance_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: Instance_Get.
     * 
     * @param manager Entry point to IoTOperationsManager.
     */
    public static void instanceGet(com.azure.resourcemanager.iotoperations.IoTOperationsManager manager) {
        manager.instances().getByResourceGroupWithResponse("rgiotoperations", "aio-instance",
            com.azure.core.util.Context.NONE);
    }
}
