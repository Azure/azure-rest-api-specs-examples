
/**
 * Samples for ChangeDataCapture ListByFactory.
 */
public final class Main {
    /*
     * x-ms-original-file: 2018-06-01/ChangeDataCapture_ListByFactory.json
     */
    /**
     * Sample code: ChangeDataCapture_ListByFactory.
     * 
     * @param manager Entry point to DataFactoryManager.
     */
    public static void
        changeDataCaptureListByFactory(com.azure.resourcemanager.datafactory.DataFactoryManager manager) {
        manager.changeDataCaptures().listByFactory("exampleResourceGroup", "exampleFactoryName",
            com.azure.core.util.Context.NONE);
    }
}
