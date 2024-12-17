
/**
 * Samples for DataflowProfile Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-11-01/DataflowProfile_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: DataflowProfile_Delete.
     * 
     * @param manager Entry point to IoTOperationsManager.
     */
    public static void dataflowProfileDelete(com.azure.resourcemanager.iotoperations.IoTOperationsManager manager) {
        manager.dataflowProfiles().delete("rgiotoperations", "resource-name123", "resource-name123",
            com.azure.core.util.Context.NONE);
    }
}
