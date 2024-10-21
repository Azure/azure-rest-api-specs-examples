
/**
 * Samples for DataflowProfile ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-15-preview/DataflowProfile_ListByResourceGroup_MaximumSet_Gen.json
     */
    /**
     * Sample code: DataflowProfile_ListByResourceGroup.
     * 
     * @param manager Entry point to IoTOperationsManager.
     */
    public static void
        dataflowProfileListByResourceGroup(com.azure.resourcemanager.iotoperations.IoTOperationsManager manager) {
        manager.dataflowProfiles().listByResourceGroup("rgiotoperations", "resource-name123",
            com.azure.core.util.Context.NONE);
    }
}
