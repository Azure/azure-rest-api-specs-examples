
import com.azure.resourcemanager.datafactory.models.ChangeDataCaptureResource;

/**
 * Samples for ChangeDataCapture CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/
     * ChangeDataCapture_Update.json
     */
    /**
     * Sample code: ChangeDataCapture_Update.
     * 
     * @param manager Entry point to DataFactoryManager.
     */
    public static void changeDataCaptureUpdate(com.azure.resourcemanager.datafactory.DataFactoryManager manager) {
        ChangeDataCaptureResource resource = manager.changeDataCaptures().getWithResponse("exampleResourceGroup",
            "exampleFactoryName", "exampleChangeDataCapture", null, com.azure.core.util.Context.NONE).getValue();
        resource.update().withDescription(
            "Sample demo change data capture to transfer data from delimited (csv) to Azure SQL Database. Updating table mappings.")
            .withAllowVNetOverride(false).withStatus("Stopped").apply();
    }
}
