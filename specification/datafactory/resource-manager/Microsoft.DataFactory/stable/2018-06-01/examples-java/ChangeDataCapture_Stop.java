
/**
 * Samples for ChangeDataCapture Stop.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/
     * ChangeDataCapture_Stop.json
     */
    /**
     * Sample code: ChangeDataCapture_Stop.
     * 
     * @param manager Entry point to DataFactoryManager.
     */
    public static void changeDataCaptureStop(com.azure.resourcemanager.datafactory.DataFactoryManager manager) {
        manager.changeDataCaptures().stopWithResponse("exampleResourceGroup", "exampleFactoryName",
            "exampleChangeDataCapture", com.azure.core.util.Context.NONE);
    }
}
