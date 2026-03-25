
/**
 * Samples for ChangeDataCapture Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2018-06-01/ChangeDataCapture_Get.json
     */
    /**
     * Sample code: ChangeDataCapture_Get.
     * 
     * @param manager Entry point to DataFactoryManager.
     */
    public static void changeDataCaptureGet(com.azure.resourcemanager.datafactory.DataFactoryManager manager) {
        manager.changeDataCaptures().getWithResponse("exampleResourceGroup", "exampleFactoryName",
            "exampleChangeDataCapture", null, com.azure.core.util.Context.NONE);
    }
}
