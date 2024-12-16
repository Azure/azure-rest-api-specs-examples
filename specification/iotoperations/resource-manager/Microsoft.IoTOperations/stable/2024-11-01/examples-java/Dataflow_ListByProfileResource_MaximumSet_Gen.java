
/**
 * Samples for Dataflow ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-11-01/Dataflow_ListByProfileResource_MaximumSet_Gen.json
     */
    /**
     * Sample code: Dataflow_ListByProfileResource.
     * 
     * @param manager Entry point to IoTOperationsManager.
     */
    public static void
        dataflowListByProfileResource(com.azure.resourcemanager.iotoperations.IoTOperationsManager manager) {
        manager.dataflows().listByResourceGroup("rgiotoperations", "resource-name123", "resource-name123",
            com.azure.core.util.Context.NONE);
    }
}
