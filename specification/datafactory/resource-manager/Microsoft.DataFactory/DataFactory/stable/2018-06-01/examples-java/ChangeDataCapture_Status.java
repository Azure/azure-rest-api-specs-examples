
/**
 * Samples for ChangeDataCapture Status.
 */
public final class Main {
    /*
     * x-ms-original-file: 2018-06-01/ChangeDataCapture_Status.json
     */
    /**
     * Sample code: ChangeDataCapture_Start.
     * 
     * @param manager Entry point to DataFactoryManager.
     */
    public static void changeDataCaptureStart(com.azure.resourcemanager.datafactory.DataFactoryManager manager) {
        manager.changeDataCaptures().statusWithResponse("exampleResourceGroup", "exampleFactoryName",
            "exampleChangeDataCapture", com.azure.core.util.Context.NONE);
    }
}
