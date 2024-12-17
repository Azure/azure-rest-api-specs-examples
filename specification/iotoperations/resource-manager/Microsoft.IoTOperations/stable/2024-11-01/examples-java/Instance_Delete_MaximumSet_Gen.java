
/**
 * Samples for Instance Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-11-01/Instance_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: Instance_Delete.
     * 
     * @param manager Entry point to IoTOperationsManager.
     */
    public static void instanceDelete(com.azure.resourcemanager.iotoperations.IoTOperationsManager manager) {
        manager.instances().delete("rgiotoperations", "aio-instance", com.azure.core.util.Context.NONE);
    }
}
