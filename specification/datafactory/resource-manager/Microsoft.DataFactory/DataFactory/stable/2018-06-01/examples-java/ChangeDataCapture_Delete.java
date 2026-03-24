
/**
 * Samples for ChangeDataCapture Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2018-06-01/ChangeDataCapture_Delete.json
     */
    /**
     * Sample code: ChangeDataCapture_Delete.
     * 
     * @param manager Entry point to DataFactoryManager.
     */
    public static void changeDataCaptureDelete(com.azure.resourcemanager.datafactory.DataFactoryManager manager) {
        manager.changeDataCaptures().deleteWithResponse("exampleResourceGroup", "exampleFactoryName",
            "exampleChangeDataCapture", com.azure.core.util.Context.NONE);
    }
}
