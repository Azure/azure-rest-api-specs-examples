
import com.azure.resourcemanager.datafactory.models.GetSsisObjectMetadataRequest;

/**
 * Samples for IntegrationRuntimeObjectMetadata Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/
     * IntegrationRuntimeObjectMetadata_Get.json
     */
    /**
     * Sample code: IntegrationRuntimeObjectMetadata_Get.
     * 
     * @param manager Entry point to DataFactoryManager.
     */
    public static void
        integrationRuntimeObjectMetadataGet(com.azure.resourcemanager.datafactory.DataFactoryManager manager) {
        manager.integrationRuntimeObjectMetadatas().getWithResponse("exampleResourceGroup", "exampleFactoryName",
            "testactivityv2", new GetSsisObjectMetadataRequest().withMetadataPath("ssisFolders"),
            com.azure.core.util.Context.NONE);
    }
}
