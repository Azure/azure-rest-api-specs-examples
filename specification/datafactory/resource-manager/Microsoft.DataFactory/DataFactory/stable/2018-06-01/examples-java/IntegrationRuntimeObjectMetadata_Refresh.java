
/**
 * Samples for IntegrationRuntimeObjectMetadata Refresh.
 */
public final class Main {
    /*
     * x-ms-original-file: 2018-06-01/IntegrationRuntimeObjectMetadata_Refresh.json
     */
    /**
     * Sample code: IntegrationRuntimeObjectMetadata_Refresh.
     * 
     * @param manager Entry point to DataFactoryManager.
     */
    public static void
        integrationRuntimeObjectMetadataRefresh(com.azure.resourcemanager.datafactory.DataFactoryManager manager) {
        manager.integrationRuntimeObjectMetadatas().refresh("exampleResourceGroup", "exampleFactoryName",
            "testactivityv2", com.azure.core.util.Context.NONE);
    }
}
