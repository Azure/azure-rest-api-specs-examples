
/**
 * Samples for ChangeDataCapture Start.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/
     * ChangeDataCapture_Start.json
     */
    /**
     * Sample code: ChangeDataCapture_Start.
     * 
     * @param manager Entry point to DataFactoryManager.
     */
    public static void changeDataCaptureStart(com.azure.resourcemanager.datafactory.DataFactoryManager manager) {
        manager.changeDataCaptures().startWithResponse("exampleResourceGroup", "exampleFactoryName",
            "exampleChangeDataCapture", com.azure.core.util.Context.NONE);
    }
}
