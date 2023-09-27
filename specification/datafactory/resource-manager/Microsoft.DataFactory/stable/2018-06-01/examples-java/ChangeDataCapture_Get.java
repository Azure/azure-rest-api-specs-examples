/** Samples for ChangeDataCapture Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/ChangeDataCapture_Get.json
     */
    /**
     * Sample code: ChangeDataCapture_Get.
     *
     * @param manager Entry point to DataFactoryManager.
     */
    public static void changeDataCaptureGet(com.azure.resourcemanager.datafactory.DataFactoryManager manager) {
        manager
            .changeDataCaptures()
            .getWithResponse(
                "exampleResourceGroup",
                "exampleFactoryName",
                "exampleChangeDataCapture",
                null,
                com.azure.core.util.Context.NONE);
    }
}
